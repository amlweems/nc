package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	listen := flag.Bool("l", false, "listening mode")
	flag.Parse()

	if *listen {
		fmt.Fprintf(os.Stderr, "listening on %s\n", flag.Arg(0))
		listener, err := net.Listen("tcp", ":"+flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer listener.Close()

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		go io.Copy(os.Stdout, conn)
		io.Copy(conn, os.Stdin)
	} else {
		fmt.Fprintf(os.Stderr, "connected to %s\n", flag.Arg(0))
		conn, err := net.Dial("tcp", flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		go io.Copy(os.Stdout, conn)
		io.Copy(conn, os.Stdin)
	}
}
