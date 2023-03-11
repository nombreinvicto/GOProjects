package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(n int) {
	// Your code here!
	initial := 1
	for initial <= n {
		fmt.Printf("current number is: %d\n", initial)
		switch {
		case initial%15 == 0:
			fmt.Println("FizzBuzz")
		case initial%5 == 0:
			fmt.Println("Buzz")
		case initial%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println("no condition fulfilled")
		}
		initial += 1
	}
}
func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
