package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//Create a conditional that prints the price of an item when a condition is matched.
	rand.Seed(time.Now().UnixNano())
	price := rand.Intn(140)
	fmt.Println(price)
	if price < 10 {
		fmt.Println("Le prix est inferieur à 10")
	} else if price < 20 {
		fmt.Println("Le prix est inferieur à 20")
	} else if price < 40 {
		fmt.Println("Le prix est inferieur à 40")
	} else if price < 60 {
		fmt.Println("Le prix est inferieur à 60")
	} else if price < 80 {
		fmt.Println("Le prix est inferieur à 80")
	} else if price < 100 {
		fmt.Println("Le prix est inferieur à 100")
	} else if price < 120 {
		fmt.Println("Le prix est inferieur à 120")
	} else {
		fmt.Println("Le prix est supérieur à 120")
	}
}
