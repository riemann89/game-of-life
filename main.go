package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var currentState State
	//var newState State

	states := [2]bool{true, false}
	//nState := 10
	size := 10
	currentState = make([][]bool, size)
	for i:= 0; i< size; i++{
		currentState[i] = make([]bool, size)
	}

	//init a random state
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			currentState[i][j] = states[rand.Intn(2)]
		}
	}

	for i:= 0; i < size; i++{
		fmt.Println(currentState[i])
	}



}
/*
	for i := 0; i < nState; i++ {
		for j := 0; j < size; j++ {

		}
	}
*/
