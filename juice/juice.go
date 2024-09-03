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
	fruits := strings.Split(input.Text(), " ")
	input.Scan()
	ratio := strings.Split(input.Text(), " ")

	
	n, err := strconv.Atoi(input.Text())
	if err != nil {
		fmt.Println("Can't convert this to an int!")
	}

	
	fmt.Print("makes sense")
}
