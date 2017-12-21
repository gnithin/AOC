package main

import (
	"fmt"
	"math"
)

type Triplet struct {
	X int
	Y int
	Z int
}

// The Individual particle
type Particle struct {
	position     Triplet
	velocity     Triplet
	acceleration Triplet
}

func (p *Particle) printP() {
	fmt.Println("position    ", p.position)
	fmt.Println("velocity    ", p.velocity)
	fmt.Println("acceleration", p.acceleration)
	fmt.Println("***")
}

func (self *Particle) overallAccln() uint {
	accln := self.acceleration
	overallAccln := uint(math.Abs(float64(accln.X))) +
		uint(math.Abs(float64(accln.Y))) +
		uint(math.Abs(float64(accln.Z)))
	return overallAccln
}

// The particle Buffer
type ParticleBuffer struct {
	particlesList []Particle
}

func (self *ParticleBuffer) findClosestToOrigin() int {
	minIndex := -1
	minAccln := ^uint(0)

	for currIndex, p := range self.particlesList {
		overallAccln := p.overallAccln()
		if overallAccln < minAccln {
			minIndex = currIndex
			minAccln = overallAccln
		}
	}
	return minIndex
}

func (self *ParticleBuffer) printBuffer() {
	fmt.Println("Number of particles - ", len(self.particlesList))
	for _, p := range self.particlesList {
		p.printP()
	}
}

func (self *ParticleBuffer) performCollisions() int {
	numValidParticles := 0
	var particleQueue []Particle = make([]Particle, len(self.particlesList))
	copy(particleQueue, self.particlesList)

	for len(particleQueue) > 0 {
		poppedParticle := popSlice(0, &particleQueue)

		// Finding min tick -
		minTickVal := 1000000000000
		minTickIndex := -1
		for i, p := range particleQueue {
			intTick := self.findIntersectionTick(p, poppedParticle)
			if intTick != -1 && intTick < minTickVal {
				minTickVal = intTick
				minTickIndex = i
			}
		}

		if minTickIndex == -1 {
			numValidParticles += 1
		} else {
			// Remove the item
			popSlice(minTickIndex, &particleQueue)

			// Remove all items with similar ticks
			var newSlice []Particle
			for _, p := range particleQueue {
				if self.findIntersectionTick(poppedParticle, p) != minTickVal {
					newSlice = append(newSlice, p)
				}
			}
			particleQueue = newSlice
		}
	}
	return numValidParticles
}

func popSlice(pos int, slice *[]Particle) Particle {
	particle := (*slice)[pos]
	(*slice) = append((*slice)[:pos], (*slice)[pos+1:]...)
	return particle
}

func (self *ParticleBuffer) findIntersectionTick(p1, p2 Particle) int {
	xStatus, commonXTicks := self.findCommonTick(
		p1.acceleration.X,
		p1.velocity.X,
		p1.position.X,
		p2.acceleration.X,
		p2.velocity.X,
		p2.position.X,
	)

	yStatus, commonYTicks := self.findCommonTick(
		p1.acceleration.Y,
		p1.velocity.Y,
		p1.position.Y,
		p2.acceleration.Y,
		p2.velocity.Y,
		p2.position.Y,
	)

	zStatus, commonZTicks := self.findCommonTick(
		p1.acceleration.Z,
		p1.velocity.Z,
		p1.position.Z,
		p2.acceleration.Z,
		p2.velocity.Z,
		p2.position.Z,
	)

	if xStatus == -1 || yStatus == -1 || zStatus == -1 {
		return -1
	}

	if xStatus == -2 && yStatus == -2 && zStatus == -2 {
		return -1
	}

	commonTick := self.findCommonElements(commonXTicks, commonYTicks, commonZTicks)
	return commonTick
}

func (p *ParticleBuffer) findCommonElements(
	xTicks []int,
	yTicks []int,
	zTicks []int,
) int {
	elementCountMap := make(map[int]int)
	updateMap := func(l []int) {
		for _, i := range l {
			val, ok := elementCountMap[i]
			newVal := 1
			if ok {
				newVal = val + 1
			}
			elementCountMap[i] = newVal
		}
	}

	updateMap(xTicks)
	updateMap(yTicks)
	updateMap(zTicks)

	// Find the most common among the maps
	for key, val := range elementCountMap {
		if val > 2 {
			return key
		}
	}
	return -1
}

/*
Derived this beautiful formula all on my own! Fuck yeah \m/

(t^2)*(A2-A1) + t*(2*(V2 - V1) + A1 - A2) - 2(P1 - P2 + V2 - V1) = 0

The positive,integral roots above formula finds the tick for the intersection of the path!
Note that there could be multiple. findCommonTick returns a list.

The formula is basically derived from -
pos(t) = pos1 + (t-1)*V1 + ((t*(t-1))/2)*A

pos(t) for 2 different particles will the same if they intersect.
So posx1(t) == posx2(t) && posy1(t) == posy2(t) && posz1(t) == posz2(t) individually
for each coordinate plane
*/
func (self *ParticleBuffer) findCommonTick(
	A1, V1, P1,
	A2, V2, P2 int,
) (int, []int) {
	var rootsList []int

	a := A2 - A1
	b := 2*(V2-V1) + A1 - A2
	c := -2 * (P1 - P2 + V2 - V1)

	if a == 0 {
		// It's not a quadriatic equation then
		velDiff := V2 - V1
		if velDiff == 0 {
			if P1 == P2 {
				return -2, rootsList
			} else {
				//fmt.Println("There is no velocity difference")
				return -1, rootsList
			}
		}

		r1 := float64(P1-P2+V2-V1) / float64(velDiff)
		if math.Floor(r1)-r1 != 0 {
			//fmt.Println("Never intersect even without quadriatic equation")
			return -1, rootsList
		}

		r1Val := int(r1)
		rootsList = append(rootsList, r1Val)
		return 1, rootsList
	}

	// Find the roots r1 and r2
	determinant := (b * b) - (4 * a * c)
	// Determinant is non-neg
	if determinant < 0 {
		//fmt.Println("determinant is < 0")
		return -1, rootsList
	}

	// Determinant is a perfect square
	dSqrt := math.Sqrt(float64(determinant))
	if dSqrt-math.Floor(dSqrt) != 0 {
		//fmt.Println(dSqrt, " determinant is not a perfect square")
		return -1, rootsList
	}

	r1Num := float64(-1*b) + (dSqrt)
	r2Num := float64(-1*b) - (dSqrt)
	deno := float64(2 * a)

	r1 := r1Num / deno
	r2 := r2Num / deno

	if (r1-math.Floor(r1)) == 0 && r1 >= 0 {
		rootsList = append(rootsList, int(r1))
	}
	if (r2-math.Floor(r2)) == 0 && r2 >= 0 {
		rootsList = append(rootsList, int(r2))
	}
	return 1, rootsList
}
