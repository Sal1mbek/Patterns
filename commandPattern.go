package main

import "fmt"

// CarCommand is the command interface.
type CarCommand interface {
	Execute()
}

// Cars represents the receiver.
type Cars struct {
	isOpen bool
}

func (c *Cars) Open(key string) {
	c.isOpen = true
	fmt.Printf("Car is open with %s\n", key)
}

func (c *Cars) Close(key string) {
	c.isOpen = false
	fmt.Printf("Car is closed with %s\n", key)
}

// OpenCarCommand is a concrete command for opening the car.
type OpenCarCommand struct {
	carName string
	car     *Cars
}

func (o *OpenCarCommand) Execute() {
	o.car.Open(o.carName)
}

// CloseCarCommand is a concrete command for closing the car.
type CloseCarCommand struct {
	carName string
	car     *Cars
}

func (c *CloseCarCommand) Execute() {
	c.car.Close(c.carName)
}

// CarInvoker represents the car invoker.
type CarInvoker struct {
	command CarCommand
}

func (c *CarInvoker) SetCommand(command CarCommand) {
	c.command = command
}

func (c *CarInvoker) ExecuteCommand() {
	c.command.Execute()
}

/* func main() {
	car := &Cars{}
	openCommand := &OpenCarCommand{car}
	closeCommand := &CloseCarCommand{car}

	remoteControl := &RemoteControl{}
	remoteControl.SetOpenCommand(openCommand)
	remoteControl.SetCloseCommand(closeCommand)

	remoteControl.OpenCar()
	remoteControl.CloseCar()

	// Можно использовать отмену (Undo)
	remoteControl.closeCommand.Undo()
} */
