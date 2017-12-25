package main

type Node struct {
	value    int
	children []Node
}

func CreateNode(Value int) Node {
	return Node{
		value: Value,
	}
}

type Graph struct {
	rootNode Node
}
