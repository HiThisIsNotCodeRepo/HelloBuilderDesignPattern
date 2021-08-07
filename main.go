package main

import "fmt"

func main() {
	aNormalBuilder := getBuilder("normal")
	anIglooBuilder := getBuilder("igloo")

	aDirector := newDirector(aNormalBuilder)
	aNormalHouse := aDirector.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", aNormalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", aNormalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", aNormalHouse.floor)

	aDirector.setBuilder(anIglooBuilder)
	anIglooHouse := aDirector.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", anIglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", anIglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", anIglooHouse.floor)

}
