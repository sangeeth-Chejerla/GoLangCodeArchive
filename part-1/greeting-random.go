package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define a slice of slices of strings named greetings.
// Each inner slice contains two strings: a greeting and the language it's in.
var greetings = [][]string{
	{"Hello, World!", "English"},
	{"Salut Monde", "French"},
	{"世界您好", "Simplified Chinese"},
	{"qo' vIvan", "Klingon"},
	{"हैलो वर्ल्ड", "Hindi"},
	{"안녕하세요", "Korean"},
	{"привет мир", "Russian"},
	{"Wapendwa Dunia", "Swahili"},
	{"Hola Mundo", "Spanish"},
	{"Merhaba Dünya", "Turkish"},
}

// Define a function named greeting that returns a slice of strings.
func greeting() []string {
	// Generate a seed for the random number generator using the current Unix time in nanoseconds.
	seed := time.Now().UnixNano()
	// Create a new random number generator using the seed.
	rnd := rand.New(rand.NewSource(seed))
	// Return a random greeting from the greetings slice.
	return greetings[rnd.Intn(len(greetings))]
}

// The main function.
func main() {
	// Call the greeting function and store the returned slice in g.
	g := greeting()
	// Print the greeting and the language it's in.
	fmt.Printf("%s (%s)\n", g[0], g[1])
}
