package main

import "fmt"

// Third-party CarInfoProvider with a different interface
type ThirdPartyCarInfoProvider struct {
	ModelName string
	TopSpeed  int
}

// Interface expected by your project
type CarInfoProvider interface {
	GetModel() string
	GetMaxSpeed() int
}

// Adapter to make ThirdPartyCarInfoProvider compatible
type Adapter struct {
	ThirdPartyProvider *ThirdPartyCarInfoProvider
}

func (a *Adapter) GetModel() string {
	return a.ThirdPartyProvider.ModelName
}

func (a *Adapter) GetMaxSpeed() int {
	return a.ThirdPartyProvider.TopSpeed
}

// Function in your project that uses CarInfoProvider
func DisplayCarInfo(provider CarInfoProvider) {
	fmt.Printf("Car Model: %s\n", provider.GetModel())
	fmt.Printf("Max Speed: %d mph\n", provider.GetMaxSpeed())
}

func main() {
	// Simulating data from a third-party library
	thirdPartyProvider := &ThirdPartyCarInfoProvider{
		ModelName: "Example Car",
		TopSpeed:  200,
	}

	// Creating an adapter
	adapter := &Adapter{ThirdPartyProvider: thirdPartyProvider}

	// Using the adapter in your project
	DisplayCarInfo(adapter)
}
