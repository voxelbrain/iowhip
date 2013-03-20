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
	2013/02/23 10:42:57 Starting 4 workers on 4 cores writing 1.000G bytes (4.000K per call)...
	2013/02/23 10:42:57 Thread 0: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T/iowhip_1361612577280469000/0
	2013/02/23 10:42:57 Thread 1: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T/iowhip_1361612577280469000/1
	2013/02/23 10:42:57 Thread 2: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T/iowhip_1361612577280469000/2
	2013/02/23 10:42:57 Thread 3: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T/iowhip_1361612577280469000/3
	2013/02/23 10:43:15 Thread 2: 17.871586s, 57.298M/s
	2013/02/23 10:43:15 Thread 0: 17.910851s, 57.172M/s
	2013/02/23 10:43:15 Thread 3: 17.970065s, 56.984M/s
	2013/02/23 10:43:15 Thread 1: 18.010156s, 56.857M/s

## Binaries

* [Darwin 386](http://filedump.surmair.de/binaries/iowhip/darwin_386/iowhip)
* [Darwin amd64](http://filedump.surmair.de/binaries/iowhip/darwin_amd64/iowhip)
* [Freebsd 386](http://filedump.surmair.de/binaries/iowhip/freebsd_386/iowhip)
* [Freebsd amd64](http://filedump.surmair.de/binaries/iowhip/freebsd_amd64/iowhip)
* [Linux 386](http://filedump.surmair.de/binaries/iowhip/linux_386/iowhip)
* [Linux amd64](http://filedump.surmair.de/binaries/iowhip/linux_amd64/iowhip)
* [Linux arm](http://filedump.surmair.de/binaries/iowhip/linux_arm/iowhip)
* [Windows 386](http://filedump.surmair.de/binaries/iowhip/windows_386/iowhip.exe)
* [Windows amd64](http://filedump.surmair.de/binaries/iowhip/windows_amd64/iowhip.exe)

---
Version 1.3.0
