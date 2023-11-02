package main


// CarBuyingStrategy defines the strategy interface for buying a car.
type CarBuyingStrategy interface {
	BuyCar() string
}

// SUVStrategy-buying an SUV.
type SUVStrategy struct{}

func (s *SUVStrategy) BuyCar() string {
	return "Buying an SUV"
}

// CrossoverStrategy-buying an crossover.
type CrossoverStrategy struct{}

func (c *CrossoverStrategy) BuyCar() string {
	return "Buying an crossover"
}

// WagonStrategy-buying a wagon.
type WagonStrategy struct{}

func (w *WagonStrategy) BuyCar() string {
	return "Buying a wagon car"
}

// MinivanStrategy-buying a minivan.
type MinivanStrategy struct{}

func (m *MinivanStrategy) BuyCar() string {
	return "Buying a minivan"
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

// PickUPStrategy-buying a pick-UP.
type PickUPStrategy struct{}

func (p *PickUPStrategy) BuyCar() string {
	return "Buying a Pick-UP"
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

/* func main() {
	buyer := &CarBuyer{}
	// Buying a sedan
	
	buyer.SetStrategy(&SedanStrategy{})
	fmt.Println(buyer.BuyCar())
} */
