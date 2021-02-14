/* Bank Heist

We often include conditionals for when we need to account for different
situations and outcomes. When it comes to coming up with a plan and
executing it, there’s nothing more uncertain than a bank heist!
(Watch any robbery themed movie if you need convincing).
In this project, we’re going to use conditionals to simulate
this situation along with hi-jinks and hiccups that may pop up.
Who knows? Maybe, just maybe, if all goes well, we’ll have a few more dollars after.
*/

/* If you want to challenge yourself:

   Change the maximum generated values and then adjust conditionals in response.
   Add more conditionals that account for more scenarios, e.g.
       What time of day is it?
       Does that affect the heist?
       Do the number of guards matter?
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100) // 0-99

	// First Conditional (Sneak past guards)
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	// Second Conditional (Open the vault)
	openedVault := rand.Intn(100)
	fmt.Println(isHeistOn)

	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn == true {
		isHeistOn = false
		fmt.Println("Vault can't be opened!")
	}
	// Third Conditional (Leaving)
	leftSafely := rand.Intn(5)
	if isHeistOn == true {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Print("Turns out vault doors don't open from the inside...")
		case 2:
			// Code for another failure
		case 3:
			// Code for another failure
		case 4:
			// Code for another failure
		default:
			fmt.Println("Start the getaway car!")
		}
	}
	// Wrapping up
	if isHeistOn == true {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("$", amtStolen, "not bad!")
	}
}
