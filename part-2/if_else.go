package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Currency struct represents a currency with its name, country, and number.
type Currency struct {
	Name    string
	Country string
	Number  int
}

// Define currency variables
var (
	CAD = Currency{Name: "Canadian Dollar", Country: "Canada", Number: 124}
	INR = Currency{Name: "Indian Rupee", Country: "India", Number: 356}
	JPY = Currency{Name: "Japanese Yen", Country: "Japan", Number: 392}
	USD = Currency{Name: "US Dollar", Country: "USA", Number: 840}
)

// main function
func main() {
	// Seed random number generator
	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(seed))
	// Generate a random number between 0 and 999
	num := rnd.Intn(1000)
	fmt.Println("Generated random number:", num)

	// Check if the generated number is within the valid currency number range
	if num < 100 || num > 900 {
		fmt.Println("Invalid currency number")
	} else {
		// Try to find the currency corresponding to the generated number
		if curr, err := findCurr(num); err == nil {
			fmt.Println("Currency found:", curr)
		} else {
			fmt.Println("Currency code", num, "not found")
		}
	}
}

// findCurr function searches for a currency by its number.
func findCurr(number int) (Currency, error) {
	// Check each currency's number to find a match
	switch number {
	case CAD.Number:
		return CAD, nil
	case INR.Number:
		return INR, nil
	case JPY.Number:
		return JPY, nil
	case USD.Number:
		return USD, nil
	default:
		// If no match is found, return an error
		return Currency{}, errors.New("Currency not found")
	}
}
