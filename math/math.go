package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var number1 = rand.Intn(10)
	var number2 = rand.Intn(10)
	var correctAnswer int = number1 + number2
	var yourAnswer int

	fmt.Println("What is:", number1, "+", number2, "?")
	fmt.Println("Enter answer:")
	fmt.Scanf("%d", &yourAnswer)
	fmt.Println("Your answer is:", yourAnswer)

	if yourAnswer == correctAnswer {
		fmt.Println("You are correct")
	} else {
		fmt.Println("You are wrong, the correct answer is:", correctAnswer)
	}
}
