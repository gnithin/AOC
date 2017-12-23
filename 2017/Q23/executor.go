package main

import (
	"strconv"
)

const (
	INSTN_SET = "set"
	INSTN_MUL = "mul"
	INSTN_JNZ = "jnz"
	INSTN_SUB = "sub"
)

// Holds the instruction details
type Instruction struct {
	cmdName  string
	register string
	argument interface{}
}

// Interprets the instruction set
type Interpreter struct {
	instnList  []Instruction
	varEnv     map[string]int
	instnIndex int
	mulFreq    int
}

func createInterpreter(instnList []Instruction) Interpreter {
	varEnv := make(map[string]int)
	return Interpreter{
		instnList:  instnList,
		varEnv:     varEnv,
		instnIndex: 0,
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
		intVal, err := strconv.Atoi(tempArgVal)
		if err == nil {
			argVal = intVal
		} else {
			argVal = self.getValOfRegister(tempArgVal)
		}
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
	case INSTN_SUB:
		oldVal := self.getValOfRegister(register)
		newVal := oldVal - argVal
		self.setValToRegister(register, newVal)
		self.instnIndex += 1

	case INSTN_MUL:
		oldVal := self.getValOfRegister(register)
		newVal := oldVal * argVal
		self.setValToRegister(register, newVal)
		self.mulFreq += 1
		self.instnIndex += 1

	case INSTN_JNZ:
		regVal := 0
		val, err := strconv.Atoi(register)
		if err == nil {
			regVal = val
		} else {
			regVal = self.getValOfRegister(register)
		}

		if regVal != 0 {
			self.instnIndex += argVal
		} else {
			self.instnIndex += 1
		}

	default:
		panic("Unknown instruction!")
	}

	return false
}
