package main

import (
	"fmt"
)

type PortType int

// Component
type Component struct {
	port1 PortType
	port2 PortType
}

func CreateComponentWithPorts(portA, portB PortType) Component {
	port1 := portA
	port2 := portB

	if port1 > port2 {
		port1, port2 = port2, port1
	}
	return Component{
		port1: port1,
		port2: port2,
	}
}

// Component manager
type ComponentManager struct {
	componentList []Component
	componentMap  map[PortType][]Component
	graph         Graph
}

func CreateManagerWithComponentList(compList []Component) ComponentManager {
	componentMap := make(map[PortType][]Component)

	// Find the zero entries in the map
	zeroPortType := PortType(0)
	var zeroComponentList []Component
	for _, component := range compList {
		if component.port1 == zeroPortType {
			zeroComponentList = append(zeroComponentList, component)
		}
	}
	if len(zeroComponentList) == 0 {
		panic("Does not have any components with zero width port")
	}

	componentMap[zeroPortType] = zeroComponentList

	componentManager := ComponentManager{
		componentList: compList,
		componentMap:  componentMap,
	}
	componentManager.linkComponents(zeroPortType)
	return componentManager
}
func (self *ComponentManager) printComponents() {
	for _, comp := range self.componentList {
		fmt.Println(comp)
	}
}

func (self *ComponentManager) printComponentsMap() {
	for key, val := range self.componentMap {
		fmt.Println(key, " => ", val)
	}
}

func (self *ComponentManager) linkComponents(initialPort PortType) {
	// Would a recursive solution work here?
	rootNode := CreateNode(int(initialPort))
	self.graph = Graph{
		rootNode: rootNode,
	}

	for _, component := range self.componentList {

	}
}
