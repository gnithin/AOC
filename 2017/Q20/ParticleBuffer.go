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
		fmt.Println("position    ", p.position)
		fmt.Println("velocity    ", p.velocity)
		fmt.Println("acceleration", p.acceleration)
		fmt.Println("***")
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
				if self.findIntersectionTick(p, poppedParticle) != minTickVal {
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
	A1 := p1.acceleration.X
	V1 := p1.velocity.X
	P1 := p1.position.X

	A2 := p2.acceleration.X
	V2 := p2.velocity.X
	P2 := p2.position.X

	commonXTick := self.findCommonTick(
		A1, V1, P1,
		A2, V2, P2,
	)

	commonYTick := self.findCommonTick(
		p1.acceleration.Y,
		p1.velocity.Y,
		p1.position.Y,
		p2.acceleration.Y,
		p2.velocity.Y,
		p2.position.Y,
	)

	commonZTick := self.findCommonTick(
		p1.acceleration.Z,
		p1.velocity.Z,
		p1.position.Z,
		p2.acceleration.Z,
		p2.velocity.Z,
		p2.position.Z,
	)

	if commonXTick == -1 || commonYTick == -1 || commonZTick == -1 {
		return -1
	}

	if commonXTick == -2 && commonYTick == -2 && commonZTick == -2 {
		return 1
	}

	var compList []int
	if commonXTick != -2 {
		compList = append(compList, commonXTick)
	}
	if commonYTick != -2 {
		compList = append(compList, commonYTick)
	}
	if commonZTick != -2 {
		compList = append(compList, commonZTick)
	}

	fVal := compList[0]
	matched := true
	for _, c := range compList[1:] {
		if c != fVal {
			matched = false
		}
	}

	if matched {
		return fVal
	}
	return -1
}

func (self *ParticleBuffer) findCommonTick(
	A1, V1, P1,
	A2, V2, P2 int,
) int {
	/*
		fmt.Println("A V P")
		fmt.Println(A1, V1, P1)
		fmt.Println(A2, V2, P2)
		fmt.Println("****\n")
	*/

	a := A2 - A1
	b := 2*(V2-V1) + A1 - A2
	c := -2 * (P1 - P2 + V2 - V1)

	if a == 0 {
		// It's not a quadriatic equation then
		velDiff := V2 - V1
		if velDiff == 0 {
			if P1 == P2 {
				return -2
			} else {
				fmt.Println("There is no velocity difference")
				return -1
			}
		}

		r1 := float64(P1-P2+V2-V1) / float64(velDiff)
		if math.Floor(r1)-r1 != 0 {
			fmt.Println("Never intersect even without quadriatic equation")
			return -1
		}

		r1Val := int(r1)
		return r1Val
	}

	// Find the roots r1 and r2
	determinant := (b * b) - (4 * a * c)
	// Determinant is non-neg
	if determinant < 0 {
		fmt.Println("determinant is < 0")
		return -1
	}

	// Determinant is a perfect square
	dSqrt := math.Sqrt(float64(determinant))
	if dSqrt-math.Floor(dSqrt) != 0 {
		fmt.Println(dSqrt, " determinant is not a perfect square")
		return -1
	}

	r1Num := float64(-1*b) + (dSqrt)
	r2Num := float64(-1*b) - (dSqrt)
	deno := float64(2 * a)

	r1 := r1Num / deno
	r2 := r2Num / deno

	validRoot := -1
	if r1-math.Floor(r1) == 0 && r1 > 0 {
		validRoot = int(r1)
	} else if r2-math.Floor(r2) == 0 && r2 > 0 {
		validRoot = int(r2)
	}
	return validRoot
}
