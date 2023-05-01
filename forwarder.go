package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

type Forwarder struct {
	src         Address
	dst         Address
	dialTimeout int
}

func NewForwarder(src Address, dst Address, dialTimeout int) *Forwarder {
	return &Forwarder{
		src:         src,
		dst:         dst,
		dialTimeout: dialTimeout,
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
			dstConn, err := net.DialTimeout("tcp", f.dst.String(), time.Duration(f.dialTimeout)*time.Second)
			if err != nil {
				if strings.Contains(err.Error(), "i/o timeout") {
					println(fmt.Sprintf(`%s -- %s >-< %s %s`, time.Now().Format(time.RFC3339), srcConn.RemoteAddr().String(), f.dst.String(), "dial timed out"))
				} else {
					println(err.Error())
				}
				_ = srcConn.Close()
				return
			}
			println(fmt.Sprintf(`%s -- %s <-> %s`, time.Now().Format(time.RFC3339), srcConn.RemoteAddr().String(), f.dst.String()))
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
