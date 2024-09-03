package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, err := strconv.Atoi(input.Text())
	if err != nil {
		fmt.Println("Can't convert this to an int!")
	}
	input.Scan()
	str := strings.Split(input.Text(), " ")
	for i := 0; i < n; i++ {
		c, err := strconv.Atoi(str[i])
		if err != nil {
			continue
		}

		if c != i+1 {
			fmt.Print("something is fishy")
			return
		}
	}
	fmt.Print("makes sense")
}
