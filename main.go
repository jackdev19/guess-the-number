package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Ask user to guess a numnber
	fmt.Println("Guess the number between 1-10!")

	fmt.Println("Choose a difficulty level \n 1. Easy (10 Chances) \n 2. Medium (5 Chances) \n 3. Hard (3 Chances)")
	fmt.Println("Enter your choice: ")

	var difficulty int

	n, err := fmt.Scanln(&difficulty)

	if err != nil || n != 1 {
		fmt.Println("Invalid Input")
		return
	}

	var maxGuesses int
	switch difficulty {
	case 1:
		maxGuesses = 10
		fmt.Printf("You chose the Easy level. You have %d guesses \n", maxGuesses)
	case 2:
		maxGuesses = 5
		fmt.Printf("You chose the Medium level. You have %d guesses \n", maxGuesses)
	case 3:
		maxGuesses = 3
		fmt.Printf("You chose the Hard level. You have %d guesses \n", maxGuesses)
	default:
		fmt.Println("Invalid difficulty! Defaulting to easy.")
		maxGuesses = 10
	}

	// Generate random number and store inside a variable
	number := rand.Intn(10) + 1

	// Declare variable which will store the users guess and guess amount
	var guess, guessAmount int

	// Declared boolean variable for loop control
	var correctNumber bool = false

	// Loop to check the user guess and compare to the random number
	// Checks against strings, throws error if a string is entered.
	for !correctNumber {
		if guessAmount >= maxGuesses {
			fmt.Printf("You are out of guesses you lose! The correct number was %d. ", number)
			break
		}
		fmt.Println("Enter your guess: ")
		n, err := fmt.Scanln(&guess)
		guessAmount++

		if err != nil || n != 1 {
			fmt.Println("Invalid Input. Enter a number!")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if guess == number {
			fmt.Println("Correct! 🎉")
			correctNumber = true
			fmt.Printf("It took you %d guesses to guess the correct number! \n", guessAmount)
		} else if guess > number {
			fmt.Println("Too high, guess again!")
		} else {
			fmt.Println("Too low, guess again!")
		}
	}

	fmt.Println("Thnx4Playin")
}
