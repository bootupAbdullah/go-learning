package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
fmt.Printf("Current fuel levels are: %v \n", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  
  // Declare ""fuel" variable
  var fuel int

  // Create switch statement to select "fuel" based on planet
  switch planet {
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    default:
      fuel = 0
    }

    // return fuel
    return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) string {
   
   // Delcare var to hold greeting
    var greeting string

    // Create switch statement to decide message based on planet
    switch planet {
      
      case "Venus":
        greeting = "What up, yall? It's ya boy Claptrap, coming to you live form Venus, the hottest planet in our solar system. Get ready to have yo ass burned alive at a moment's notice. Pack ya cooling gear because we got volcanoes and this ain't dragon ball z MF so you best beleive you ass gon be crushed alive if you step outside this here gravity zone, ya heard?"
      
      case "Mercury":
        greeting = "Welcome to Mercury, son."

      case "Mars":
        greeting = "You haven't made it that far. Only Mars, welcome."
      
      default: 
        greeting = "We haven't taken off yet."
    }

    return greeting
  }

// Create the function cantFly() here
func cantFly() {
  // var msg string
  // msg = "We do not have the availabe fuel to fly there."
  fmt.Println("We do not have the availabe fuel to fly there.")
} 

// Create the function flyToPlanet() here
 func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int

  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)

  if (fuelRemaining >= fuelCost) {
    fmt.Println(greetPlanet(planet))
    fuelCost -= fuelRemaining
  } else if (fuelCost > fuelRemaining) {
    cantFly()
  }

  return fuelRemaining

 }

func main() {

  var fuel int
  fuel = 1000000

  var planetChoice string
  planetChoice = "Venus"

  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)

  // And then liftoff!
  
}