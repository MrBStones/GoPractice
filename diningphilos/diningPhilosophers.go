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

var finished = make(chan bool)
var times = 3
var timeunit = time.Millisecond

func main() {

	// make L and R Channel per fork
	fcL1 := make(chan bool)
	fcL2 := make(chan bool)
	fcL3 := make(chan bool)
	fcL4 := make(chan bool)
	fcL5 := make(chan bool)
	fcR1 := make(chan bool)
	fcR2 := make(chan bool)
	fcR3 := make(chan bool)
	fcR4 := make(chan bool)
	fcR5 := make(chan bool)

	// spawn threads
	go fork(fcL1, fcR1)
	go fork(fcL2, fcR2)
	go fork(fcL3, fcR3)
	go fork(fcL4, fcR4)
	go fork(fcL5, fcR5)

	go phill(1, fcR1, fcL2)
	go phill(2, fcR2, fcL3)
	go phill(3, fcR3, fcL4)
	go phill(4, fcR4, fcL5)
	go phill(5, fcR5, fcL1)

	for i := 0; i < 5; i++ {
		<-finished
	}
	fmt.Println("All philosophers have eaten", times, "times")
}

func phill(n int, L chan bool, R chan bool) {
	var timesNomNommed = 0
	LeftForkNum := n
	RightForkNum := (n + 1) % 5

	for {
		// Determine if the philosopher is thinking or eating
		randnum := rand.IntN(100)
		if randnum > 50 {
			fmt.Println(n, "thinking")
			time.Sleep(time.Duration(randnum) * timeunit)
			continue
		}

		// Determine which fork to pick up first by comparing the fork numbers
		// Smallest numbered fork is picked up first to avoid deadlock
		if LeftForkNum < RightForkNum {
			L <- true
			R <- true
		} else {
			R <- true
			L <- true
		}

		fmt.Println(n, "eating")
		time.Sleep(time.Duration(randnum+100) * timeunit)

		<-L
		fmt.Println(n, "put down left fork")
		<-R
		fmt.Println(n, "put down right fork")

		fmt.Println(n, "done eating")
		time.Sleep(100 * timeunit)

		timesNomNommed++
		if timesNomNommed == times {
			fmt.Println(n, "ate", times, "times")
			finished <- true
		}
	}
}

func fork(L chan bool, R chan bool) {
	for {
		// select the first channel with a signal
		select {
		case <-L:
			L <- false
		case <-R:
			R <- false
		}
	}
}

// Code does not deadlock because the philosophers pick up the forks in a way that avoids deadlock.
// Every philosopher exept for the 5th philosopher picks up their left fork first.
// The 5th philosopher picks up the right fork first.
// This way, if every philosopher tries to pick up forks at the same time, at least one will succeed

// --- output --- (ofc this is non deterministic so results will never be the same)
// 5 thinking
// 1 thinking
// 2 thinking
// 3 eating
// 1 thinking
// 5 thinking
// 1 thinking
// 3 put down left fork
// 3 put down right fork
// 3 done eating
// 4 eating
// 2 eating
// 5 thinking
// 5 thinking
// 3 thinking
// 4 put down left fork
// 4 put down right fork
// 4 done eating
// 2 put down left fork
// 2 put down right fork
// 1 eating
// 2 done eating
// 5 thinking
// 3 eating
// 2 thinking
// 1 put down left fork
// 1 put down right fork
// 1 done eating
// 5 thinking
// 3 put down left fork
// 3 put down right fork
// 3 done eating
// 4 eating
// 2 eating
// 1 thinking
// 1 thinking
// 4 put down left fork
// 4 put down right fork
// 4 done eating
// 5 eating
// 2 put down left fork
// 2 put down right fork
// 2 done eating
// 3 eating
// 1 thinking
// 1 thinking
// 5 put down left fork
// 5 put down right fork
// 5 done eating
// 3 put down left fork
// 3 put down right fork
// 3 done eating
// 4 eating
// 2 eating
// 1 thinking
// 3 ate 3 times
// 4 put down left fork
// 4 put down right fork
// 4 done eating
// 5 eating
// 2 put down left fork
// 2 put down right fork
// 2 done eating
// 3 eating
// 4 ate 3 times
// 2 ate 3 times
// 2 thinking
// 5 put down left fork
// 5 put down right fork
// 5 done eating
// 1 eating
// 3 put down left fork
// 3 put down right fork
// 3 done eating
// 4 eating
// 2 thinking
// 1 put down left fork
// 1 put down right fork
// 1 done eating
// 4 put down left fork
// 4 put down right fork
// 4 done eating
// 5 eating
// 3 eating
// 4 thinking
// 3 put down left fork
// 3 put down right fork
// 3 done eating
// 2 eating
// 5 put down left fork
// 5 put down right fork
// 5 done eating
// 4 thinking
// 3 thinking
// 5 ate 3 times
// 2 put down left fork
// 2 put down right fork
// 2 done eating
// 1 eating
// 3 thinking
// 4 thinking
// 3 thinking
// 1 put down left fork
// 1 put down right fork
// 1 done eating
// 5 eating
// 2 eating
// 3 thinking
// 1 ate 3 times
// All philosophers have eaten 3 times
