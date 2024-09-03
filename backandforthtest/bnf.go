package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {
	go backnf()
	fmt.Println("helloWorld")
	time.Sleep(5 * time.Second)
}

func backnf() {
	var i int
	for {
		i++
		fmt.Println("sending", i)
		ch <- i
		fmt.Printf("it never makes it here since its waiting for the message to be recived")
		x := <-ch
		fmt.Println("reciving", x)
	}
}
