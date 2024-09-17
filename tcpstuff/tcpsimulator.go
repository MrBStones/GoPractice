package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var channel = make(chan int)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		client()
		wg.Done()
	}()
	go func() {
		server()
		wg.Done()
	}()
	wg.Wait()
}

func sendHandshake(SYN int) {
	ch := make(chan int)
	go func() {
		ch <- SYN
	}()
	select {
	case <-ch:
		fmt.Println("Client sending SYN sequence", SYN)
		channel <- SYN
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! Retrying...")
		sendHandshake(SYN)
	}
}

func receiveHandshake() int {
	select {
	case res := <-channel:
		fmt.Println("Client recived", res)
		return res
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! Retrying...")
		return receiveHandshake()
	}
}

func client() {
	SYN := rand.IntN(100)
	sendHandshake(SYN)

	// Wait for server response
	ACKR := receiveHandshake()
	SYNR := receiveHandshake()
	if ACKR == SYN+1 {
		fmt.Println("Handshake successful! Client recived:", ACKR, SYNR)
		channel <- SYNR + 1
		channel <- ACKR
	} else {
		fmt.Println("Handshake unsuccessful! Client recived:", ACKR, "but expected", SYN+1)
	}

}

func server() {
	ACK := <-channel
	fmt.Println("Server received ACK sequence", ACK)

	SYN := rand.IntN(100)
	ACK++
	fmt.Println("Server sending SYN-ACK sequence", SYN, ACK)
	channel <- ACK
	channel <- SYN

	ACKR := <-channel
	SYNR := <-channel
	if ACKR == SYN+1 {
		fmt.Println("Handshake successful! Server recived:", ACKR, SYNR)
	} else {
		fmt.Println("Handshake unsuccessful! Server recived:", ACKR, "but expected", SYN+1)
	}
}
