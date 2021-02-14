package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("You have", fuel, "gallons of fuel left!")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	case "Earth":
		fuel = 250000
	/*Challenge: Add more destinations and allow for traveling between different planets.
	Venus to ...*/
	case "Venus-Mercury":
		fuel = 0
	case "Venus-Mars":
		fuel = 0
	case "Venus-Earth":
		fuel = 0
	// Mercury to ...
	case "Mercury-Venus":
		fuel = 0
	case "Mercury-Mars":
		fuel = 0
	case "Mercury-Earth":
		fuel = 0
	// Mars to ...
	case "Mars-Venus":
		fuel = 0
	case "Mars-Mercury":
		fuel = 0
	case "Mars-Earth":
		fuel = 0
	// Earth to ...
	case "Earth-Venus":
		fuel = 0
	case "Earth-Mercury":
		fuel = 0
	case "Earth-Mars":
		fuel = 0
	default:
		fuel = 0
	}
	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Println("Welcome to planet", planet)
}

//Challenge: Add a variable that keeps track of which planet your spaceship is on.

//Challenge: Create a function that returns your spaceship back to your home planet.

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
		fuel = fuelRemaining
		/*fmt.Println(fuelRemaining)*/
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	// Test FuelRemaining
	/*var fuelRemain int = 100
	  FuelGauge(fuelRemain)*/

	// Test calculateFuel to Planet
	/*var planet string = "Venus"
	  var calcFuel int
	  calcFuel = calculateFuel(planet)
	  fmt.Printf("\nIl faut %v gallons de fuel pour aller sur %v.", calcFuel, planet)*/

	// Test greetPlanet
	/*var planet string = "Venus"
	  greetPlanet(planet)*/

	// Test cantFly
	/*cantFly()*/

	//Test flyToPlanet
	/*var planet string = "Mercury"
	  fmt.Println(flyToPlanet(planet, 600000))*/

	// Create `planetChoice` and `fuel`

	// Fly Me To The Stars
	fuel := 1000000
	planetChoice := "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)

	// And then liftoff!

}
