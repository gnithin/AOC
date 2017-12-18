package main

import (
//"fmt"
)

const (
	INSTN_SET = "set"
	INSTN_ADD = "add"
	INSTN_MUL = "mul"
	INSTN_MOD = "mod"
	INSTN_SND = "snd"
	INSTN_RCV = "rcv"
	INSTN_JGZ = "jgz"
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
	lastSoundFreq int
	recoveredFreq int
	varEnv        map[string]int
	instnIndex    int
}

func createInterpreter(instnList []Instruction) Interpreter {
	varEnv := make(map[string]int)
	return Interpreter{
		instnList:     instnList,
		lastSoundFreq: -1,
		recoveredFreq: -1,
		varEnv:        varEnv,
		instnIndex:    0,
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
		if shouldStop {
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
	case nil:
		// Pass
	default:
		panic("Unknown type for argument")
	}

	//fmt.Println(currInstn.cmdName, register, currInstn.argument, argVal)

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
		newVal := oldVal % argVal
		self.setValToRegister(register, newVal)
		self.instnIndex += 1

	case INSTN_SND:
		self.lastSoundFreq = argVal
		self.instnIndex += 1

	case INSTN_RCV:
		if argVal != 0 {
			self.recoveredFreq = self.lastSoundFreq
			return true
		}
		self.instnIndex += 1

	case INSTN_JGZ:
		regVal := self.getValOfRegister(register)
		if regVal == 0 {
			self.instnIndex += 1
		} else {
			self.instnIndex += argVal
		}

	default:
		panic("Unknown instruction!")
	}

	return false
}
