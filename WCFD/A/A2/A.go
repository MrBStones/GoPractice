package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a [4][6]bool
	bottles := 0

	for i := 0; i < 4; i++ {

		text, _ := reader.ReadString('\n')
		// split string into char array
		count := 0
		for _, c := range text {
			if c == 'o' {
				a[i][count] = true
				bottles++
			}
			if c == '.' {
				a[i][count] = false
			}
			count++
		}
	}

	// cluster together all the true values

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			if bottles > 0 {
				fmt.Print("o")
				bottles--
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
