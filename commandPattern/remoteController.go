package commandPattern

type RemoteController struct {
	onSlot      [2]Command
	offSlot     [2]Command
	unDoCommand Command
}

func (r *RemoteController) SetCommand(slotId int, onCommand, offCommand Command) {
	r.onSlot[slotId] = onCommand
	r.offSlot[slotId] = offCommand
}

func (r *RemoteController) On(slotId int) {
	r.onSlot[slotId].Execute()
	r.unDoCommand = r.onSlot[slotId]
}

func (r *RemoteController) Off(slotId int) {
	r.offSlot[slotId].Execute()
	r.unDoCommand = r.offSlot[slotId]
}

func (r *RemoteController) Undo() {
	r.unDoCommand.Undo()
}
