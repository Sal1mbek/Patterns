package main

import "fmt"

// Engine represents the car's engine.
type Engine struct{}

func (e *Engine) start() {
	fmt.Println("Car engine is started")
}

func (e *Engine) stop() {
	fmt.Println("Car engine is stopped")
}

// Lights represents the car's lights.
type Lights struct{}

func (l *Lights) turnOn() {
	fmt.Println("Car lights are turned on")
}

func (l *Lights) turnOff() {
	fmt.Println("Car lights are turned off")
}

// AirConditioner represents the car's air conditioner.
type AirConditioner struct{}

func (ac *AirConditioner) turnOn() {
	fmt.Println("Car air conditioner is turned on")
}

func (ac *AirConditioner) turnOff() {
	fmt.Println("Car air conditioner is turned off")
}

// CarSystemFacade provides a simplified interface for the car system.
type CarSystemFacade struct {
	engine         *Engine
	lights         *Lights
	airConditioner *AirConditioner
}

// NewCarSystemFacade creates a new CarSystemFacade.
func NewCarSystemFacade() *CarSystemFacade {
	return &CarSystemFacade{
		engine:         &Engine{},
		lights:         &Lights{},
		airConditioner: &AirConditioner{},
	}
}

// StartCar is a simplified method for starting the car.
func (facade *CarSystemFacade) StartCar() {
	fmt.Println("Starting the car...")
	facade.engine.start()
	facade.lights.turnOn()
	facade.airConditioner.turnOn()
	fmt.Println("Car is ready to go!")
}

// StopCar is a simplified method for stopping the car.
func (facade *CarSystemFacade) StopCar() {
	fmt.Println("Stopping the car...")
	facade.engine.stop()
	facade.lights.turnOff()
	facade.airConditioner.turnOff()
	fmt.Println("Car is safely parked.")
}

func main() {
	// Create CarSystemFacade
	carSystem := NewCarSystemFacade()

	// Start the car using the facade
	carSystem.StartCar()

	// Stop the car using the facade
	carSystem.StopCar()
}
