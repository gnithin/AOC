package main

import (
	"fmt"
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
	cmdName string
	arg1    string
	arg2    string
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
	stepsLimit := 10000000
	steps := 0
	for self.instnIndex >= 0 && self.instnIndex < len(self.instnList) {
		if steps > stepsLimit {
			fmt.Println("Stopping!")
			break
		} else {
			//fmt.Println(steps, "g - ", self.getValOfRegister("g"))
			fmt.Println("Step number - ", self.instnIndex, " -> ", self.instnList[self.instnIndex])
			steps += 1
		}
		shouldStop := self.runInstn()
		if shouldStop {
			fmt.Println("Stopping automatically!")
			break
		}
	}
	fmt.Println(steps)
	fmt.Println(self.instnIndex)
}

func (self *Interpreter) runInstn() bool {
	currInstn := self.instnList[self.instnIndex]
	arg1Register := ""
	arg2Register := ""
	arg1Val := 0
	arg2Val := 0

	intVal1, err := strconv.Atoi(currInstn.arg1)
	if err == nil {
		arg1Val = intVal1
	} else {
		arg1Register = currInstn.arg1
		arg1Val = self.getValOfRegister(arg1Register)
	}

	intVal2, err2 := strconv.Atoi(currInstn.arg2)
	if err2 == nil {
		arg2Val = intVal2
	} else {
		arg2Register = currInstn.arg2
		arg2Val = self.getValOfRegister(arg2Register)
	}

	//fmt.Println(currInstn.cmdName, currInstn.arg1, currInstn.arg2)

	// Add the switch cases here
	switch currInstn.cmdName {
	case INSTN_SET:
		//fmt.Println(arg1Register, " = ", arg2Val)

		self.setValToRegister(arg1Register, arg2Val)
		self.instnIndex += 1
	case INSTN_SUB:
		//fmt.Println(arg1Register, " = ", arg1Val, " - ", arg2Val)

		newVal := arg1Val - arg2Val
		self.setValToRegister(arg1Register, newVal)
		self.instnIndex += 1

	case INSTN_MUL:
		//fmt.Println(arg1Register, " = ", arg1Val, " * ", arg2Val)

		newVal := arg1Val * arg2Val
		self.setValToRegister(arg1Register, newVal)
		self.mulFreq += 1
		self.instnIndex += 1

	case INSTN_JNZ:
		//fmt.Println("***j")
		if arg1Val != 0 {
			self.instnIndex += arg2Val
			//fmt.Println("JUMPING - ", arg2Val)
		} else {
			self.instnIndex += 1
			//fmt.Println("NOT JUMPING")
		}

	default:
		panic("Unknown instruction!")
	}

	return false
}
