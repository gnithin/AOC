package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	//filename := "ip.txt"
	filename := "trial.txt"
	compList := getIpListFromFilename(filename)
	fmt.Println(compList)
	fmt.Println("*****")

	maxScore := getMaxScoreFromCompList(compList)
	fmt.Println("*****")
	fmt.Println("Max score - ", maxScore)
	fmt.Println("*****")
}

func getIpListFromFilename(filename string) []Component {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")
	var compList []Component
	for _, ipStr := range ipStrList {
		var portComponents []int
		for _, c := range strings.Split(ipStr, "/") {
			intVal, _ := strconv.Atoi(c)
			portComponents = append(portComponents, intVal)
		}

		compList = append(compList,
			CreateComponentWithPorts(
				portComponents[0],
				portComponents[1],
			),
		)
	}
	return compList
}
