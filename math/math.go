package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var number1 = rand.Intn(5)
	var number2 = rand.Intn(5)
	fmt.Println("What is:", number1, "+", number2, "?")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter answer: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("Your answer is: ", text)
	your_answer, _ := strconv.Atoi(text)
	var correct_answer = number1 + number2
	switch your_answer == correct_answer {
	case true:
		fmt.Println("You are correct")
	case false:
		fmt.Println("You are wrong, the correct answer is:", correct_answer)
	}
}
