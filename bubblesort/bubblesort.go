package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
	rand.Seed(time.Now().Unix())
	// set numbers
	numbers := [5]int{rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)}
	fmt.Println(numbers)
	sort(numbers[:])
}

func sort(numbers []int) {
	// var for length of array
	arrl := len(numbers) - 1
	// set a var; while a is less than length of array; add 1 to a
	for a := 0; a < arrl; a++ {
		// set b var; while b is less than array length minus a; add 1 to b
		for b := 0; b < arrl-a; b++ {
			// if first in set is larger than second in set
			if numbers[b] > numbers[b+1] {
				// switch sets
				numbers[b], numbers[b+1] = numbers[b+1], numbers[b]
			}
		}
	}
	fmt.Println(numbers)
}
