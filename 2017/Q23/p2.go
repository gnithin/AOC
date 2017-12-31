package main

import (
	"fmt"
)

var a = 1
var b = 0
var c = 0
var d = 0
var e = 0
var f = 0
var g = 0
var h = 0

func prog() {
	b = 65
	c = b
	if a != 0 {
		b = (b * 100) + 100000
		c = b + 17000
	}

	//for {
	f = 1
	d = 2
	e = 2

	for {
		// section 1
		for {
			g = (d * e) - b
			if g == 0 {
				f = 0
			}
			e++
			g = e - b
			fmt.Println(e - b)
			fmt.Println(b)
			if g == 0 {
				break
			}
		}
		// section 1 ends
		d++
		g = d - b
		if g == 0 {
			break
		}
		printVars()
	}

	/*
		// 29
		if f == 0 {
			h = h - 1
		}
		g = b - c
		if g == 0 {
			return
		}
		b = b + 17

		fmt.Println("Internal iterations done!")
		break
	*/
	//}
}

func printVars() {
	fmt.Println("*********")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
	fmt.Println("*********")
}

func main() {
	prog()
	printVars()
	fmt.Println("h - ", h)
}
