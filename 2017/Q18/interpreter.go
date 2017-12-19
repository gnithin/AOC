package main

import (
	"fmt"
	"strconv"
	"unicode"
)

const (
	INSTN_SET   = "set"
	INSTN_ADD   = "add"
	INSTN_MUL   = "mul"
	INSTN_MOD   = "mod"
	INSTN_SND   = "snd"
	INSTN_RCV   = "rcv"
	INSTN_JGZ   = "jgz"
	REG_PROG_ID = "p"
)

// Holds the instruction details
type Instruction struct {
	cmdName string
	arg1    string
	arg2    string
}

// Interprets the instruction set
type Interpreter struct {
	instnList     []Instruction
	varEnv        map[string]int
	instnIndex    int
	readingChan   chan int
	writingChan   chan int
	sendCount     int
	responseChan  chan int
	stopInterrupt bool
	progId        int
}

func createInterpreter(
	progId int,
	instnList []Instruction,
	readingChan,
	writingChan chan int,
	responseChan chan int,
) Interpreter {
	varEnv := make(map[string]int)
	varEnv[REG_PROG_ID] = progId

	return Interpreter{
		progId:        progId,
		instnList:     instnList,
		varEnv:        varEnv,
		instnIndex:    0,
		readingChan:   readingChan,
		writingChan:   writingChan,
		responseChan:  responseChan,
		stopInterrupt: false,
	}
}

func (self *Interpreter) getValOfRegister(regName string) int {
	regVal, _ := self.varEnv[regName]
	return regVal
}

func (self *Interpreter) setValToRegister(regName string, val int) {
	self.varEnv[regName] = val
}

func (self *Interpreter) run() {
	for self.instnIndex >= 0 && self.instnIndex < len(self.instnList) {
		shouldStop := self.runInstn()
		if shouldStop || self.stopInterrupt == true {
			self.responseChan <- self.varEnv[REG_PROG_ID]
			break
		}
	}
}

func (self *Interpreter) runInstn() bool {
	currInstn := self.instnList[self.instnIndex]
	r1 := ""
	r2 := ""
	a1 := 0
	a2 := 0

	if unicode.IsLetter([]rune(currInstn.arg1)[0]) {
		r1 = currInstn.arg1
		a1 = self.getValOfRegister(r1)
	} else {
		a1, _ = strconv.Atoi(currInstn.arg1)
	}

	if currInstn.arg2 != "" {
		if unicode.IsLetter([]rune(currInstn.arg2)[0]) {
			r2 = currInstn.arg2
			a2 = self.getValOfRegister(r2)
		} else {
			a2, _ = strconv.Atoi(currInstn.arg2)
		}
	}

	//fmt.Println(self.progId, currInstn.cmdName, r1, a1, r2, a2)
	//fmt.Println(self.progId, currInstn.cmdName, currInstn.register, currInstn.argument)

	// Add the switch cases here
	switch currInstn.cmdName {
	case INSTN_SET:
		self.setValToRegister(r1, a2)
		self.instnIndex += 1

	case INSTN_ADD:
		oldVal := self.getValOfRegister(r1)
		newVal := oldVal + a2
		self.setValToRegister(r1, newVal)
		self.instnIndex += 1

	case INSTN_MUL:
		oldVal := self.getValOfRegister(r1)
		newVal := oldVal * a2
		self.setValToRegister(r1, newVal)
		self.instnIndex += 1

	case INSTN_MOD:
		oldVal := self.getValOfRegister(r1)
		newVal := oldVal % a2
		self.setValToRegister(r1, newVal)
		self.instnIndex += 1

	case INSTN_SND:
		self.sendCount += 1
		//fmt.Println(self.progId, "sending - ", argVal)
		self.writingChan <- a1
		self.instnIndex += 1

	case INSTN_RCV:
		if self.progId == 1 {
			fmt.Println("Reg val - ", self.sendCount)
		}
		//fmt.Println(self.progId, "Received - ", newVal, "setting register - ", rcvRegister)
		newVal := <-self.readingChan
		self.setValToRegister(r1, newVal)
		self.instnIndex += 1

	case INSTN_JGZ:
		if a1 > 0 {
			self.instnIndex += a2
		} else {
			self.instnIndex += 1
		}

	default:
		panic("Unknown instruction!")
	}

	return false
}
