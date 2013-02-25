package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/voxelbrain/goptions"
)

const (
	VERSION = "1.2.0"
)

var (
	DefaultBlocksize = 4 * KiloByte
	options          = struct {
		Cores     int           `goptions:"-c, --cores, description='Number of cores to use'"`
		Threads   int           `goptions:"-t, --threads, description='Number of threads to use'"`
		Blocksize *Datasize     `goptions:"-b, --block-size, description='Number of bytes to write with each call'"`
		Filesize  *Datasize     `goptions:"-f, --file-size, obligatory, description='Number of zeroes to write to each file'"`
		OutputDir string        `goptions:"-o, --output-dir, description='Output directory'"`
		Sync      bool          `goptions:"-s, --sync, description='Sync after every written block'"`
		KeepFiles bool          `goptions:"-k, --keep-files, description='Dont delete files when done'"`
		Help      goptions.Help `goptions:"-h, --help, description='Show this help'"`
	}{
		Cores:     runtime.NumCPU(),
		Threads:   runtime.NumCPU(),
		Blocksize: &DefaultBlocksize,
		OutputDir: os.TempDir(),
	}
	Timestamp = fmt.Sprintf("%d", time.Now().UnixNano())
)

func init() {
	goptions.ParseAndFail(&options)
	options.OutputDir = filepath.Clean(options.OutputDir + "/iowhip_" + Timestamp)
}

type Result struct {
	Index       int
	TotalTime   time.Duration
	WritingTime time.Duration
}

func main() {
	err := os.MkdirAll(options.OutputDir, os.FileMode(0755))
	if err != nil {
		log.Fatalf("Could not create output directory %s: %s", options.OutputDir, err)
	}
	cleanup := make(chan os.Signal)
	signal.Notify(cleanup, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	go func() {
		<-cleanup
		if !options.KeepFiles {
			os.RemoveAll(options.OutputDir)
		}
	}()

	runtime.GOMAXPROCS(options.Cores)
	log.Printf("Starting %d workers on %d cores writing %s bytes (%s per call)...", options.Threads, options.Cores, options.Filesize, options.Blocksize)

	results := make(chan Result)
	wg := &sync.WaitGroup{}
	wg.Add(options.Threads)
	for i := 0; i < options.Threads; i++ {
		go writeFile(i, results, wg)
	}
	wg.Wait()

	prettyPrint(results)
	cleanup <- syscall.SIGINT
}

func prettyPrint(results <-chan Result) {
	w := tabwriter.NewWriter(os.Stdout, 3, 1, 2, ' ', tabwriter.AlignRight)
	defer w.Flush()
	fmt.Fprintf(w, "\nThread\tTotal Time\tTotal Speed\tWriting Time\tWriting Speed\t\n")
	for i := 0; i < options.Threads; i++ {
		r := <-results
		fmt.Fprintf(w, "%d\t%s\t%s/s\t%s\t%s/s\t\n", r.Index,
			r.TotalTime, Datasize(int64(float64(*options.Filesize)*float64(time.Second)/float64(r.TotalTime))),
			r.WritingTime, Datasize(int64(float64(*options.Filesize)*float64(time.Second)/float64(r.WritingTime))))
	}
}

func writeFile(idx int, c chan Result, wg *sync.WaitGroup) {
	result := Result{
		Index: idx,
	}
	defer func() { c <- result }()
	defer wg.Done()
	filename := fmt.Sprintf("%s/%d", options.OutputDir, idx)
	log.Printf("Thread %d: Using %s", idx, filename)

	amount := *options.Filesize
	data := make([]byte, int(*options.Blocksize))
	start_creating := time.Now()
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("Thread %d: Could not open file %s: %s", idx, filename, err)
		return
	}
	defer f.Close()

	start_writing := time.Now()
	for amount > 0 {
		n, err := f.Write(data)
		if err != nil {
			log.Printf("Thread %d: Write to %s failed: %s", idx, filename, err)
			return
		}
		amount -= Datasize(n)
		if options.Sync {
			f.Sync()
		}
	}
	f.Sync()
	result.WritingTime = time.Since(start_writing)
	f.Close()
	result.TotalTime = time.Since(start_creating)
}
