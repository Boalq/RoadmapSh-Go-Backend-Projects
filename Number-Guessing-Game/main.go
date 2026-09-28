package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Printf("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.\n\n")

	fmt.Println("Please select the difficulty level:")
	fmt.Printf("1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)\n\n")

	var choice int

	fmt.Printf("Enter your choice: ")
	fmt.Scanln(&choice)

	if choice == 0 || choice > 3 {
		fmt.Println("Invalid Choice.")
		return
	}

	fmt.Println()

	var tries int
	switch choice {
	case 1:
		tries = 10
	case 2:
		tries = 5
	case 3:
		tries = 3
	}

	var guess int

	randomNum := rand.IntN(100)

	fmt.Println("Great! You have selected the Medium difficulty level.")
	fmt.Println("Let's start the game!")
	
	// Cheat code for testing!
	//fmt.Println(randomNum)

	for i := 0; i < tries; i++ {
		fmt.Printf("\nEnter your choice:")
		fmt.Scanln(&guess)

		if guess == randomNum {
			fmt.Printf("Congratulations! You guessed the number in %d attempts!\n", i + 1)
			return
		} else if guess < randomNum {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
		} else {
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		}

		if i == tries / 2 {
			fmt.Printf("\nTip: ")
			switch{
			case randomNum < 20: fmt.Printf("The number is less than 20\n")
			case randomNum < 50: fmt.Printf("The number is between 20 - 49\n")
			case randomNum < 70: fmt.Printf("The number is between 50 - 69\n")
			default: fmt.Printf("The number is equal or greater than 70\n")
			}
		}
	}

	fmt.Printf("\nSorry, You used all of your tries! Try again next time.\n")
}
