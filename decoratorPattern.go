package main

// Car represents the base component interface.
type Car interface {
	GetDescription() string
	GetCost() float64
}

// BaseCar is a concrete component representing a basic car type.
type BaseCar struct {
	description string
	cost        float64
}

func (c *BaseCar) GetDescription() string {
	return c.description
}

func (c *BaseCar) GetCost() float64 {
	return c.cost
}

// SUV is a concrete component representing an SUV.
type SUV struct {
	BaseCar
}

func NewSUV() *SUV {
	return &SUV{
		BaseCar: BaseCar{
			description: "SUV",
			cost:        25000.0,
		},
	}
}

// Сrossover is a concrete component representing an crossover.
type Crossover struct {
	BaseCar
}

func NewCrossover() *Crossover {
	return &Crossover{
		BaseCar: BaseCar{
			description: "Crossover",
			cost:        25000.0,
		},
	}
}

// Wagon is a concrete component representing a wagon.
type Wagon struct {
	BaseCar
}

func NewWagon() *Wagon {
	return &Wagon{
		BaseCar: BaseCar{
			description: "Wagon",
			cost:        25000.0,
		},
	}
}

// Minivan is a concrete component representing a minivan.
type Minivan struct {
	BaseCar
}

func NewMinivan() *Minivan {
	return &Minivan{
		BaseCar: BaseCar{
			description: "Minivan",
			cost:        25000.0,
		},
	}
}

// Sedan is a concrete component representing a sedan.
type Sedan struct {
	BaseCar
}

func NewSedan() *Sedan {
	return &Sedan{
		BaseCar: BaseCar{
			description: "Sedan",
			cost:        20000.0,
		},
	}
}

// Hatchback is a concrete component representing a hatchback.
type Hatchback struct {
	BaseCar
}

func NewHatchback() *Hatchback {
	return &Hatchback{
		BaseCar: BaseCar{
			description: "Hatchback",
			cost:        18000.0,
		},
	}
}

// PickUp is a concrete component representing a PickUp.
type PickUp struct {
	BaseCar
}

func NewPickUp() *PickUp {
	return &PickUp{
		BaseCar: BaseCar{
			description: "Pick-Up",
			cost:        18000.0,
		},
	}
}

// CarDecorator is the base decorator.
type CarDecorator struct {
	car Car
}

func (d *CarDecorator) GetDescription() string {
	return d.car.GetDescription()
}

func (d *CarDecorator) GetCost() float64 {
	return d.car.GetCost()
}

// SunroofDecorator is a concrete decorator for adding a sunroof to a car.
type SunroofDecorator struct {
	CarDecorator
}

func NewSunroofDecorator(car Car) *SunroofDecorator {
	return &SunroofDecorator{CarDecorator{car}}
}

func (s *SunroofDecorator) GetDescription() string {
	return s.car.GetDescription() + " with Sunroof"
}

func (s *SunroofDecorator) GetCost() float64 {
	return s.car.GetCost() + 1000.0
}

// AdvancedAudioDecorator is a concrete decorator for adding an advanced audio system to a car.
type AdvancedAudioDecorator struct {
	CarDecorator
}

func NewAdvancedAudioDecorator(car Car) *AdvancedAudioDecorator {
	return &AdvancedAudioDecorator{CarDecorator{car}}
}

func (a *AdvancedAudioDecorator) GetDescription() string {
	return a.car.GetDescription() + " with Advanced Audio"
}

func (a *AdvancedAudioDecorator) GetCost() float64 {
	return a.car.GetCost() + 800.0
}

// AirConditioner is a concrete decorator for adding a air conditioner to a car.
type AirConditionerDecorator struct {
	CarDecorator
}

func NewAirConditionerDecorator(car Car) *AirConditionerDecorator {
	return &AirConditionerDecorator{CarDecorator{car}}
}

func (a *AirConditionerDecorator) GetDescription() string {
	return a.car.GetDescription() + " with Air conditioner"
}

func (a *AirConditionerDecorator) GetCost() float64 {
	return a.car.GetCost() + 500.0
}

/* func main() {
	// Create different car types
	TLC_300 := NewSUV()
	Lancer := NewSedan()
	Solaris := NewHatchback()

	// Add optional features to cars
	hatchbackWithSunroof := NewSunroofDecorator(Solaris)
	sedanWithAdvancedAudio := NewAdvancedAudioDecorator(Lancer)
	suvWithBoth := NewAdvancedAudioDecorator(NewSunroofDecorator(TLC_300))

	// Print car descriptions and prices
	fmt.Println("Car Options:")
	fmt.Println("--------------")
	fmt.Printf("Car: %s\nPrice: $%.2f\n\n", hatchbackWithSunroof.GetDescription(), hatchbackWithSunroof.GetCost())
	fmt.Printf("Car: %s\nPrice: $%.2f\n\n", sedanWithAdvancedAudio.GetDescription(), sedanWithAdvancedAudio.GetCost())
	fmt.Printf("Car: %s\nPrice: $%.2f\n\n", suvWithBoth.GetDescription(), suvWithBoth.GetCost())
} */
