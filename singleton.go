package main

import (
	"fmt"
	"sync"
)

// CarFactory represents the Singleton using the lazy initialization approach.
type CarsFactory struct {
	carsCreated int
}

var instance *CarsFactory
var once sync.Once

func getInstance() *CarsFactory {
	once.Do(func() {
		instance = &CarsFactory{}
	})
	return instance
}

func (cf *CarsFactory) CreateCar(carType string) {
	cf.carsCreated++
	fmt.Printf("Created %s car #%d\n", carType, cf.carsCreated)
}


/* func main() {
	carFactory := getInstance()

	carFactory.CreateCar("SUV")
	carFactory.CreateCar("Sedan")

	anotherCarFactory := getInstance()

	anotherCarFactory.CreateCar("Hatchback")
}
 */