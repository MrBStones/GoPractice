package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	cards := strings.Split(input.Text(), "")
	counts := make(map[string]bool)
	P, K, H, T := 13, 13, 13, 13
	for i := 0; i < len(cards); i += 3 {
		str := cards[i] + cards[i+1] + cards[i+2]

		if counts[str] == true {
			fmt.Print("GRESKA")
			return
		}

		counts[str] = true
		if cards[i] == "P" {
			P--
		}
		if cards[i] == "K" {
			K--
		}
		if cards[i] == "H" {
			H--
		}
		if cards[i] == "T" {
			T--
		}
	}
	fmt.Print(P, " ", K, " ", H, " ", T)
}
