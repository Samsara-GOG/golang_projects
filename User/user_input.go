package main

import "fmt"

func main() {
	fmt.Println("What's your name?")
	var response1 string
	var response2 string
	fmt.Scan(&response1)
	fmt.Scan(&response2)
	fmt.Printf("I'm %v %v", response1, response2)
}
