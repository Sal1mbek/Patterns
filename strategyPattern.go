package main

import (
	"fmt"
)

// CarBuyingStrategy defines the strategy interface for buying a car.
type CarBuyingStrategy interface {
	BuyCar() string
}

// SUVStrategy-buying an SUV.
type SUVStrategy struct{}

func (s *SUVStrategy) BuyCar() string {
	return "Buying an SUV"
}

// SedanStrategy-buying a sedan.
type SedanStrategy struct{}

func (s *SedanStrategy) BuyCar() string {
	return "Buying a sedan"
}

// HatchbackStrategy-buying a hatchback.
type HatchbackStrategy struct{}

func (h *HatchbackStrategy) BuyCar() string {
	return "Buying a hatchback"
}

// CarBuyer is the context that uses the car buying strategy.
type CarBuyer struct {
	strategy CarBuyingStrategy
}

func (c *CarBuyer) SetStrategy(strategy CarBuyingStrategy) {
	c.strategy = strategy
}

func (c *CarBuyer) BuyCar() string {
	return c.strategy.BuyCar()
}

func main() {
	buyer := &CarBuyer{}
	// Buying a sedan
	
	buyer.SetStrategy(&SedanStrategy{})
	fmt.Println(buyer.BuyCar())
}
