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
	"time"
)

func main() {
	var finished = make(chan bool)
	var forks = make([]chan int, 5)
	for i := 0; i < 5; i++ {
		forks[i] = make(chan int)
		go fork(forks[i])
	}

	for i := 0; i < 5; i++ {
		go func(i int, forkChannel chan int) {
			phill(i, forkChannel, forks[(i+1)%5], finished)
		}(i, forks[i])
	}

	for i := 0; i < 5; i++ {
		<-finished
	}
	fmt.Println("All philosophers have eaten 3 times")
}

func phill(pNumber int, L chan int, R chan int, F chan bool) {
	var timesNomNommed = 0
	LeftForkNum := pNumber
	RightForkNum := (pNumber + 1) % 5

	for {

		// Determine if the philosopher is thinking or eating
		randnum := rand.IntN(100)
		if randnum > 50 {
			fmt.Println(pNumber, "thinking")
			time.Sleep(time.Duration(randnum) * time.Millisecond)
			continue
		}

		// Determine which fork to pick up first by comparing the fork numbers
		// Smallest numbered fork is picked up first to avoid deadlock
		if LeftForkNum < RightForkNum {
			L <- 1
			R <- 1
		} else {
			R <- 1
			L <- 1
		}

		fmt.Println(pNumber, "eating")
		time.Sleep(100 * time.Millisecond)

		res1 := <-L
		res2 := <-R

		if res1 != 2 || res2 != 2 {

		}

		fmt.Println(pNumber, "done eating")
		time.Sleep(100 * time.Millisecond)

		timesNomNommed++
		if timesNomNommed == 3 {
			fmt.Println(pNumber, "ate 3 times")
			F <- true
		}
	}
}

func fork(fc chan int) {
	for {
		<-fc // philosopher is eating
		time.Sleep(100 * time.Millisecond)
		fc <- 2 // philosopher is done eating
	}
}

// Code does not deadlock because the philosophers pick up the forks in a way that avoids deadlock.
// Every philosopher exept for the 5th philosopher picks up their left fork first.
// The 5th philosopher picks up the right fork first.
// This way, the philosopher with the smallest fork number will always pick up the left fork first.
// This way, the philosophers will never be in a situation where they are waiting for each other to put down the fork they are holding.
// This way, the philosophers can always pick up the forks they need to eat.
