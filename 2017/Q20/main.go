package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	filename := "ip.txt"
	//filename := "trial.txt"
	particlesList := getIpListFromFilename(filename)
	particlesBuffer := ParticleBuffer{
		particlesList,
	}
	//fmt.Println("Particle closest to the origin - ", particlesBuffer.findClosestToOrigin())
	// Part 2
	numParticlesLeft := particlesBuffer.performCollisions()
	fmt.Println("Num particles left after collision - ", numParticlesLeft)

	fmt.Println()
	// 516 is too high
}

// FYI: This method ran properly the very first time! No compiler or run time errors!
// :cool_shades:
func getIpListFromFilename(filename string) []Particle {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")
	var particlesList []Particle

	for _, ipStr := range ipStrList {
		constituents := strings.Split(ipStr, ", ")
		var tripletList []Triplet
		for _, c := range constituents {
			keyVal := strings.Split(c, "=")
			val := keyVal[len(keyVal)-1]
			val = strings.Trim(val, "<>")
			pointsStrList := strings.Split(val, ",")
			xVal, _ := strconv.Atoi(pointsStrList[0])
			yVal, _ := strconv.Atoi(pointsStrList[1])
			zVal, _ := strconv.Atoi(pointsStrList[2])

			triplet := Triplet{
				xVal, yVal, zVal,
			}
			tripletList = append(tripletList, triplet)
		}

		particle := Particle{
			tripletList[0],
			tripletList[1],
			tripletList[2],
		}
		particlesList = append(particlesList, particle)
	}
	return particlesList
}
