package main

import (
	"fmt"
)

// Observer represents an observer (car dealership).
type Observer interface {
	Update(message string)
}

// ConcreteObserver is a concrete implementation of the Observer interface.
type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(message string) {
	fmt.Printf("Dealer %s received message: %s\n", co.name, message)
}

// Subject represents the car buyer.
type Subject struct {
	name string
	observers []Observer
}

func NewSubject(name string) *Subject{
	return &Subject{name: name}
}

func (s *Subject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) DeregisterObserver(observer Observer) {
	indexToRemove := -1
	for i, o := range s.observers {
		if o == observer {
			indexToRemove = i
			break
		}
	}
	if indexToRemove != -1 {
		s.observers = append(s.observers[:indexToRemove], s.observers[indexToRemove+1:]...)
	}
}

func (s *Subject) NotifyObservers(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

func main() {
	
	// Create car dealerships as observers
	dealer1 := &ConcreteObserver{name: "Astana Motors"}
	dealer2 := &ConcreteObserver{name: "My Car"}

	// Create car buyers (subjects) and register them with dealerships
	salimbek := NewSubject("Salimbek")
	homyak := NewSubject("Homyak")

	salimbek.RegisterObserver(dealer1)
	homyak.RegisterObserver(dealer2)

	// Car buyer makes a purchase and notifies dealerships
	message := "Just bought a new SUV!"
	homyak.NotifyObservers(message)

	// Car buyer decides to deregister a dealership
	homyak.DeregisterObserver(dealer1)

	// Car buyer makes another purchase and notifies remaining dealership
	message = "Just bought a sedan!"
	salimbek.NotifyObservers(message)
}