package main

import "fmt"

func performAddition() {
	x := 5
	y := 7
	fmt.Println("The sum of", x, "and", y, "is", x+y)
}

func main() {
	performAddition()
	fmt.Println("What if", x, "was different?")
}
