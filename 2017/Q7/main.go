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

func main() {
	filename := "ip.txt"
	ipList := getIpListFromFilename(filename)
	baseProgram := getBaseFromIpList(ipList)
	fmt.Println("Base name - ", baseProgram.name)
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
