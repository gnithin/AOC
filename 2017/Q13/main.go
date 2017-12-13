package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Part 1

type Layer struct {
	depth          int
	scannerRange   int
	scannerPos     int
	containsPacket bool
	scannerDirnInc bool
}

type Packet struct {
	currLayer int
}

func (l *Layer) String() string {
	str := fmt.Sprintf("%d : ", l.depth)
	if l.scannerRange == 0 {
		s := "x"
		if l.containsPacket == true {
			s = "()"
		}
		str += fmt.Sprintf(s)
	}

	for i := 0; i < l.scannerRange; i++ {
		if i == 0 && l.containsPacket == false || i != 0 {
			item := "[]"
			if i == l.scannerPos {
				item = "[S]"
			}
			str += fmt.Sprintf("%s ", item)
		} else {
			item := "[()]"
			if i == l.scannerPos {
				item = "[(S)]"
			}
			str += fmt.Sprintf("%s ", item)
		}
	}
	return str
}

func (l *Layer) incrementScanner() {
	if l.scannerRange == 0 {
		return
	}

	if l.scannerDirnInc {
		if l.scannerPos == (l.scannerRange - 1) {
			l.scannerPos = l.scannerPos - 1
			l.scannerDirnInc = false
		} else {
			l.scannerPos += 1
		}
	} else {
		if l.scannerPos == 0 {
			l.scannerPos = 1
			l.scannerDirnInc = true
		} else {
			l.scannerPos -= 1
		}
	}
}

func findSeverity(depthRangeMap map[int]int, delay int) (bool, int) {
	var firewall []*Layer
	var severeDepths []int

	// Create a packet with initial position as -1
	packet := Packet{-1}

	// Create a firewall
	maxDepth := -1

	for depth, _ := range depthRangeMap {
		if maxDepth < depth {
			maxDepth = depth
		}
	}

	for i := 0; i <= maxDepth; i++ {
		rangeVal, _ := depthRangeMap[i]
		currLayer := Layer{
			depth:        i,
			scannerRange: rangeVal,
		}
		if packet.currLayer == i {
			currLayer.containsPacket = true
		} else {
			currLayer.containsPacket = false
		}
		currLayer.scannerDirnInc = true
		firewall = append(firewall, &currLayer)
	}

	// Go psec by psec with both modes
	numPicoSecs := maxDepth + 1 + delay
	for i := 0; i < numPicoSecs; i++ {
		/*
			banner := "*******"
			fmt.Printf("%s Picosecond - %d %s\n", banner, i, banner)
		*/

		// Update the packet before doing anything else
		if i >= delay {
			packet.currLayer += 1
		}
		for j := 0; j < len(firewall); j++ {
			layer := firewall[j]
			if packet.currLayer == j {
				layer.containsPacket = true
			} else {
				layer.containsPacket = false
			}
		}

		// Printing the firewall
		/*
			fmt.Println("Pre-picosecond")
			for j := 0; j < len(firewall); j++ {
				layer := firewall[j]
				fmt.Println(layer)
			}
		*/

		// Checking severity
		currPacketLayerNum := packet.currLayer
		if currPacketLayerNum != -1 {
			currPacketLayer := firewall[currPacketLayerNum]
			if currPacketLayer.scannerRange > 0 && currPacketLayer.scannerPos == 0 {
				severeDepths = append(severeDepths, packet.currLayer)
			}
		}

		// Update the firewall
		for j := 0; j < len(firewall); j++ {
			layer := firewall[j]
			layer.incrementScanner()
		}

		/*
			fmt.Println("Post-picosecond")
			// Printing the firewall
			for j := 0; j < len(firewall); j++ {
				layer := firewall[j]
				fmt.Println(layer)
			}
		*/
	}

	//fmt.Println("Severe - ", severeDepths)
	severity := 0
	for _, s := range severeDepths {
		r, _ := depthRangeMap[s]
		severity += s * r
	}

	isCaught := false
	if len(severeDepths) != 0 {
		isCaught = true
	}
	return isCaught, severity
}

// Part - 2
// Had a brute force solution that took a loong time, and could never complete
// This is an elegant solution, intellectually awesome (my opinion ofcourse)
// that presented itself after a really good dinner :)
func findMinDelay(depthRangeMap map[int]int) int {
	layerOrder := (*findLayerOrder(depthRangeMap))
	for delay := 0; ; delay += 1 {
		pathFound := true
		for _, layer := range layerOrder {
			currDepth, _ := depthRangeMap[layer]
			currPos := (delay + layer) % (2 * (currDepth - 1))
			if currPos == 0 {
				pathFound = false
				break
			}
		}
		if pathFound {
			return delay
		}
	}
	return -1
}

func findLayerOrder(depthRangeMap map[int]int) *[]int {
	maxRange := -1

	rangeDepthMap := make(map[int]*[]int)
	for d, r := range depthRangeMap {
		if r > maxRange {
			maxRange = r
		}

		dMap, ok := rangeDepthMap[r]
		if !ok {
			dMap := []int{d}
			rangeDepthMap[r] = &dMap
		} else {
			*dMap = append(*dMap, d)
		}
	}

	var layerOrder []int
	for i := 0; i <= maxRange; i++ {
		dList, ok := rangeDepthMap[i]
		if !ok {
			continue
		}
		layerOrder = append(layerOrder, (*dList)...)
	}

	return &layerOrder
}

// Main
func main() {
	//filename := "trial.txt"
	filename := "ip.txt"
	depthRangeMap := getIpListFromFilename(filename)

	_, severity := findSeverity(depthRangeMap, 0)
	fmt.Println("Severity -", severity)

	minDelay := findMinDelay(depthRangeMap)
	fmt.Println("Min delay - ", minDelay)

	//testMinDelay(depthRangeMap)
}

func testMinDelay(depthRangeMap map[int]int) {
	var timeDiffList []int64
	counter := 10

	for i := 0; i < counter; i++ {
		startTime := getCurrentTimeMillis()
		findMinDelay(depthRangeMap)
		endTime := getCurrentTimeMillis()
		timeDiff := endTime - startTime
		timeDiffList = append(timeDiffList, timeDiff)
	}
	fmt.Println("time - ", timeDiffList)
	sum := int64(0)
	for _, timeDiff := range timeDiffList {
		sum += timeDiff
	}
	fmt.Println("Avg time - ", sum/int64(len(timeDiffList)))

}

func getCurrentTimeMillis() int64 {
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	return millis
}

func getIpListFromFilename(filename string) map[int]int {
	ipListBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileContents := string(ipListBytes)
	fileContents = strings.TrimSpace(fileContents)
	ipStrList := strings.Split(fileContents, "\n")

	depthRangeMap := make(map[int]int)

	for _, ip := range ipStrList {
		depthRangeComp := strings.Split(ip, ":")
		depthVal, _ := strconv.Atoi(strings.TrimSpace(depthRangeComp[0]))
		rangeVal, _ := strconv.Atoi(strings.TrimSpace(depthRangeComp[1]))
		depthRangeMap[depthVal] = rangeVal
	}
	return depthRangeMap
}
