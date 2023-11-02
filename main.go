package main

import (
	"fmt"
)

func main() {
	
	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%s\n", &name)
 
	var cash int
	fmt.Println("Enter your cash: ($)")
	fmt.Scanf("%d\n", &cash)
 
	fmt.Printf("Welcome to car factory project, %s, your cash - %d $\n", name, cash)

	buyer := &CarBuyer{}	
	//Using Strategy pattern
	fmt.Println("Which car you are going to buy?")
	fmt.Println("1) SUV\n", "2) Crossover\n", "3) Wagon\n", "4) Minivan\n", "5) Sedan\n", "6) Hatchback\n", "7) Pick-UP\n")

	var answer int
	fmt.Scanf("%d\n", &answer)

	mySUV := NewSUV()
	myCrossover := NewCrossover()
	myWagon := NewWagon()
	myMinivan := NewMinivan()
	mySedan := NewSedan()
	myHatchback := NewHatchback()
	myPickUp := NewPickUp()

	switch answer {
 
    case 1:
		buyer.SetStrategy(&SUVStrategy{})
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", mySUV.GetDescription(), mySUV.GetCost())

	case 2:
		buyer.SetStrategy(&CrossoverStrategy{})
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myCrossover.GetDescription(), myCrossover.GetCost())

	case 3:
		buyer.SetStrategy(&WagonStrategy{})
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myWagon.GetDescription(), myWagon.GetCost())

    case 4:
		buyer.SetStrategy(&MinivanStrategy{})	
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myMinivan.GetDescription(), myMinivan.GetCost())

    case 5:
		buyer.SetStrategy(&SedanStrategy{})	
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", mySedan.GetDescription(), mySedan.GetCost())

    case 6:
		buyer.SetStrategy(&HatchbackStrategy{})	
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myHatchback.GetDescription(), myHatchback.GetCost())

    case 7:
		buyer.SetStrategy(&PickUPStrategy{})	
		fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myPickUp.GetDescription(), myPickUp.GetCost())

    default:
        fmt.Println("неизвестно")
    }

	fmt.Println(buyer.BuyCar())
	fmt.Printf("Car: %s\nPrice: $%.2f\n\n", myPickUp.GetDescription(), myPickUp.GetCost())

}

		
	/* carFactory := getInstance()

	carFactory.CreateCar("SUV")
	carFactory.CreateCar("Sedan")
	carFactory.CreateCar("Hatchback")
	carFactory.CreateCar("Crossover")
	carFactory.CreateCar("Wagon")
	carFactory.CreateCar("Minivan")
	carFactory.CreateCar("Pick-UP") */

