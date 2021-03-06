package medtronic

const (
	Suspend Command = 0x4D
)

func (pump *Pump) Suspend(suspend bool) {
	if suspend {
		pump.Execute(Suspend, 1)
	} else {
		pump.Execute(Suspend, 0)
	}
}
