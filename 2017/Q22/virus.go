package main

type Direction int

const (
	N Direction = iota
	S
	E
	W
)

type Virus struct {
	currPos       Position
	currDirection Direction
	grid          *Grid
	numInfected   int
}

func createVirusWithGrid(grid *Grid) Virus {
	initialPos := Position{
		0, 0,
	}
	initialDirn := N

	return Virus{
		currPos:       initialPos,
		currDirection: initialDirn,
		grid:          grid,
		numInfected:   0,
	}
}

func (self *Virus) infectWithBurstSize(burstSize int) {
	for i := 0; i < burstSize; i++ {
		self.infect()
	}
}

func (self *Virus) infect() {
	nodeState := self.grid.GetNodeStateForPos(self.currPos)
	switch nodeState {
	case NODE_CLEAN:
		self.turnLeft()
		self.grid.WeakenNodeAtPos(self.currPos)
	case NODE_FLAGGED:
		self.turnReverse()
		self.grid.CleanNodeAtPos(self.currPos)

	case NODE_WEAKENED:
		// no turn
		self.grid.InfectNodeAtPos(self.currPos)
		self.numInfected += 1
	case NODE_INFECTED:
		self.turnRight()
		self.grid.FlagNodeAtPos(self.currPos)
	}
	self.move()
	//fmt.Println((*self.grid))
	//fmt.Println("****")
}
func (self *Virus) turnReverse() {
	switch self.currDirection {
	case N:
		self.currDirection = S
	case S:
		self.currDirection = N
	case E:
		self.currDirection = W
	case W:
		self.currDirection = E
	}
}

func (self *Virus) turnRight() {
	switch self.currDirection {
	case N:
		self.currDirection = E
	case S:
		self.currDirection = W
	case E:
		self.currDirection = S
	case W:
		self.currDirection = N
	}
}

func (self *Virus) turnLeft() {
	switch self.currDirection {
	case N:
		self.currDirection = W
	case S:
		self.currDirection = E
	case E:
		self.currDirection = N
	case W:
		self.currDirection = S
	}
}

// Move in the current direction
func (self *Virus) move() {
	xinc, yinc := 0, 0
	switch self.currDirection {
	case N:
		xinc, yinc = 1, 0
	case S:
		xinc, yinc = -1, 0
	case E:
		xinc, yinc = 0, 1
	case W:
		xinc, yinc = 0, -1
	}

	self.currPos.row += xinc
	self.currPos.col += yinc
}
