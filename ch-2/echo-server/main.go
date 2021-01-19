package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func echoService(conn net.Conn,done chan int){
	defer conn.Close()
//	v1.0 带有缓冲区的低级调用
//	buf := make([]byte,512)
//	for{
//		size,err := conn.Read(buf)
//		if err == io.EOF {
//			fmt.Println("client disconnected")
//			break
//		}
//		if err != nil {fmt.Println(err)}
//		fmt.Printf("received %d bytes\n> %s",size,string(buf))
//		if string(buf[:size]) == "exit\n"{done<-1}
//		fmt.Println(buf[:size])
//		if _,err := conn.Write(buf[:size]);err != nil{
//			fmt.Println("cannot write data")
//		}
//
//	}
//	v2.0 bufio 使用
//	for{
//		reader:=bufio.NewReader(conn)
//		s,err := reader.ReadString('\n')
//		if err == io.EOF {
//			fmt.Println("client disconnected")
//			break
//		}
//		if err != nil {fmt.Println(err)}
//		fmt.Printf("received %d bytes\n> %s",len(s),s)
//		if s == "exit\n"{done<-1}
//		writer := bufio.NewWriter(conn)
//		if _,err := writer.WriteString(s);err != nil{
//			fmt.Println("cannot write data")
//		}
//		writer.Flush()
//	}
//	v3.0 直接io.Copy，但是这样不能捕获用户输入，适用于不需要对传输数据做处理的场景（譬如端口转发
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
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
	done := make(chan int)
	go func() {
		for{
			conn, err:=listener.Accept()
			if err != nil{
				fmt.Println(err)
				continue
			}
			go echoService(conn,done)
		}
	}()
	<-done
}
