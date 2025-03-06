package main

import "fmt"

func main() {
	// Generate a random number between 1 and 100
	randNum := generateRandomNumber(1, 100)
	fmt.Println("The random number is:", randNum)
}

func generateRandomNumber(min int, max int) int {
	// Return a random number between min and max
	return min + (max-min)*rand.Float64()
}
