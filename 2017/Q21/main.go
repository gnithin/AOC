package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filename := "ip.txt"
	//filename := "trial.txt"
	enhancementRules := getIpListFromFilename(filename)
	initialImg := [][]string{
		{".", "#", "."},
		{".", ".", "#"},
		{"#", "#", "#"},
	}
	img := CreateImage(initialImg, enhancementRules)

	// Part 1
	numTimesIterate := 5
	img.iterateNTimes(numTimesIterate)
	fmt.Println("********")
	numOnPixels := img.getNumOnPixels()
	fmt.Println("P1: Number of on pixels", numOnPixels)

	// Part 2
	numTimesIterate = 13
	img.iterateNTimes(numTimesIterate)
	fmt.Println("********")
	numOnPixels = img.getNumOnPixels()
	fmt.Println("P2: Number of on pixels", numOnPixels)

}

func getIpListFromFilename(filename string) map[string]string {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")

	rules := make(map[string]string)
	for _, ipStr := range ipStrList {
		components := strings.Split(ipStr, "=>")
		key := strings.TrimSpace(components[0])
		val := strings.TrimSpace(components[1])
		rules[key] = val
	}
	return rules
}
