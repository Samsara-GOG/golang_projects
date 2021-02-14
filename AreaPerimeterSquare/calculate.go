package main

import (
	"fmt"
	"time"
)

func calculateAreaSquare(side int) int {
	return side * side
}

func calculateAreaAndPerimeterSquare(side int) (int, int) {
	area := side * side
	perimeter := side * 4
	return area, perimeter
}

func joinTwoStrings(prefix string, suffix string) string {
	return prefix + suffix
}

func printOutDate() {
	t := time.Now()
	fmt.Println(t)
}

func waitForIt(message string) {
	defer fmt.Println("Done!")
	fmt.Println("Waiting")
	fmt.Println("Waiting")
	fmt.Println("Waiting")
	fmt.Println(message)
}

func main() {
	printOutDate()
	waitForIt(joinTwoStrings("Hi", " there"))

	side := 5
	resultArea := calculateAreaSquare(side)
	fmt.Printf("L'aire d'un carré avec %v de côté est égal à %v", side, resultArea)

	var area, perimeter int
	area, perimeter = calculateAreaAndPerimeterSquare(side)

	if perimeter > 0 {
		fmt.Printf("\n Le périmère d'un carré avec %v de côté est égal à %v", side, perimeter)
	}

	if area > 0 {
		fmt.Printf("\n L'aire d'un carré avec %v de côté est égal à %v", side, area)
	}

}
