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
		self.buffer = make([]int, 0, 50000000)
		self.buffer = append(self.buffer, 0)
	}

	for currCounter := 1; currCounter <= self.countLimit; currCounter++ {
		newPos := (self.currPos+self.stepCount)%len(self.buffer) + 1
		self.insertAtPos(currCounter, newPos)
		self.currPos = newPos
	}
	return self.buffer[(self.currPos+1)%len(self.buffer)]
}
func (self *SpinLock) findValAfterInitialElement() int {
	val := -1
	for currCounter := 1; currCounter <= self.countLimit; currCounter++ {
		newPos := (self.currPos+self.stepCount)%(currCounter) + 1
		self.currPos = newPos
		if self.currPos == 1 {
			val = currCounter
		}
	}
	return val
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
	limit := 2017

	s := SpinLock{
		stepCount:  stepper,
		countLimit: limit,
	}
	finalVal := s.performLock()
	fmt.Println("Final value - ", finalVal)

	// Part 2
	p2Limit := 50000000
	p2SpinLock := SpinLock{
		stepCount:  stepper,
		countLimit: p2Limit,
	}
	fmt.Println("Value after the initial element -", p2SpinLock.findValAfterInitialElement())
}
