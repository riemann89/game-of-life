package main

//State type of a system
type State [][]string

//StateTransition computes the transition between a state into another
func (currentState State) StateTransition(size int) State {

	newState := State(make([][]string, size))
	
	for i:=0; i< size; i++{
		newState[i] = make([]string, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			lifeAround := currentState.countLiving(size, i, j)
			if currentState[i][j] == "*" {
				switch {
				case lifeAround < 2:
					newState[i][j] = " "
				case lifeAround == 2 || lifeAround == 3:
					newState[i][j] = "*"
				default:
					newState[i][j] = " "
				}
			} else {
				if lifeAround == 3 {
					newState[i][j] = "*"
				} else {
					newState[i][j] = " "
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
			if (x+i) == x && (y+j) == y {
				continue
			}
			//check if the neighbour is in the square
			if 0 <= (x+i) && (x+i) < size && 0 <= (y+j) && (y+j) < size {
				if currentState[x+i][y+j] == "*" {
					totalLiving++
				}
			}
		}
	}

	return totalLiving
}


//BlockState initilize a still life block example
func BlockState(size int) (State){
	
	Block := State(make([][]string, size))
	
	for i:=0; i< size; i++{
		Block[i] = make([]string, size)
	}

	//init stable status
	Block[0][0] = " "
	Block[0][1] = " "
	Block[0][2] = " "
	Block[0][3] = " "

	Block[1][0] = " "
	Block[1][1] = "*"
	Block[1][2] = "*"
	Block[1][3] = " "

	Block[2][0] = " "
	Block[2][1] = "*"
	Block[2][2] = "*"
	Block[2][3] = " "

	Block[3][0] = " "
	Block[3][1] = " "
	Block[3][2] = " "
	Block[3][3] = " "

	return Block
}