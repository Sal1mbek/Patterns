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
	
	fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myCar.GetDescription(), myCar.GetCost())

	if cash >= myCar.GetCost(){
		fmt.Println(buyer.BuyCar())
		fmt.Printf("Your cash - $%.2f.", cash-myCar.GetCost())

		fmt.Println("\nIn our project we also have a function to add optional features to cars")
	fmt.Println("\n 1)Sunroof/Люк на крыше ($1000.0)\n", "2)AdvancedAudio ($800.0)\n", "3)AirConditioner ($500.0)\n", "4)Sunroof+AdvencedAudio ($1800.0)\n", "5)Sunroof+AirConditioner ($1500.0)\n", "6)AdvencedAudio+AirConditioner ($1300.0)\n", "7)All of them:Sunroof, AdvencedAudio, AirConditioner ($2300.0)")
	fmt.Println("\nDo you want to add optional features to car? (yes/no)")
	var userWish string
	fmt.Scanf("%s\n", &userWish)
	if userWish == "yes" {
		fmt.Println("\nPlease enter which option you want to add: (1-7)")
		var decorAns int
		fmt.Scanf("%d\n", &decorAns)
		switch decorAns {
		case 1:
		sedanWithSunroof := NewSunroofDecorator(myCar)
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", sedanWithSunroof.GetDescription(), sedanWithSunroof.GetCost())
		case 2:
		sedanWithAdvancedAudio := NewAdvancedAudioDecorator(myCar)
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", sedanWithAdvancedAudio.GetDescription(), sedanWithAdvancedAudio.GetCost())
		case 3:
		sedanWithAirConditioner := NewAirConditionerDecorator(myCar)
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", sedanWithAirConditioner.GetDescription(), sedanWithAirConditioner.GetCost())
		case 4:
		sedanWith_Sunroof_and_AdvancedAudio := NewAdvancedAudioDecorator(NewSunroofDecorator(myCar))
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", sedanWith_Sunroof_and_AdvancedAudio.GetDescription(), sedanWith_Sunroof_and_AdvancedAudio.GetCost())
		case 5:
		sedanWith_Sunroof_and_AirConditioner := NewSunroofDecorator(NewAirConditionerDecorator(myCar))
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", sedanWith_Sunroof_and_AirConditioner.GetDescription(), sedanWith_Sunroof_and_AirConditioner.GetCost())
		case 6:
		sedanWith_AdvancedAudio_and_AirConditioner := NewAdvancedAudioDecorator(NewAirConditionerDecorator(myCar))
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", sedanWith_AdvancedAudio_and_AirConditioner.GetDescription(), sedanWith_AdvancedAudio_and_AirConditioner.GetCost())
		case 7:
		sedanWith_allOptions := NewSunroofDecorator(NewAdvancedAudioDecorator(NewAirConditionerDecorator(myCar)))
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", sedanWith_allOptions.GetDescription(), sedanWith_allOptions.GetCost())
		default:
				fmt.Println("undefined")
		
		}
	} else if userWish == "no" {
		fmt.Println("Ok, thank you!")
	} else {
        fmt.Println("undefined")
	}
	} else {
		fmt.Printf("You don't have enough money!!\nYour cash - $%.2f.", cash)
	}

	//Using decorator pattern
	
	

	// Using command pattern

	/* openCommand := &OpenCarCommand{myCar}
	closeCommand := &CloseCarCommand{myCar}

	remoteControl := &RemoteControl{}
	remoteControl.SetOpenCommand(openCommand)
	remoteControl.SetCloseCommand(closeCommand)

	remoteControl.OpenCar()
	remoteControl.CloseCar()

	// Можно использовать отмену (Undo)
	remoteControl.closeCommand.Undo() */

}
