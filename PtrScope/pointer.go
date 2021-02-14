package main

import "fmt"

/* If you want to practice more with pointers and addresses:

    Create a function that changes a boolean value from a different scope.
    Create a function that changes a float value from a different scope.
	Find the zero value of a pointer that is initialized but hasn’t been assigned
	a value.

*/

func fuckBoolean(fuck *bool) {
	*fuck = false
}

func fuckfloat(flot *float32) {
	*flot = 10.0
}

func brainwash(saying *string) {
	*saying = "Beep Boop."
}

func main() {
	// Change a boolean value from a different scope
	/*verite := true
	fmt.Printf("Before: It's %t", verite)
	fuckBoolean(&verite)
	fmt.Printf("\nAfter: It's %t", verite)*/

	// Create a function that changes a float value from a different scope.
	/*var num float32 = 5.6
	fmt.Printf("\nBefore: %g", num)
	fuckfloat(&num)
	fmt.Printf("\nAfter: %g", num)*/

	// Find the zero value of a pointer that is initialized but hasn’t been assigned a value.

	var varpointerForInt *int
	fmt.Println(varpointerForInt)

	/* Change of Greeting from a different scope
	*greeting := "Hello there!"
	brainwash(&greeting)
	fmt.Println("\ngreeting is now:", greeting)*
	*/

}
