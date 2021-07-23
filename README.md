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
       CPU : 8 processors, AMD EPYC 7K62 48-Core Processor
    Memory : 16.0 GiB

# Go version and settings
Go Version : go1.16.5 linux/amd64
GOMAXPROCS : 8

# Netwokr settings
TCP connections : 500/1000/5000/10000
Packet size     : 512/1024/2048/4096/8192/16384/32768/65536
Test duration   : 15s
```

#### Echo Server

![](https://github.com/panjf2000/gnet_benchmarks/raw/master/results/echo_conn_linux.png)

![](https://github.com/panjf2000/gnet_benchmarks/raw/master/results/echo_packet_linux.png)