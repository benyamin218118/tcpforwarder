package main

import (
	"flag"
	"fmt"
)

type Address struct {
	Host string
	Port int
}

func (a *Address) String() string {
	return fmt.Sprintf(`%s:%d`, a.Host, a.Port)
}

func main() {
	println(`
find me on telegram: @unsafepointer
`)
	listenHost := flag.String("lHost", "0.0.0.0", "listen host")
	listenPort := flag.Int("lPort", 8080, "listen port")
	remoteHost := flag.String("rHost", "8.8.8.8", "remote host")
	remotePort := flag.Int("rPort", 53, "remote port")
	dialTimeout := flag.Int("timeout", 4, "dial timeout in seconds")
	help := flag.Bool("help", false, "print help")

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}

	if *dialTimeout <= 0 || *dialTimeout > 32 {
		panic("invalid dial timeout, it should be bigger than 0 and smaller than 32")
	}

	src := Address{
		Host: *listenHost,
		Port: *listenPort,
	}
	dst := Address{
		Host: *remoteHost,
		Port: *remotePort,
	}

	println(fmt.Sprintf(`tcp://%s <-> tcp://%s`+"\n", src.String(), dst.String()))

	forwarder := NewForwarder(src, dst, *dialTimeout)

	forwarder.Start()
}
