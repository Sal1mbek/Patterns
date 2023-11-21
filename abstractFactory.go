package main

import "fmt"

// Car parts interfaces
type Engine interface {
	Start()
}

type Interior interface {
	Design()
}

type Wheels interface {
	Roll()
}

// Car parts factories
type CarPartsFactory interface {
	CreateEngine() Engine
	CreateInterior() Interior
	CreateWheels() Wheels
}

// SUV parts
type SUVEngine struct{}

func (e *SUVEngine) Start() {
	fmt.Println("SUV Engine started")
}

type SUVInterior struct{}

func (i *SUVInterior) Design() {
	fmt.Println("Luxurious SUV Interior designed")
}

type SUVWheels struct{}

func (w *SUVWheels) Roll() {
	fmt.Println("SUV Wheels rolling")
}

type SUVCarPartsFactory struct{}

func (f *SUVCarPartsFactory) CreateEngine() Engine {
	return &SUVEngine{}
}

func (f *SUVCarPartsFactory) CreateInterior() Interior {
	return &SUVInterior{}
}

func (f *SUVCarPartsFactory) CreateWheels() Wheels {
	return &SUVWheels{}
}

// Sedan parts
type SedanEngine struct{}

func (e *SedanEngine) Start() {
	fmt.Println("Sedan Engine started")
}

type SedanInterior struct{}

func (i *SedanInterior) Design() {
	fmt.Println("Comfortable Sedan Interior designed")
}

type SedanWheels struct{}

func (w *SedanWheels) Roll() {
	fmt.Println("Sedan Wheels rolling")
}

type SedanCarPartsFactory struct{}

func (f *SedanCarPartsFactory) CreateEngine() Engine {
	return &SedanEngine{}
}

func (f *SedanCarPartsFactory) CreateInterior() Interior {
	return &SedanInterior{}
}

func (f *SedanCarPartsFactory) CreateWheels() Wheels {
	return &SedanWheels{}
}

// Car creation function using Abstract Factory
func createCar(factory CarPartsFactory) {
	engine := factory.CreateEngine()
	interior := factory.CreateInterior()
	wheels := factory.CreateWheels()

	fmt.Println("Assembling the car:")
	engine.Start()
	interior.Design()
	wheels.Roll()
}

// Main function
func main() {
	var carType string
	fmt.Println("Which car type you are going to buy? (SUV/Sedan)")
	fmt.Scanf("%s\n", &carType)

	var carFactory CarPartsFactory

	switch carType {
	case "SUV":
		carFactory = &SUVCarPartsFactory{}
	case "Sedan":
		carFactory = &SedanCarPartsFactory{}
	default:
		fmt.Println("Undefined car type")
		return
	}

	createCar(carFactory)
}
