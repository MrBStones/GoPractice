package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
)

func main() {
	fmt.Println("start")
	for i := 0; i < 10000; i++ {
		go setleader(i)
	}
}

func setleader(i int) {
	mutex.Lock()
	fmt.Println(i, "leader")
	mutex.Unlock()
}
