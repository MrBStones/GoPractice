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

- the system must be designed in a way that does not lead to  a deadlock (and each philosopher must eat at least 3 times).
Comment in the code why the system does not deadlock.

- A sequentialisation of the system (executing only one philosopher at a time) is not acceptable. I.e., philosophers must
be able to request a fork at any time.

- Philosophers must display (print on screen) any state change (eating or thinking) during their execution.


used https://en.wikipedia.org/wiki/Dining_philosophers_problem
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var arbiter sync.Mutex // makes sure only one phill can use the waiter at a time
var ateThrice = make(chan bool)
var minTimes int = 3

func main() {
	// make a Channel per fork
	fc1 := make(chan int)
	fc2 := make(chan int)
	fc3 := make(chan int)
	fc4 := make(chan int)
	fc5 := make(chan int)

	// spawn threads
	go fork(fc1)
	go fork(fc2)
	go fork(fc3)
	go fork(fc4)
	go fork(fc5)

	go phill(1, fc1, fc2)
	go phill(2, fc2, fc3)
	go phill(3, fc3, fc4)
	go phill(4, fc4, fc5)
	go phill(5, fc5, fc1)

	// recive 5 communications and throw them out
	<-ateThrice
	<-ateThrice
	<-ateThrice
	<-ateThrice
	<-ateThrice
	fmt.Println("EVERYONE ATE", minTimes, "TIMES")
}

func phill(pNumber int, LeftFork chan int, RightFork chan int) {
	var timesAte int
	for {
		// should bro think or eat ?
		if r.Intn(3) == 2 {
			fmt.Println(pNumber, "is asking for the waiter")
			arbiter.Lock()
			fmt.Println(pNumber, "got the waiter")
			result := waiter(LeftFork, RightFork)
			if result {
				fmt.Println(pNumber, "got the okay to take forks")
				LeftFork <- pNumber
				RightFork <- pNumber
			}
			arbiter.Unlock()
			fmt.Println(pNumber, "finished with waiter")

			if result {
				fmt.Println(pNumber, "is eating")
				// time.Sleep(time.Second)

				arbiter.Lock()
				fmt.Println(pNumber, "is dropping their forks")
				LeftFork <- -1
				RightFork <- -1
				arbiter.Unlock()

				timesAte++
				if timesAte == minTimes {
					fmt.Println(pNumber, "ate", minTimes, "times")
					ateThrice <- true
				}
				continue
			}
		}

		// bro is thinking
		fmt.Println(pNumber, "is thinking")
		// time.Sleep(time.Second / 2)
	}
}

func waiter(LeftFork chan int, RightFork chan int) bool {
	LeftFork <- -2
	res1 := <-LeftFork
	RightFork <- -2
	res2 := <-RightFork

	if res1 == 1 || res2 == 1 { // are the forks in use
		return false
	} else {
		return true
	}
}

func fork(Channel chan int) {
	var heldBy int = 0
	for {
		x := <-Channel
		if x == -2 {
			// report if fork is grabbed
			if heldBy == 0 {
				Channel <- -1 // -1 fork is not grabbed
			} else {
				Channel <- 1 // 1 fork is grabbed
			}
			continue
		}

		if x == -1 {
			// drop fork
			fmt.Println(heldBy, "dropped a fork")
			heldBy = 0
			continue
		}

		if heldBy != 0 {
			fmt.Println(x, "tried to grab a fork that is allready held by", heldBy, "!!!")
			continue
		}

		// any other number is good
		heldBy = x
		fmt.Println(heldBy, "grabbed a fork")
	}
}
