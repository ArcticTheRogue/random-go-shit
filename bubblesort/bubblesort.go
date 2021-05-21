package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	// set numbers
	numbers := [5]int{5, 1, 4, 2, 8}
	fmt.Println("Numbers before sort:", numbers)
	// set vars for first set of numbers
	var a = 0
	var b = 1
	// set boolean to see if it is sorted
	sorted := false
	fmt.Println(numbers[a], numbers[b])
	// set for loop for if set does is less or equal to the last set and it is not sorted
	for a <= 3 && b <= 4 && sorted == false {
		// if statement for if the first number is larger than the second
		if numbers[a] > numbers[b] {
			//swap numbers
			numbers[a], numbers[b] = numbers[b], numbers[a]
			fmt.Println(numbers)
			// add one to each var so it moves to the next set
			a += 1
			b += 1
		} else {
			// same as end of other if statement
			a += 1
			a += 1
		}
	}
}
