`iowhip` is a simple I/O stressing tool.

## Installation

	go get github.com/voxelbrain/iowhip

## Usage

	$ iowhip --help
	Usage: iowhip [global options]

	Global options:
	    -c, --cores      Number of cores to use (default: 4)
	    -t, --threads    Number of threads to use (default: 4)
	    -b, --block-size Number of bytes to write with each call (default: 4.000K)
	    -f, --file-size  Number of zeroes to write to each file (*)
	    -o, --output-dir Output directory (default: /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T/)
	    -s, --sync       Sync after every written block
	    -k, --keep-files Dont delete files when done
	        --dsync      Open files with O_DSYNC (Linux only)
	        --direct     Open files with O_DIRECT (Linux only)
	        --osync      Open files with O_SYNC
	    -h, --help       Show this help

	$ iowhip -f 1G
	2013/03/20 16:04:42 Starting 4 workers on 4 cores writing 1.000G bytes (4.000K per call)...
	2013/03/20 16:04:42 Thread 0: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T/iowhip_1363791882334221000/0
	2013/03/20 16:04:42 Thread 3: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T/iowhip_1363791882334221000/3
	2013/03/20 16:04:42 Thread 1: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T/iowhip_1363791882334221000/1
	2013/03/20 16:04:42 Thread 2: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T/iowhip_1363791882334221000/2

	  Thread  Total Time  Total Speed  Writing Time  Writing Speed
	       1  19.995106s    51.213M/s     19.99464s      51.214M/s
	       2   20.07883s    50.999M/s    20.078598s      51.000M/s
	       0  20.237617s    50.599M/s    20.237483s      50.599M/s
	       3  20.302615s    50.437M/s    20.302451s      50.437M/s

## Binaries

* [Darwin 386](http://downloads.voxelbrain.com/iowhip/master/darwin_386/iowhip)
* [Darwin amd64](http://downloads.voxelbrain.com/iowhip/master/darwin_amd64/iowhip)
* [Linux 386](http://downloads.voxelbrain.com/iowhip/master/linux_386/iowhip)
* [Linux amd64](http://downloads.voxelbrain.com/iowhip/master/linux_amd64/iowhip)
* [Linux arm](http://downloads.voxelbrain.com/iowhip/master/linux_arm/iowhip)

---
Version 1.3.2
