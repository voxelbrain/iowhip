`iowhip` is simple I/O stressing tool.

## Installation

	go get github.com/voxelbrain/iowhip

## Usage

	$ iowhip --help
	Usage: iowhip [global options]

	Global options:
	    -c, --cores      Number of cores to use (default: 4)
	    -t, --threads    Number of threads to use (default: 4)
	    -b, --block-size Number of bytes to write with each call (default: 4.000K)
	    -f, --file-size  Number of zeroes to write to each file (default: 10.000G)
	    -k, --keep-files Dont delete files when done
	    -h, --help       Show this help

	$ iowhip
	2013/02/22 12:44:03 Starting 4 workers on 4 cores writing 10.000G bytes (4.000K per call)...
	2013/02/22 12:44:03 Thread 3: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T//iowhip_3_1361533443840381000
	2013/02/22 12:44:03 Thread 1: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T//iowhip_1_1361533443840380000
	2013/02/22 12:44:03 Thread 2: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T//iowhip_2_1361533443840401000
	2013/02/22 12:44:03 Thread 0: Using /var/folders/k4/4rlzbxl14d9_2k_jwyfv_nzh0000gn/T//iowhip_0_1361533443840421000
	2013/02/22 12:47:07 Thread 1: 2m56.284958s, 58.088M/s
	2013/02/22 12:47:07 Thread 0: 2m58.190676s, 57.467M/s
	2013/02/22 12:47:07 Thread 2: 2m59.765777s, 56.963M/s
	2013/02/22 12:47:07 Thread 3: 3m3.512097s, 55.800M/s


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
Version 1.1.0
