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
