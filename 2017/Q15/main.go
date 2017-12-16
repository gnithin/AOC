package main

import (
	"fmt"
	"strconv"
)

type Generator struct {
	currVal     int64
	factor      int64
	maxVal      int64
	pickyFactor int64
}

func (g *Generator) generateNextVal() int64 {
	g.currVal = (g.currVal * g.factor) % g.maxVal
	return g.currVal
}

func (g *Generator) generateNextValPicky() int64 {
	for {
		g.currVal = (g.currVal * g.factor) % g.maxVal
		if g.currVal%g.pickyFactor == 0 {
			break
		}
	}
	return g.currVal
}

func (g *Generator) getBaseBinStrOfVal(numDigits int64) string {
	binaryRep := strconv.FormatInt(g.currVal, 2)
	binaryRepLen := int64(len(binaryRep))
	if binaryRepLen < numDigits {
		for i := int64(0); i < (numDigits - binaryRepLen); i++ {
			binaryRep = "0" + binaryRep
		}
	} else if binaryRepLen > numDigits {
		binaryRep = binaryRep[binaryRepLen-numDigits:]
	}

	return binaryRep
}

// Part 1
func getNumMatches(genA, genB Generator, genPicky bool, limit int64) int64 {
	numMatches := int64(0)
	for i := int64(0); i < limit; i++ {
		if !genPicky {
			genA.generateNextVal()
			genB.generateNextVal()
		} else {
			genA.generateNextValPicky()
			genB.generateNextValPicky()
		}

		aStr := genA.getBaseBinStrOfVal(16)
		bStr := genB.getBaseBinStrOfVal(16)

		if aStr == bStr {
			numMatches += 1
		}
	}
	return numMatches
}

// Main
const MAX_VAL = 2147483647

func main() {
	//genAInit := int64(65)
	//genBInit := int64(8921)
	genAInit := int64(722)
	genBInit := int64(354)

	genA := Generator{
		currVal: genAInit,
		factor:  int64(16807),
		maxVal:  MAX_VAL,
	}

	genB := Generator{
		currVal: genBInit,
		factor:  int64(48271),
		maxVal:  MAX_VAL,
	}

	p1Limit := int64(40 * 1000000)
	numMatches := getNumMatches(genA, genB, false, p1Limit)
	fmt.Println("P1 Num matches - ", numMatches)

	p2Limit := int64(5 * 1000000)
	genA.currVal = genAInit
	genA.pickyFactor = int64(4)
	genB.currVal = genBInit
	genB.pickyFactor = int64(8)
	numP2Matches := getNumMatches(genA, genB, true, p2Limit)
	fmt.Println("P2 Num matches - ", numP2Matches)
}
