package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var name string
	fmt.Println("Welcome to rock, paper, scissor!")
	options := [3]string{"rock", "paper", "scissor"}
	fmt.Println("The options are:", options)
	fmt.Printf("Please choose an option: ")
	fmt.Scanf("%s", &name)
	fmt.Println("You chose: ", name)
	switch name {
	case options[0]:
		c := 0
		solve(c)
	case options[1]:
		c := 1
		solve(c)
	case options[2]:
		c := 2
		solve(c)
	default:
		fmt.Println("Not a valid option")
	}
}

func solve(c int) {
	rand.Seed(time.Now().UnixNano())
	var a = rand.Intn(2)
	switch {
	case c == a:
		fmt.Println("Tie")
	default:
		switch c == 0 {
		case a == 2:
			fmt.Println("Bot chooses", a)
			fmt.Println("You win")
		case a == 1:
			fmt.Println("Bot chooses", a)
			fmt.Println("You lose")
		}
		switch c == 1 {
		case a == 0:
			fmt.Println("Bot chooses", a)
			fmt.Println("You win")
		case a == 2:
			fmt.Println("Bot chooses", a)
			fmt.Println("You lose")
		}
		switch c == 2 {
		case a == 1:
			fmt.Println("Bot chooses", a)
			fmt.Println("You win")
		case a == 0:
			fmt.Println("Bot chooses", a)
			fmt.Println("You lose")
		}
	}
}
