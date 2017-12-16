package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Part 1
type CmdManager struct {
	cmdList []Cmd
	ipList  []rune
}

func (manager *CmdManager) execute() {
	for _, cmd := range manager.cmdList {
		manager.execCmd(cmd)
	}
}

func (manager *CmdManager) execCmd(cmd Cmd) {
	// Parse the command
	switch cmd.name {
	case "s":
		size, _ := strconv.Atoi(cmd.args)
		manager.cmdSpin(size)

	case "x":
		posStrList := strings.Split(cmd.args, "/")
		var posList []int
		for _, posStr := range posStrList {
			posInt, _ := strconv.Atoi(string(posStr))
			posList = append(posList, posInt)
		}
		manager.cmdExchange(posList[0], posList[1])

	case "p":
		posStrList := strings.Split(cmd.args, "/")
		var posList []rune
		for _, posStr := range posStrList {
			posList = append(posList, []rune(posStr)[0])
		}
		manager.cmdPartner(posList[0], posList[1])
	}

}

func (manager *CmdManager) cmdSpin(size int) {
	//fmt.Println("Spin - ", size)
	pivot := len(manager.ipList) - size
	manager.ipList = append(manager.ipList[pivot:], manager.ipList[:pivot]...)
}

func (manager *CmdManager) cmdExchange(pos1, pos2 int) {
	//fmt.Println("Exchange - ", pos1, pos2)
	manager.ipList[pos2], manager.ipList[pos1] = manager.ipList[pos1], manager.ipList[pos2]
}

func (manager *CmdManager) cmdPartner(name1, name2 rune) {
	//fmt.Println("Partner - ", string(name1), string(name2))
	pos1 := manager.findPos(name1)
	pos2 := manager.findPos(name2)
	manager.cmdExchange(pos1, pos2)
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
	name string
	args string
}

func main() {
	filename := "ip.txt"
	numEntries := 16
	//filename := "trial.txt"
	//numEntries := 5

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
	runCount := 1000000000
	var p2List = make([]rune, len(ipList))
	copy(p2List, ipList)
	p2Manager := CmdManager{
		cmdList: cmdList,
		ipList:  p2List,
	}

	for i := 0; i < runCount; i++ {
		p2Manager.execute()
	}
	p2Manager.printList()
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
		cmd := Cmd{
			name: cmdStr,
			args: args,
		}
		cmdList = append(cmdList, cmd)
	}
	return cmdList
}
