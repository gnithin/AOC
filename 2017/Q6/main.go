package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func findNumStepsForBlocksBeforeRepeat(originalBlocks []int) int {
	blocksLen := len(originalBlocks)
	blocks := make([]int, blocksLen)
	copy(blocks, originalBlocks)
	history := [][]int{}
	runCount := 0

	for {
		blocksCopy := make([]int, blocksLen)
		copy(blocksCopy, blocks)
		history = append(history, blocksCopy)

		maxIndex := getIndexOfMaxInList(blocks)
		maxVal := blocks[maxIndex]
		blocks[maxIndex] = 0
		currIndex := (maxIndex + 1) % blocksLen

		for i := maxVal; i > 0; i-- {
			blocks[currIndex] += 1
			currIndex = (currIndex + 1) % blocksLen
		}

		runCount += 1

		if listContainsBlock(history, blocks) {
			break
		}
	}
	return runCount
}

func listContainsBlock(list [][]int, block []int) bool {
	//fmt.Println("*********")
	//fmt.Println(block)
	//fmt.Println(list)

	for _, currBlock := range list {
		if len(currBlock) != len(block) {
			continue
		}

		isEqual := true
		for i := 0; i < len(currBlock); i++ {
			if currBlock[i] != block[i] {
				isEqual = false
				break
			}
		}

		if isEqual {
			return true
		}
	}
	return false

}

func getIndexOfMaxInList(list []int) int {
	if len(list) == 0 {
		return -1
	}

	maxIndex := -1
	maxVal := -1

	for index, val := range list {
		if val > maxVal {
			maxVal = val
			maxIndex = index
		}
	}

	return maxIndex
}

func main() {
	filename := "ip.txt"
	blocks := getIpListFromFilename(filename)
	fmt.Println("Input blocks - ", blocks)
	fmt.Println("Number of steps before recursion - ", findNumStepsForBlocksBeforeRepeat(blocks))
}

func getIpListFromFilename(filename string) []int {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\t")

	ipList := []int{}

	for _, ipByte := range ipStrList {
		ipStr := string(ipByte)
		intVal, strConvErr := strconv.Atoi(ipStr)
		if strConvErr != nil {
			panic(strConvErr)
		}
		ipList = append(ipList, intVal)
	}

	return ipList
}
