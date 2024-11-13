package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal house door type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal house window type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal house num floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("Igloo house door type %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo house window type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo house num floor: %d\n", iglooHouse.floor)
}