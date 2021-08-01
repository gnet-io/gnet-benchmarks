## gnet benchmark tools

Required tools:

- [wrk](https://github.com/wg/wrk) for HTTP
- [tcpkali](https://github.com/machinezone/tcpkali) for Echo

Required Go packages:

```
go get gonum.org/v1/plot/...
go get -u github.com/valyala/fasthttp
```

And of course [Go](https://golang.org) is required.

Run `bench.sh` for all benchmarks.

## Notes

- The current results were run on both Linux and FreeBSD.
- The servers started in multiple-threaded mode (GOMAXPROCS=Default).
- Network clients connected over Ipv4 localhost.

Like all benchmarks ever made in the history of whatever, YMMV. Please tweak and run in your environment and let me know if you see any glaring issues.

# Benchmark Test

## On Linux (epoll)

### Test Environment

```powershell
# Machine information
        OS : Ubuntu 20.04/x86_64
       CPU : 8 CPU cores, AMD EPYC 7K62 48-Core Processor
    Memory : 16.0 GiB

# Go version and settings
Go Version : go1.16.5 linux/amd64
GOMAXPROCS : 8

# Network settings
TCP connections : 500/1000/5000/10000
Packet size     : 512/1024/2048/4096/8192/16384/32768/65536 bytes
Test duration   : 15s
```

#### Echo benchmark

![](https://github.com/panjf2000/gnet_benchmarks/raw/master/results/echo_conn_linux.png)

![](https://github.com/panjf2000/gnet_benchmarks/raw/master/results/echo_packet_linux.png)

## On MacOS (kqueue)

### Test Environment

```powershell
# Machine information
        OS : MacOS Big Sur/x86_64
       CPU : 6 CPU cores, Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
    Memory : 16.0 GiB

# Go version and settings
Go Version : go1.16.5 darwin/amd64
GOMAXPROCS : 12

# Network settings
TCP connections : 300/400/500/600/700
Packet size     : 512/1024/2048/4096/8192 bytes
Test duration   : 15s
```

#### Echo benchmark

![](https://github.com/panjf2000/gnet_benchmarks/raw/master/results/echo_conn_macos.png)

![](https://github.com/panjf2000/gnet_benchmarks/raw/master/results/echo_packet_macos.png)