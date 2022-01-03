package commandPattern

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type Light struct {
}

func (*Light) On() {
	fmt.Println("light is on")
}

func (*Light) Off() {
	fmt.Println("light is off")
}

type LightOnCommand struct {
	light Light
}

func (l *LightOnCommand) Undo() {
	l.light.Off()
}

func (l *LightOnCommand) Execute() {
	l.light.On()
}

type LightOffCommand struct {
	light Light
}

func (l *LightOffCommand) Execute() {
	l.light.Off()
}

func (l *LightOffCommand) Undo() {
	l.light.On()
}

type Garage struct {
}

func (*Garage) Open() {
	fmt.Println("garageDoor is open")
}

func (*Garage) Close() {
	fmt.Println("garageDoor is close")
}

type GarageDoorOpenCommand struct {
	garage Garage
}

func (g *GarageDoorOpenCommand) Execute() {
	g.garage.Open()
}

func (g *GarageDoorOpenCommand) Undo() {
	g.garage.Close()
}

type GarageDoorCloseCommand struct {
	garage Garage
}

func (g *GarageDoorCloseCommand) Undo() {
	g.garage.Open()
}

func (g *GarageDoorCloseCommand) Execute() {
	g.garage.Close()
}
