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
	str := strings.Split(input.Text(), " ")

	num1, err := strconv.Atoi(str[0])
	if err != nil {
		fmt.Println("Can't convert this to an int!")
	}

	num2, err := strconv.Atoi(str[1])
	if err != nil {
		fmt.Println("Can't convert this to an int!")
	} else {
		fmt.Println(num1)
	}

	res := num1 + num2

	fmt.Print(res)
}
