package main

//State type of a system
type State [][]bool

//StateTransition computes the transition between a state into another
func (currentState State) StateTransition(size int) State {

	newState := State(make([][]bool, size))

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			lifeAround := currentState.countLiving(size, i, j)
			if currentState[i][j] == true {
				switch {
				case lifeAround < 2:
					newState[i][j] = false
				case lifeAround == 2 || lifeAround == 3:
					newState[i][j] = true
				default:
					newState[i][j] = false
				}
			} else {
				if lifeAround == 3 {
					newState[i][j] = true
				}
			}
		}
	}

	return newState
}

func (currentState State) countLiving(size, x, y int) uint8 {

	totalLiving := uint8(0)

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if (x+i) == 0 && (y+j) == 0 {
				continue
			}
			//check if the neighbour is in the square
			if 0 <= (x+i) && (x+i) < size && 0 <= (y+j) && (y+j) < size {
				if currentState[x+i][y+j] == true {
					totalLiving++
				}
			}
		}
	}

	return totalLiving
}
