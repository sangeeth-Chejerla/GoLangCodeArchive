package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	// Generate a random number between 1 and 100
	secretNumber := rnd.Intn(100) + 1

	maxAttempts := 5
	attemptsLeft := maxAttempts
	fmt.Println("Welcome to the Secret Number Game!")
	fmt.Println("I've picked a secret number between 1 and 100.")
	fmt.Printf("You have %d attempts to guess it.\n", maxAttempts)

	// Game loop
	for attemptsLeft > 0 {
		fmt.Printf("\nEnter your number: ")

		var guess int
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("Error reading input. Please enter a valid number .")
			continue
		}

		if guess == secretNumber {
			fmt.Println("Congratulations! You guessed the secret number!")
			return
		} else if guess < secretNumber {
			fmt.Println("Oops! Your guess was LOW.")
		} else {
			fmt.Println("Oops! Your guess was HIGH.")
		}

		attemptsLeft--
		fmt.Printf("You have %d attempts left.\n", attemptsLeft)
	}

	fmt.Println("Sorry, you've run out of attempts.")
	fmt.Printf("The secret number was: %d\n", secretNumber)
}
