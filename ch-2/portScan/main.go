package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
)

func scaner(ports,results chan int){
	for p := range ports{
		conn,err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(p))
		if err != nil {
			fmt.Printf("%d off\n",)
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int,100)
	results := make(chan int)
	var openports []int
	for i:= 0;i < cap(ports);i++{
		go scaner(ports,results)
	}
	go func(){
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}