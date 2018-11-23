package main

import (
	"fmt"
	"time"
)

func main() {

	var currentState State
	//states := [2]string{" ", "*"}
	size := 4
	currentState = make([][]string, size)
	for i := 0; i < size; i++ {
		currentState[i] = make([]string, size)
	}
	
	//init stable status
	currentState = BlockState(size)

	for k := 0; k <= 10; k++ {
		for i := 0; i < size; i++ {
			fmt.Println(currentState[i])
		}
		fmt.Println("")
		time.Sleep(time.Second)
		currentState = currentState.StateTransition(size)
	}

}
