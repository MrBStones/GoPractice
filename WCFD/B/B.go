package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Split(bufio.ScanWords)

	// Parse input
	reader.Scan()
	a := reader.Text()
	reader.Scan()
	b := reader.Text()
	// split string into 2 integers
	n, _ := strconv.Atoi(a)
	t, _ := strconv.Atoi(b)

	visitors := make([]Visitor, n)
	for i := 0; i < n; i++ {
		reader.Scan()
		c := reader.Text()
		reader.Scan()
		d := reader.Text()
		// split string into 2 integers
		bi, _ := strconv.Atoi(c)
		di, _ := strconv.Atoi(d)
		visitors[i] = Visitor{bi, di}
	}

	// Calculate result
	maxOn := 0
	isOnNow := 0
	time := 1
	for t > time {

		minTime := 10000000000
		for vis := range visitors {
			total := visitors[vis].bi + visitors[vis].di
			println("total", total)

			timemodtotal := time % total
			println("timemodtotal", timemodtotal)
			if timemodtotal == visitors[vis].bi {
				isOnNow++
				fmt.Println("on", vis)
			}
			if timemodtotal == total {
				isOnNow--
				fmt.Println("off", vis)
			}

			if timemodtotal < visitors[vis].bi {
				totes := visitors[vis].bi - timemodtotal
				if totes < minTime {
					minTime = totes
					fmt.Println("min", vis, minTime)
				}
			}

			if timemodtotal >= visitors[vis].bi {
				totes := total - timemodtotal
				if totes < minTime {
					minTime = totes
					fmt.Println("min", vis, minTime)
				}
			}

		}

		time += minTime
		fmt.Println(time)

		if isOnNow > maxOn {
			maxOn = isOnNow
		}
	}

	fmt.Println(maxOn)
}

type Visitor struct {
	bi int
	di int
}
