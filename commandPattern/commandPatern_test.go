package commandPattern

import "testing"

func TestCommandPattern(t *testing.T) {
	lightOn := new(LightOnCommand)
	lightOff := new(LightOffCommand)

	garageDoorOpen := new(GarageDoorOpenCommand)
	garageDoorClose := new(GarageDoorCloseCommand)

	controller := new(RemoteController)
	controller.SetCommand(0, lightOn, lightOff)
	controller.SetCommand(1, garageDoorOpen, garageDoorClose)

	controller.On(0)
	controller.Off(0)
	controller.Undo()
	controller.On(1)
	controller.Off(1)
	controller.Undo()
	// output
	// light is on
	// light is off
	// light is on
	// garageDoor is open
	// garageDoor is close
	// garageDoor is open
}
