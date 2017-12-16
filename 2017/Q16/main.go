package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Part 1

type CmdType int

const (
	S = iota
	X
	P
)

type CmdManager struct {
	cmdList []Cmd
	ipList  []rune
	cache   []string
}

func (manager *CmdManager) execute() bool {
	currConfig := string(manager.ipList)
	cachePos := manager.findPosInCache(currConfig)
	if cachePos != -1 {
		manager.cache = manager.cache[:len(manager.cache)-1]
		nextElemStr := manager.cache[(cachePos+1)%len(manager.cache)]
		var updatedRuneList []rune
		for _, r := range nextElemStr {
			updatedRuneList = append(updatedRuneList, r)
		}
		manager.ipList = updatedRuneList
		return true
	}

	for _, cmd := range manager.cmdList {
		args1 := cmd.args1
		args2 := cmd.args2
		switch cmd.name {
		case S:
			pivot := len(manager.ipList) - cmd.args1
			manager.ipList = append(manager.ipList[pivot:], manager.ipList[:pivot]...)

		case X:
			manager.ipList[args1], manager.ipList[args2] = manager.ipList[args2], manager.ipList[args1]

		case P:
			pos1 := manager.findPos(rune(args1))
			pos2 := manager.findPos(rune(args2))
			manager.ipList[pos1], manager.ipList[pos2] = manager.ipList[pos2], manager.ipList[pos1]
		}
	}

	currList := string(manager.ipList)
	manager.cache = append(manager.cache, currList)

	return false
}

func (manager *CmdManager) findPosInCache(needle string) int {
	if len(manager.cache) == 0 {
		return -1
	}
	for i, entry := range manager.cache[:(len(manager.cache) - 1)] {
		if entry == needle {
			return i
		}
	}
	return -1
}

func (manager *CmdManager) findPos(needle rune) int {
	for i, entry := range manager.ipList {
		if entry == needle {
			return i
		}
	}
	return -1
}

func (manager *CmdManager) printList() {
	fmt.Printf("Iplist : ")
	for _, val := range manager.ipList {
		fmt.Printf("%s", string(val))
	}
	fmt.Println()
}

// Main
type Cmd struct {
	name  CmdType
	args1 int
	args2 int
}

func main() {
	filename := "ip.txt"
	numEntries := 16

	cmdList := getCmdListFromFilename(filename)

	var ipList []rune
	charVal := rune("a"[0])
	for i := 0; i < numEntries; i++ {
		ipList = append(ipList, charVal)
		charVal = (charVal + 1)
	}

	fmt.Println("*****P1*****")
	var p1List = make([]rune, len(ipList))
	copy(p1List, ipList)
	p1Manager := CmdManager{
		cmdList: cmdList,
		ipList:  p1List,
	}
	p1Manager.execute()
	p1Manager.printList()

	// Part 2
	fmt.Println("*****P2*****")
	runCount := int64(1000000000)
	var p2List = make([]rune, len(ipList))
	copy(p2List, ipList)
	p2Manager := CmdManager{
		cmdList: cmdList,
		ipList:  p2List,
	}

	foundCycle := false
	for i := int64(0); i < runCount; i++ {
		foundCycle = p2Manager.execute()
		if foundCycle {
			break
		}
	}

	fmt.Println("Final - runcount - ", runCount)
	if foundCycle {
		index := (runCount - 1) % int64(len(p2Manager.cache))
		fmt.Println(p2Manager.cache[int(index)])
	} else {
		p2Manager.printList()
	}
}

func getCmdListFromFilename(filename string) []Cmd {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, ",")
	var cmdList []Cmd
	for _, ipStr := range ipStrList {
		cmdStr := string(ipStr[0])
		args := ipStr[1:]
		cmdList = append(cmdList, parseCommand(cmdStr, args))
	}
	return cmdList
}

func parseCommand(name, args string) Cmd {
	cmd := Cmd{}
	switch name {
	case "s":
		cmd.name = S
		size, _ := strconv.Atoi(args)
		cmd.args1 = size

	case "x":
		cmd.name = X
		posStrList := strings.Split(args, "/")
		var posList []int
		for _, posStr := range posStrList {
			posInt, _ := strconv.Atoi(string(posStr))
			posList = append(posList, posInt)
		}
		cmd.args1 = posList[0]
		cmd.args2 = posList[1]

	case "p":
		cmd.name = P
		posStrList := strings.Split(args, "/")
		var posList []rune
		for _, posStr := range posStrList {
			posList = append(posList, []rune(posStr)[0])
		}
		cmd.args1 = int(posList[0])
		cmd.args2 = int(posList[1])
	}
	return cmd
}
