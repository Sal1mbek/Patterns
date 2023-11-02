package main

import "fmt"

// Car is the interface for all types of cars.
type FactoryCar interface {
	Drive()
}

// EngineCar is a concrete type that implements the Car interface.
type EngineCar struct{}

func (c *EngineCar) Drive() {
	fmt.Println("You are going to drive an engine car")
}

// ElectricCar is a concrete type that implements the Car interface.
type ElectricCar struct{}

func (c *ElectricCar) Drive() {
	fmt.Println("You are going to drive an electric car")
}

// HybridCar is a concrete type that implements the Car interface.
type HybridCar struct{}

func (c *HybridCar) Drive() {
	fmt.Println("You are going to drive a hybrid car")
}

// CarFactory is the factory interface for creating cars.
type CarFactory interface {
	CreateCar() FactoryCar
}

// EngineCarFactory is a concrete factory that creates engine cars.
type EngineCarFactory struct{}

func (cf EngineCarFactory) CreateCar() FactoryCar {
	return &EngineCar{}
}

// ElectricCarFactory is a concrete factory that creates electric cars.
type ElectricCarFactory struct{}

func (cf ElectricCarFactory) CreateCar() FactoryCar {
	return &ElectricCar{}
}

// HybridCarFactory is a concrete factory that creates hybrid cars.
type HybridCarFactory struct{}

func (cf HybridCarFactory) CreateCar() FactoryCar {
	return &HybridCar{}
}

/* func main() {
	// Create an engine car using the EngineCarFactory
	Dodge := EngineCarFactory{}
	RAM_trx := Dodge.CreateCar()
	RAM_trx.Drive()

	// Create an electric car using the ElectricCarFactory
	Zeekr := ElectricCarFactory{}
	z_001 := Zeekr.CreateCar()
	z_001.Drive()

	// Create a hybrid car using the HybridCarFactory
	Fisker := HybridCarFactory{}
	Carma := Fisker.CreateCar()
	Carma.Drive()
} */
