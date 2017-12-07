package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var programStash map[string]*ProgramData = make(map[string]*ProgramData)
var baseProgramFilterList []string

type ProgramData struct {
	name   string
	weight int
	onTop  []*ProgramData
}

func (p *ProgramData) populateFromString(ipStr string) {
	parentChildSplit := strings.Split(ipStr, "->")
	parentStr := strings.TrimSpace(parentChildSplit[0])
	nameWeightSplit := strings.Split(parentStr, " ")
	p.name = strings.TrimSpace(nameWeightSplit[0])
	weightStr := strings.Trim(nameWeightSplit[1], "() ")
	weight, err := strconv.Atoi(weightStr)
	if err != nil {
		panic(err)
	}
	p.weight = weight
}

func (p *ProgramData) getCorrectedWeight() int {
	var weightDistribution map[int][]int = make(map[int][]int)

	for topIndex, childProgram := range p.onTop {
		weight := childProgram.getWeight()
		_, isFound := weightDistribution[weight]
		if !isFound {
			weightDistribution[weight] = []int{topIndex}
		} else {
			weightDistribution[weight] = append(weightDistribution[weight], topIndex)
		}
	}

	commonWeight := -1
	outlierWeight := -1
	for weight, indices := range weightDistribution {
		if len(indices) == 1 {
			outlierWeight = weight
		} else {
			commonWeight = weight
		}
	}

	weightDiff := outlierWeight - commonWeight

	unequalProgram := p.getUnEqualProgram()
	return unequalProgram.weight - weightDiff
}

func (p *ProgramData) getUnEqualProgram() *ProgramData {
	if len(p.onTop) == 0 {
		return p
	}

	var weightDistribution map[int][]int = make(map[int][]int)

	for topIndex, childProgram := range p.onTop {
		weight := childProgram.getWeight()
		_, isFound := weightDistribution[weight]
		if !isFound {
			weightDistribution[weight] = []int{topIndex}
		} else {
			weightDistribution[weight] = append(weightDistribution[weight], topIndex)
		}
	}

	outlierIndex := -1
	for _, indices := range weightDistribution {
		if len(indices) == 1 {
			outlierIndex = indices[0]
		}
	}

	if outlierIndex == -1 {
		return p
	}

	unEqualProgram := p.onTop[outlierIndex]

	return unEqualProgram.getUnEqualProgram()

}

func (p *ProgramData) getWeight() int {
	weight := p.weight
	for _, childProgram := range p.onTop {
		weight += childProgram.getWeight()
	}

	return weight
}

// Part 1

func getBaseFromIpList(ipList []string) *ProgramData {
	sep := "->"

	// First pass
	for _, ip := range ipList {
		program := ProgramData{}
		program.populateFromString(ip)
		programStash[program.name] = &program
		baseProgramFilterList = append(baseProgramFilterList, program.name)
	}

	// Second pass
	for _, ip := range ipList {
		parentChildSplit := strings.Split(ip, sep)

		parentNameStr := strings.Split(parentChildSplit[0], " ")[0]
		if !strings.Contains(ip, sep) {
			removeProgramFromBaseList(parentNameStr)
			continue
		}

		parentProgram, parentFound := programStash[parentNameStr]
		if !parentFound {
			panic(fmt.Sprintf("Could not find parent - '%s'", parentNameStr))
		}

		childListStr := parentChildSplit[1]
		childsList := strings.Split(childListStr, ",")
		for _, childStr := range childsList {
			childStr = strings.TrimSpace(childStr)
			childProgram, childFound := programStash[childStr]
			if !childFound {
				panic("Could not find children!")
				panic(fmt.Sprintln("Could not find children - ", childStr))
			}
			parentProgram.onTop = append(parentProgram.onTop, childProgram)
			removeProgramFromBaseList(childStr)
		}
	}

	baseName := baseProgramFilterList[0]
	removeProgramFromBaseList(baseName)
	baseProgram := programStash[baseName]

	return baseProgram
}

func removeProgramFromBaseList(parentNameStr string) {
	parentIndex := -1
	for index, name := range baseProgramFilterList {
		if name == parentNameStr {
			parentIndex = index
			break
		}
	}

	if parentIndex == -1 {
		return
	}

	baseProgramFilterList = append(baseProgramFilterList[:parentIndex], baseProgramFilterList[parentIndex+1:]...)
}

// Part 2
func findUnequalWeighForBaseProgram(baseProgram *ProgramData) int {
	return baseProgram.getCorrectedWeight()
}

// Main
func main() {
	filename := "ip.txt"
	ipList := getIpListFromFilename(filename)
	baseProgram := getBaseFromIpList(ipList)
	fmt.Println("Base name - ", baseProgram.name)

	unequalWeight := findUnequalWeighForBaseProgram(baseProgram)
	fmt.Println("Unequal weight - ", unequalWeight)

}

func getIpListFromFilename(filename string) []string {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")
	return ipStrList
}
