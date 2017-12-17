package main

import (
	"fmt"
)

type SpinLock struct {
	stepCount  int
	buffer     []int
	countLimit int
	currPos    int
}

func (self *SpinLock) performLock() int {
	if len(self.buffer) == 0 {
		self.buffer = append(self.buffer, 0)
	}

	for currCounter := 1; currCounter <= self.countLimit; currCounter++ {
		newPos := (self.currPos+self.stepCount)%len(self.buffer) + 1
		self.insertAtPos(currCounter, newPos)
		self.currPos = newPos

		//fmt.Println("******")
		//fmt.Println("Curr count - ", self.currPos)
		//fmt.Println(self.buffer)
	}
	return self.buffer[(self.currPos+1)%len(self.buffer)]
}
func (self *SpinLock) findValAfterElement(element int) int {
	index := self.findValPos(0)
	return self.buffer[(index+1)%len(self.buffer)]
}

func (self *SpinLock) insertAtPos(val, pos int) {
	self.buffer = append(self.buffer, 0)
	copy(self.buffer[pos+1:], self.buffer[pos:])
	self.buffer[pos] = val
}

func (self *SpinLock) findValPos(needle int) int {
	for pos, val := range self.buffer {
		if needle == val {
			return pos
		}
	}
	return -1

}

// Main
func main() {
	stepper := 344
	//stepper := 3

	limit := 2017

	s := SpinLock{
		stepCount:  stepper,
		countLimit: limit,
	}
	finalVal := s.performLock()
	fmt.Println("Final value - ", finalVal)

	// Part 2
	/*
		p2Limit := 50000000
		p2SpinLock := SpinLock{
			stepCount:  stepper,
			countLimit: p2Limit,
		}
		p2SpinLock.performLock()
		fmt.Println("Value after 0 ", p2SpinLock.findValAfterElement(0))
	*/
}
