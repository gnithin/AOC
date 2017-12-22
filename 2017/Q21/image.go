package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	ON_PIXEL  = "#"
	OFF_PIXEL = "."
	ROW_SEP   = "/"
)

type Image struct {
	imgPixels           [][]string
	enhancementRulesMap map[string]string
}

func CreateImage(initialImg [][]string, enhancementRules map[string]string) Image {
	return Image{
		initialImg,
		enhancementRules,
	}
}

func (self *Image) getSize() int {
	return len(self.imgPixels)

}

func (self Image) String() string {
	op := fmt.Sprintln("*** Size - ", self.getSize(), "***")
	for row := 0; row < self.getSize(); row++ {
		for col := 0; col < self.getSize(); col++ {
			op += fmt.Sprintf("%-3s", self.imgPixels[row][col])
		}
		if row != (self.getSize() - 1) {
			op += fmt.Sprintf("\n")
		}
	}
	op += fmt.Sprintln("\n*****************")
	return op
}

// Part 1
func (self *Image) getNumOnPixels() int {
	numOnPixels := 0
	for row := 0; row < self.getSize(); row++ {
		for col := 0; col < self.getSize(); col++ {
			if self.imgPixels[row][col] == ON_PIXEL {
				numOnPixels += 1
			}
		}
	}
	return numOnPixels
}

func (self *Image) iterateNTimes(n int) {
	//fmt.Println(self)
	for i := 0; i < n; i++ {
		self.iterate()
		//fmt.Println(i, " done!")
		//fmt.Println(self)
	}
}

func (self *Image) iterate() {
	breakDownSize := 3
	if self.getSize()%2 == 0 {
		breakDownSize = 2
	}

	var newImageSegments []string
	segments := self.getSegmentStrOfSize(breakDownSize)

	// For each segment find the rule that matches it!
	for _, s := range segments {
		possibleRules := self.getAllPossibleRulesForSegment(s)
		/*
			fmt.Println("Possible rules for ", s)
			for _, r := range possibleRules {
				fmt.Println(r)
			}
		*/
		matchFound := false
		newSegmentStr := ""
		for _, rule := range possibleRules {
			val, ok := self.enhancementRulesMap[rule]
			if ok == true {
				matchFound = true
				newSegmentStr = val
				break
			}
		}

		if matchFound == false {
			panic("No rules match found!")
		}

		newImageSegments = append(newImageSegments, newSegmentStr)
	}

	self.updateImageWithSegments(newImageSegments)
}

func (self *Image) updateImageWithSegments(newImageSegments []string) {
	numSegments := len(newImageSegments)
	sqrtFloatVal := math.Sqrt(float64(numSegments))
	if sqrtFloatVal-math.Floor(sqrtFloatVal) != 0 {
		panic("Num segments is not a perfect square!")
	}
	numSegmentsInARow := int(sqrtFloatVal)

	var newImagesList [][]string

	for i := 0; i < len(newImageSegments); i = i + numSegmentsInARow {
		rowSegmentsList := newImageSegments[i : i+numSegmentsInARow]
		compLen := len(strings.Split(rowSegmentsList[0], ROW_SEP)[0])

		rowCompList := make([]string, compLen)
		for _, rowVal := range rowSegmentsList {
			segComponents := strings.Split(rowVal, ROW_SEP)
			for currIndex, cVal := range segComponents {
				rowCompList[currIndex] += cVal
			}
		}
		for _, row := range rowCompList {
			var rowRunes []string
			for _, c := range row {
				rowRunes = append(rowRunes, string(c))
			}
			newImagesList = append(newImagesList, rowRunes)
		}
	}
	self.imgPixels = newImagesList
}

func (self *Image) getAllPossibleRulesForSegment(s string) []string {
	currSegment := s
	var newSegments []string

	// Perform rotations(3)
	for i := 0; i < 4; i++ {
		currSegment = self.getRotatedSegment(currSegment)
		newSegments = append(newSegments, currSegment)
		mirrorSegment := self.getMirrorSegment(currSegment)
		newSegments = append(newSegments, mirrorSegment)
	}
	return newSegments
}

func (self *Image) getRotatedSegment(s string) string {
	var compList []string
	rowsList := strings.Split(s, ROW_SEP)
	size := len(rowsList[0])
	for i := 0; i < size; i++ {
		newRow := ""
		for j := size - 1; j >= 0; j-- {
			newRow += string(rowsList[j][i])
		}
		compList = append(compList, newRow)
	}
	rotatedSegment := strings.Join(compList, ROW_SEP)
	return rotatedSegment
}

func (self *Image) getMirrorSegment(s string) string {
	// Split the original
	var finalStr []string
	for _, c := range strings.Split(s, ROW_SEP) {
		revC := self.getReversedString(c)
		finalStr = append(finalStr, revC)
	}
	return strings.Join(finalStr, ROW_SEP)
}

func (self *Image) getReversedString(s string) string {
	rList := []rune(s)
	for i, j := 0, len(rList)-1; i < (len(rList) / 2); i, j = i+1, j-1 {
		rList[i], rList[j] = rList[j], rList[i]
	}
	return string(rList)
}

func (self *Image) getSegmentStrOfSize(breakDownSize int) []string {
	var segments []string
	size := self.getSize()

	// Note that this creates the segments horizontally first,
	// then moves vertically one step, then continues horizontally
	for row := 0; row < size; row += breakDownSize {
		for col := 0; col < size; col += breakDownSize {
			var segList []string
			for i := 0; i < breakDownSize; i++ {
				segList = append(segList, strings.Join(self.imgPixels[row+i][col:col+breakDownSize], ""))
			}
			segments = append(segments, strings.Join(segList, ROW_SEP))
		}
	}
	return segments
}
