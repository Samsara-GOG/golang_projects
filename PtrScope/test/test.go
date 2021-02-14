package main

import "fmt"

func addHundred(numPtr *int) {
	*numPtr += 100
}

func main() {
	x := 1
	addHundred(&x)
	fmt.Println(x) // Prints 101
}
