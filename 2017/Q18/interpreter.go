package main

import (
	"fmt"
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
	cmdName  string
	register string
	argument interface{}
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
			//self.responseChan <- self.varEnv[REG_PROG_ID]
			break
		}
	}
}

func (self *Interpreter) runInstn() bool {
	currInstn := self.instnList[self.instnIndex]
	register := currInstn.register

	argVal := 0
	switch tempArgVal := currInstn.argument.(type) {
	case int:
		argVal = tempArgVal
	case string:
		argVal = self.getValOfRegister(tempArgVal)
		//register = tempArgVal
	case nil:
		// Pass
	default:
		panic("Unknown type for argument")
	}

	//fmt.Println(self.progId, currInstn.cmdName, register, currInstn.argument, argVal)
	//fmt.Println(self.progId, currInstn.cmdName, currInstn.register, currInstn.argument)

	// Add the switch cases here
	switch currInstn.cmdName {
	case INSTN_SET:
		self.setValToRegister(register, argVal)
		self.instnIndex += 1

	case INSTN_ADD:
		oldVal := self.getValOfRegister(register)
		newVal := oldVal + argVal
		self.setValToRegister(register, newVal)
		self.instnIndex += 1

	case INSTN_MUL:
		oldVal := self.getValOfRegister(register)
		newVal := oldVal * argVal
		self.setValToRegister(register, newVal)
		self.instnIndex += 1

	case INSTN_MOD:
		oldVal := self.getValOfRegister(register)
		//fmt.Println(oldVal, argVal)
		newVal := oldVal % argVal
		self.setValToRegister(register, newVal)
		self.instnIndex += 1

	case INSTN_SND:
		self.sendCount += 1
		//fmt.Println(self.progId, "sending - ", argVal)
		self.writingChan <- argVal
		self.instnIndex += 1

	case INSTN_RCV:
		if self.progId == 1 {
			fmt.Println("Reg val - ", self.sendCount)
		}
		newVal := <-self.readingChan
		//fmt.Println(self.progId, "Received - ", newVal)
		self.setValToRegister(register, newVal)
		self.instnIndex += 1

	case INSTN_JGZ:
		regVal := self.getValOfRegister(register)
		if regVal > 0 {
			self.instnIndex += argVal
		} else {
			self.instnIndex += 1
		}

	default:
		panic("Unknown instruction!")
	}

	return false
}
