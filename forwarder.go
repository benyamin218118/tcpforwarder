package main

import (
	"io"
	"net"
)

type Forwarder struct {
	src Address
	dst Address
}

func NewForwarder(src Address, dst Address) *Forwarder {
	return &Forwarder{
		src: src,
		dst: dst,
	}
}

func (f *Forwarder) Start() {
	listener, err := net.Listen("tcp", f.src.String())
	panicIfErr(err)

	for {
		srcConn, err := listener.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		go func(srcConn net.Conn) {
			dstConn, err := net.Dial("tcp", f.dst.String())
			if err != nil {
				println(err.Error())
				_ = srcConn.Close()
				return
			}
			go func(srcConn, dstConn net.Conn) {
				_, _ = io.Copy(srcConn, dstConn)
				_ = srcConn.Close()
				_ = dstConn.Close()
			}(srcConn, dstConn)
			go func(srcConn, dstConn net.Conn) {
				_, _ = io.Copy(dstConn, srcConn)
				_ = srcConn.Close()
				_ = dstConn.Close()
			}(srcConn, dstConn)
		}(srcConn)
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
