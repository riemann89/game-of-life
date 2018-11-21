package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var currentState State
	//var newState State

	states := [2]uint8{0, 1}
	//nState := 10
	size := 10
	currentState = make([][]uint8, size)
	for i:= 0; i< size; i++{
		currentState[i] = make([]uint8, size)
	}

	//init a random state
	//TODO FIX PSEUDORANDOM	
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			currentState[i][j] = states[rand.Intn(2)]
		}
	}

	for k:=0; k<=2; k++{
		for i:= 0; i < size; i++{
			fmt.Println(currentState[i])
		}
		fmt.Println("")
		time.Sleep(time.Second)
	}


}
