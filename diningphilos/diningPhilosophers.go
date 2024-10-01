/*
The Dining Philosophers is a well-known problem in Computer Science that concerns concurrency. At a dining round table,
there are five philosophers who are supposed to have dinner. Philosophers are kind of special and while they have dinner,
they either *eat* their food xor *think* about something. In order to be able to eat, they must get hold of two forks
(the food is very special and cannot be handled with one fork). Unfortunately, there are only five forks at the table, each
of them uniquely placed between two neighbouring philosophers (the table is round, there is exactly one fork between any two
neighbouring philosophers -- each philosopher can only reach the two forks that are nearby). As a consequence, it is never
the case that all philosophers can eat at the same time (max two at a time).  Eating is not limited by food quantity or
stomach space (which are both assumed to be infinite). This problem is interesting because, depending on how they decide to
pick the forks, the philosopher may reach a deadlock.

The goal of this project is to implement the dining philosophers problem in Go, with the following requirements:

- Each fork must have its own thread (goroutine)

- Each philosopher must have its own thread (goroutine)

- Philosophers and forks must communicate with each other *only* by  using channels

- the system must be designed in a way that does not lead to a deadlock (and each philosopher must eat at least 3 times).
Comment in the code why the system does not deadlock.

- A sequentialisation of the system (executing only one philosopher at a time) is not acceptable. I.e., philosophers must
be able to request a fork at any time.

- Philosophers must display (print on screen) any state change (eating or thinking) during their execution.
*/

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var arbiter1 sync.Mutex
var arbiter2 sync.Mutex

func main() {
	var wgP sync.WaitGroup
	var forks = make([]chan int, 5)
	for i := 0; i < 5; i++ {
		forks[i] = make(chan int)
		go fork(forks[i])
	}

	for i := 0; i < 5; i++ {
		wgP.Add(1)
		go func(i int, forkChannel chan int) {
			defer wgP.Done()
			phill(i, forkChannel, forks[(i+1)%5])
		}(i, forks[i])
	}

	wgP.Wait()
	fmt.Println("All philosophers have eaten 3 times")
}

func phill(pNumber int, L chan int, R chan int) {
	var timesNomNommed = 0
	for {
		if rand.IntN(100) < 50 {
			fmt.Println(pNumber, "thinking")
			time.Sleep(100 * time.Millisecond)
			continue
		}

		arbiter1.Lock()
		L <- 1
		R <- 1
		arbiter1.Unlock()

		fmt.Println(pNumber, "eating")
		time.Sleep(100 * time.Millisecond)

		arbiter2.Lock()
		<-L
		<-R
		arbiter2.Unlock()

		fmt.Println(pNumber, "done eating")
		time.Sleep(100 * time.Millisecond)

		timesNomNommed++
		if timesNomNommed == 3 {
			break
		}
	}
}

func fork(fc chan int) {
	for {
		<-fc // philosopher is eating
		time.Sleep(100 * time.Millisecond)
		fc <- 1 // philosopher is done eating
	}
}

// Code does not deadlock because the arbiter1 and arbiter2 mutexes are used to ensure that the forks are not taken by two at the same time
