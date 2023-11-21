package main

import (
  	"fmt"
)


func main() {

	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%s\n", &name)
	
	var cash float64
	fmt.Println("Enter your cash: ($)")
	fmt.Scanf("%f\n", &cash)
	
	fmt.Printf("\nWelcome to car factory project, %s!!!. Your cash - $%.2f.\n", name, cash)

	buyer := &CarBuyer{}  
	//Using Strategy pattern
	fmt.Println("\nWhich car you are going to buy?")
	fmt.Println("\n 1) SUV\n", "2) Crossover\n", "3) Wagon\n", "4) Minivan\n", "5) Sedan\n", "6) Hatchback\n", "7) Pick-UP\n")

	var answer int
	fmt.Println("Enter your car choice:")
	fmt.Scanf("%d\n", &answer)

	var myCar Car
  	switch answer {
 
    case 1:
		myCar = NewSUV()
    	buyer.SetStrategy(&SUVStrategy{})	
  	case 2:
		myCar = NewCrossover()
    	buyer.SetStrategy(&CrossoverStrategy{})
  	case 3:
		myCar = NewWagon()
    	buyer.SetStrategy(&WagonStrategy{})
    case 4:
		myCar = NewMinivan()
    	buyer.SetStrategy(&MinivanStrategy{})
    case 5:
		myCar = NewSedan()
    	buyer.SetStrategy(&SedanStrategy{})
    case 6:
		myCar = NewHatchback()
    	buyer.SetStrategy(&HatchbackStrategy{})
    case 7:
		myCar = NewPickUp()
    	buyer.SetStrategy(&PickUPStrategy{})
    default:
        fmt.Println("undefined")
    }
	
	fmt.Printf("Car: %s\nPrice: $%.2f\n", myCar.GetDescription(), myCar.GetCost())

	fmt.Println("\nOk, let's now choose an heart(engine) for your car")
	fmt.Println("1) Engine (petrol/diesel)\n2) Electric (Powerbank ☺)\n3) Hybrid (such as Toyota Prius)")

	var engineChoice int
	fmt.Println("\nEnter your engine choice:")
	fmt.Scanf("%d\n", &engineChoice)
	
	chooseOfEngine(engineChoice)

	if cash >= myCar.GetCost(){

		fmt.Println("\nIn our project we also have a function to add optional features to cars")
		fmt.Println("\n 1)Sunroof/Люк на крыше ($1000.0)\n", "2)AdvancedAudio ($800.0)\n", "3)AirConditioner ($500.0)\n", "4)Sunroof+AdvencedAudio ($1800.0)\n", "5)Sunroof+AirConditioner ($1500.0)\n", "6)AdvencedAudio+AirConditioner ($1300.0)\n", "7)All of them:Sunroof, AdvencedAudio, AirConditioner ($2300.0)")
		fmt.Println("\nDo you want to add optional features to car? (yes/no)")
		var userWish string
		fmt.Scanf("%s\n", &userWish)
		if userWish == "yes" {
			fmt.Println("\nPlease enter which option you want to add: (1-7)")
			var decorAns int
			fmt.Scanf("%d\n", &decorAns)
			decorator(myCar, decorAns, cash)
			fmt.Println(buyer.BuyCar())
		} else if userWish == "no" {
			fmt.Println("Ok, thank you!")
			fmt.Println(buyer.BuyCar())
			fmt.Printf("Your cash - $%.2f.\n", cash-myCar.GetCost())	
		} else {
			fmt.Println("undefined")
		}

		fmt.Println("Congratulations on your purchase of the car. Будь внимательным на дороге, и пристегни ремень! Here are your car keys")
		// User's choice of opening and closing
		fmt.Println("\nChoose an option:")
		fmt.Println("1) Open Car with Mechanical Key")
		fmt.Println("2) Open Car with Remote Key")
		fmt.Println("3) Close Car with Mechanical Key")
		fmt.Println("4) Close Car with Remote Key")

		var userChoice int
		fmt.Println("\nEnter your choice:")
		fmt.Scanf("%d\n", &userChoice)

		InteractWithCar(userChoice)
	} else {
		fmt.Printf("\nYou don't have enough money!!\nYour cash - $%.2f.\n", cash)
	}

}

// Using decorator pattern to add optional features to your car
func decorator(myCar Car, choice int, cash float64) {
	switch choice {
	case 1:
		myCarWithSunroof := NewSunroofDecorator(myCar)
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myCarWithSunroof.GetDescription(), myCarWithSunroof.GetCost())
		fmt.Printf("Your cash - $%.2f.\n", cash - myCarWithSunroof.GetCost())	
	case 2:
		myCarWithAdvancedAudio := NewAdvancedAudioDecorator(myCar)
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myCarWithAdvancedAudio.GetDescription(), myCarWithAdvancedAudio.GetCost())
		fmt.Printf("Your cash - $%.2f.\n", cash - myCarWithAdvancedAudio.GetCost())	
	case 3:
		myCarWithAirConditioner := NewAirConditionerDecorator(myCar)
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myCarWithAirConditioner.GetDescription(), myCarWithAirConditioner.GetCost())
		fmt.Printf("Your cash - $%.2f.\n", cash - myCarWithAirConditioner.GetCost())	
	case 4:
		myCarWith_Sunroof_and_AdvancedAudio := NewAdvancedAudioDecorator(NewSunroofDecorator(myCar))
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myCarWith_Sunroof_and_AdvancedAudio.GetDescription(), myCarWith_Sunroof_and_AdvancedAudio.GetCost())
		fmt.Printf("Your cash - $%.2f.\n", cash - myCarWith_Sunroof_and_AdvancedAudio.GetCost())	
	case 5:
		myCarWith_Sunroof_and_AirConditioner := NewSunroofDecorator(NewAirConditionerDecorator(myCar))
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myCarWith_Sunroof_and_AirConditioner.GetDescription(), myCarWith_Sunroof_and_AirConditioner.GetCost())
		fmt.Printf("Your cash - $%.2f.\n", cash - myCarWith_Sunroof_and_AirConditioner.GetCost())	
	case 6:
		myCarWith_AdvancedAudio_and_AirConditioner := NewAdvancedAudioDecorator(NewAirConditionerDecorator(myCar))
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myCarWith_AdvancedAudio_and_AirConditioner.GetDescription(), myCarWith_AdvancedAudio_and_AirConditioner.GetCost())
		fmt.Printf("Your cash - $%.2f.\n", cash - myCarWith_AdvancedAudio_and_AirConditioner.GetCost())	
	case 7:
		myCarWith_allOptions := NewSunroofDecorator(NewAdvancedAudioDecorator(NewAirConditionerDecorator(myCar)))
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myCarWith_allOptions.GetDescription(), myCarWith_allOptions.GetCost())
		fmt.Printf("Your cash - $%.2f.\n", cash - myCarWith_allOptions.GetCost())	
	default:
		fmt.Println("undefined")			
	}
}

// Choosing heart(двигатель) of your car. Using Factory method
func chooseOfEngine(variant int){
	switch variant {
	case 1:
		dvsFactory := EngineCarFactory{}
		dvsCar := dvsFactory.CreateCar()
		dvsCar.Drive()
	case 2:
		powerbankFactory := ElectricCarFactory{}
		powerbank := powerbankFactory.CreateCar()
		powerbank.Drive()
	case 3:
		hybridFactory := HybridCarFactory{}
		hybridCar := hybridFactory.CreateCar()
		hybridCar.Drive()
	default:
		fmt.Println("undefined")
	}
}
// Using command pattern
func InteractWithCar(userChoice int) {
	// Create the receiver
	car := &Cars{}

	var carName string

	switch userChoice {
	case 1, 3:
		carName = "Mechanical Key"
	case 2, 4:
		carName = "Remote Key"
	}

	var command CarCommand

	switch userChoice {
	case 1, 2:
		command = &OpenCarCommand{carName: carName, car: car}
	case 3, 4:
		command = &CloseCarCommand{carName: carName, car: car}
	}

	// Create the car invoker
	carInvoker := &CarInvoker{}
	carInvoker.SetCommand(command)

	carInvoker.ExecuteCommand()
}
