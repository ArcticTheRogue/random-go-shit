package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	// set numbers
	numbers := [5]int{5, 1, 4, 2, 8}
	fmt.Println(numbers)
	sort(numbers[:])
}

func sort(numbers []int) {
	arrl := len(numbers) - 1
	for a := 0; a < arrl; a++ {
		for b := 0; b < arrl-a; b++ {
			if numbers[b] > numbers[b+1] {
				numbers[b], numbers[b+1] = numbers[b+1], numbers[b]
			}
		}
	}
	fmt.Println(numbers)
}
