package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var name string
	var a = rand.Intn(2)
	fmt.Println("Welcome to rock, paper, scissor!")
	options := [3]string{"rock", "paper", "scissor"}
	fmt.Println("The options are:", options)
	fmt.Printf("Please choose an option: ")
	fmt.Scanf("%s", &name)
	fmt.Println("You chose: ", name)
	convert(name, options[:])
}

func convert(name string, options []string) {
	switch name {
	case options[0]:
		c := 0
	case options[1]:
		c := 1
	case options[2]:
		c := 2
	default:
		fmt.Println("Not a valid option")
	}
}
