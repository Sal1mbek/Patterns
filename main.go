package main

import (
	"fmt"
	"singleton" // Import the 'singleton' package
)

func main() {
	// Call 'getInstance' from the 'singleton' package to get a single instance of 'CarFactory'
	carFactory := singleton.GetCarFactory()

	// Use the 'CarFactory' instance to create cars
	carFactory.CreateCar("SUV")
	carFactory.CreateCar("Sedan")
	carFactory.CreateCar("Hatchback")
	carFactory.CreateCar("Crossover")

	// You can also access the same 'CarFactory' instance elsewhere in your 'main' or other packages
	// This demonstrates that the Singleton pattern ensures a single instance across your application.

	// Create more cars
	carFactory.CreateCar("Wagon")
	carFactory.CreateCar("Minivan")
	carFactory.CreateCar("Pick-UP")
}
