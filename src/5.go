package main

import "fmt"

func getRandomNumber(max int) int {
	return max * rand.Intn(100)
}
