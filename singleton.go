package singleton

import (
	"fmt"
	"sync"
)

// CarFactory represents the Singleton using the lazy initialization approach.
type CarFactory struct {
	carsCreated int
}

var instance *CarFactory
var once sync.Once

func getInstance() *CarFactory {
	once.Do(func() {
		instance = &CarFactory{}
	})
	return instance
}

func (cf *CarFactory) CreateCar(carType string) {
	cf.carsCreated++
	fmt.Printf("Created %s car #%d\n", carType, cf.carsCreated)
}


/* func main() {
	carFactory := getInstance()

	carFactory.CreateCar("SUV")
	carFactory.CreateCar("Sedan")

	anotherCarFactory := getInstance()

	anotherCarFactory.CreateCar("Hatchback")
} */
