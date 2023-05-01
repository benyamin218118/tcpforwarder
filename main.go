package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Address struct {
	Host string
	Port int
}

func (a *Address) String() string {
	return fmt.Sprintf(`%s:%d`, a.Host, a.Port)
}

func main() {
	listenHost := flag.String("lHost", "0.0.0.0", "listen host")
	listenPort := flag.String("lPort", "8080", "listen port ( can be a port range like 1024-5555 too )")
	remoteHost := flag.String("rHost", "8.8.8.8", "remote host")
	remotePort := flag.Int("rPort", 53, "remote port ( will be equal to lPort when doing range forward )")
	dialTimeout := flag.Int("timeout", 4, "dial timeout in seconds")
	help := flag.Bool("help", false, "print help")
	h := flag.Bool("h", false, "")

	flag.Parse()
	if *help || *h {
		println(`
find me on telegram: @unsafepointer
`)
		flag.PrintDefaults()
		return
	}

	if *dialTimeout <= 0 || *dialTimeout > 32 {
		panic("invalid dial timeout, it should be bigger than 0 and smaller than 32")
	}
	src := Address{}
	dst := Address{
		Host: *remoteHost,
		Port: *remotePort,
	}

	if strings.Contains(*listenPort, "-") {
		pRange := strings.Split(*listenPort, "-")
		if len(pRange) > 2 {
			println("[E] invalid port range format, should be like 1024-5555")
			os.Exit(1)
		}
		start, err := strconv.Atoi(pRange[0])
		if err != nil || start < 1 || start > 65534 {
			println(`[E] invalid port range start, should be a Number between 1 and 65534`)
			os.Exit(1)
		}
		end, err := strconv.Atoi(pRange[1])
		if err != nil || end <= start || end > 65534 {
			println(`[E] invalid port range start, should be a Number between 1 and 65534 ( start should be smaller than end too )`)
			os.Exit(1)
		}
		for port := start; port <= end; port++ {
			println(fmt.Sprintf(`[I] tcp://%s:%d <-> tcp://%s:%d`+"\n", *listenHost, port, *remoteHost, port))
			go startForwarder(Address{
				Host: *listenHost,
				Port: port,
			}, Address{
				Host: *remoteHost,
				Port: port,
			}, *dialTimeout)
		}
	} else {
		n, err := strconv.Atoi(*listenPort)
		if err != nil {
			println("[E] invalid listenPort, should be a Number between 1 and 65534")
			os.Exit(1)
		}
		src.Host = *listenHost
		src.Port = n
		println(fmt.Sprintf(`[I] tcp://%s <-> tcp://%s`+"\n", src.String(), dst.String()))
		go startForwarder(src, dst, *dialTimeout)
	}
	select {}
}

func startForwarder(src, dst Address, dialTimeout int) {
	forwarder := NewForwarder(src, dst, dialTimeout)
	forwarder.Start()
}
