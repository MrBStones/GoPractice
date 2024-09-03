package main

import (
	"fmt"
)

func main() {

	fmt.Println("REFERENCE EXAMPLE")
	var i int
	i = 5
	var pointer *int
	pointer = &i
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(pointer)

	fmt.Println("SLICES")
	// init 1
	var numbers []int
	numbers = []int{1, 2, 3, 5, 4}
	fmt.Println(numbers)

	// init 2
	var numbers2 = []int{99, 100, 90, 101, 69}
	fmt.Println(numbers2)

	// init 3
	numbers3 := []int{69, 420}
	fmt.Println(numbers3)

	// init 4, with allocation
	var numbers4 = make([]int, 2, 2) // length 2, capacity 2
	fmt.Println(numbers4, "cap", cap(numbers4))

	numbers4[0] = 1
	numbers4[1] = 2
	numbers4 = append(numbers4, 69)
	fmt.Println(numbers4, "cap:", cap(numbers4))

	numout(123)
	a, b := numfourtwenty(1)
	numout(a)
	numout(b)
}

func numout(i int) {
	fmt.Println(i)
}

func numfourtwenty(i int) (int, int) {
	return i + 420, i
}
