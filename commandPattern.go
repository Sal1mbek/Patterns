package main

import "fmt"

// CarCommand provides a command interface.
type CarCommand interface {
	Execute()
	Undo()
}

// OpenCarCommand implements the CarCommand interface.
type OpenCarCommand struct {
	car *Cars
}

func (c *OpenCarCommand) Execute() {
	c.car.Open()
}

func (c *OpenCarCommand) Undo() {
	c.car.Close()
}

// CloseCarCommand implements the CarCommand interface.
type CloseCarCommand struct {
	car *Cars
}

func (c *CloseCarCommand) Execute() {
	c.car.Close()
}

func (c *CloseCarCommand) Undo() {
	c.car.Open()
}

// Receiver implementation.
type Cars struct {
	isOpen bool
}

func (c *Cars) Open() {
	c.isOpen = true
	fmt.Println("Car is open")
}

func (c *Cars) Close() {
	c.isOpen = false
	fmt.Println("Car is closed")
}

// Invoker implementation.
type RemoteControl struct {
	openCommand CarCommand
	closeCommand CarCommand
}

func (rc *RemoteControl) SetOpenCommand(command CarCommand) {
	rc.openCommand = command
}

func (rc *RemoteControl) SetCloseCommand(command CarCommand) {
	rc.closeCommand = command
}

func (rc *RemoteControl) OpenCar() {
	rc.openCommand.Execute()
}

func (rc *RemoteControl) CloseCar() {
	rc.closeCommand.Execute()
}

/* func main() {
	car := &Car{}
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
