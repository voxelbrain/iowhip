package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/voxelbrain/goptions"
)

const (
	VERSION = "1.0.1"
)

var (
	DefaultFilesize  = 10 * GigaByte
	DefaultBlocksize = 4 * KiloByte
	options          = struct {
		Cores     int           `goptions:"-c, --cores, description='Number of cores to use'"`
		Threads   int           `goptions:"-t, --threads, description='Number of threads to use'"`
		Blocksize *Datasize     `goptions:"-b, --block-size, description='Number of bytes to write with each call'"`
		Filesize  *Datasize     `goptions:"-f, --file-size, description='Number of zeroes to write to each file'"`
		KeepFiles bool          `goptions:"-k, --keep-files, description='Dont delete files when done'"`
		Help      goptions.Help `goptions:"-h, --help, description='Show this help'"`
	}{
		Cores:     runtime.NumCPU(),
		Threads:   runtime.NumCPU(),
		Filesize:  &DefaultFilesize,
		Blocksize: &DefaultBlocksize,
	}
)

func init() {
	goptions.ParseAndFail(&options)
}

type Result struct {
	Index    int
	Duration time.Duration
}

func main() {
	runtime.GOMAXPROCS(options.Cores)
	log.Printf("Starting %d workers on %d cores writing %s bytes (%s per call)...", options.Threads, options.Cores, options.Filesize, options.Blocksize)

	results := make(chan Result, options.Threads)
	wg := &sync.WaitGroup{}
	wg.Add(options.Threads)
	for i := 0; i < options.Threads; i++ {
		go writeFile(i, results, wg)
	}
	wg.Wait()
	for i := 0; i < options.Threads; i++ {
		r := <-results
		log.Printf("Thread %d: %s, %s/s", r.Index, r.Duration, Datasize(int64(float64(*options.Filesize)*float64(time.Second)/float64(r.Duration))))
	}
	close(results)
}

func writeFile(idx int, c chan Result, wg *sync.WaitGroup) {
	result := Result{
		Index: idx,
	}
	defer func() { c <- result }()
	defer wg.Done()
	filename := fmt.Sprintf("%s/iowhip_%d_%d", os.TempDir(), idx, time.Now().UnixNano())
	log.Printf("Thread %d: Using %s", idx, filename)
	if !options.KeepFiles {
		defer os.Remove(filename)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Printf("Thread %d: Could not open file %s: %s", idx, filename, err)
		return
	}
	defer f.Close()

	amount := *options.Filesize
	data := make([]byte, int(*options.Blocksize))
	start := time.Now()
	for amount > 0 {
		n, err := f.Write(data)
		if err != nil {
			log.Printf("Thread %d: Write to %s failed: %s", idx, filename, err)
			return
		}
		amount -= Datasize(n)
	}
	f.Sync()
	result.Duration = time.Now().Sub(start)
}
