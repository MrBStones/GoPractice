package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	for i := 0; i < 3; i++ {
		fmt.Print(str, " ")
	}

}
