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
	var b = options[a]
	switch name {
	case b:
		fmt.Println("Tie")
	}
}
