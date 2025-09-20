package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// how to always win
// choose cap 5
// you will get 3 try just guess half number

func main() {
	var cap int
	fmt.Print("Please Enter the upper limit of your guessing range: ")
	fmt.Scan(&cap)
	if cap <= 0 {
		fmt.Println("please enter more than 0")
		return
	}
	goal := randInt(cap)
	var guess int
	maximumGuesses := int(math.Ceil(math.Log2(float64(cap))))
	for try := 0; try < maximumGuesses; try++ {
		fmt.Printf("You have %d tries left. Please enter your guess: ", maximumGuesses-try)
		fmt.Scan(&guess)
		switch {
		case guess < goal:
			fmt.Println("Too low!")
		case guess > goal:
			fmt.Println("Too high!")
		default:
			fmt.Println("you are right")
			return
		}
	}
	fmt.Println("you failed (run out of attempt)")
}

func randInt(cap int) int {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	randInstance := rand.New(source)
	goal := randInstance.Intn(cap)
	return goal
}
