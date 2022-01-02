package main

import (
	"fmt"
	"time"
)

func main() {

	// sequence := ".o0o="
	// sequence := "|/-\\="
	// sequence := "+x="
	sequence := "-+x\\//x+-="

	for i := 0; i < 11; i++ {
		for j := 0; j < 10; j++ {
			n := 10*i + j
			if n > 107 {
				break
			}
			fmt.Printf("\033[%dm %3d\033[m", n, n)
		}
		fmt.Println()
	}
	fmt.Printf("Progress: ")
	for i := 0; i < 80; i++ {
		for _, c := range sequence {
			fmt.Printf("\x1b[1D%c", c)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Printf("%c", sequence[len(sequence)-1])

	}

	//fmt.Println()
}
