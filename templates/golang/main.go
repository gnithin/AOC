package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//filename := "ip.txt"
	filename := "trial.txt"
	ipList := getIpListFromFilename(filename)
	fmt.Printf("Input - [%s]\n", strings.Join(ipList, ", "))
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
