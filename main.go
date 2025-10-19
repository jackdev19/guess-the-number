package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Ask user to guess a numnber
	fmt.Println("Guess the number between 1-10!")

	// Generate random number and store inside a variable
	number := rand.Intn(10) + 1

	// Declare variable which will store the users guess
	var guess int

	// Declared boolean variable for loop control
	var correctNumber bool = false

	// Loop to check the user guess and compare to the random number
	// Checks against strings, throws error if a string is entered.
	for !correctNumber {
		fmt.Println("Enter your guess: ")
		n, err := fmt.Scanln(&guess)

		if err != nil || n != 1 {
			fmt.Println("Invalid Input. Enter a number!")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if guess == number {
			fmt.Println("Correct! 🎉")
			correctNumber = true
		} else if guess > number {
			fmt.Println("Too high, guess again!")
		} else {
			fmt.Println("Too low, guess again!")
		}
	}

	fmt.Println("Thnx4Playin")
}
