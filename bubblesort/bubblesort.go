package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	numbers := [5]int{5, 1, 4, 2, 8}
	fmt.Println("Numbers before sort:", numbers)
	var a = 0
	var b = 1
	fmt.Println(numbers[a], numbers[b])
	if a < 3 && b < 4 {
		if numbers[a] > numbers[b] {
			numbers[a], numbers[b] = numbers[b], numbers[a]
			fmt.Println(numbers)
			return
		}
	}
}
