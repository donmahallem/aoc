package day20

type module interface {
	Receive(from int, pulse bool) (bool, bool)
	TargetIds() []int
	Reset()
}

type conjunctionModule struct {
	targetIds []int
	inputs    map[int]bool
}

func (c *conjunctionModule) Receive(from int, pulse bool) (bool, bool) {
	if c.inputs == nil {
		c.inputs = make(map[int]bool)
	}
	c.inputs[from] = pulse

	allHigh := true
	for _, state := range c.inputs {
		if !state {
			allHigh = false
			break
		}
	}

	return true, !allHigh
}

func (c *conjunctionModule) TargetIds() []int {
	return c.targetIds
}

func (c *conjunctionModule) Reset() {
	for input := range c.inputs {
		c.inputs[input] = false
	}
}

func (c *conjunctionModule) AddInput(id int) {
	if c.inputs == nil {
		c.inputs = make(map[int]bool)
	}
	c.inputs[id] = false
}

type flipFlopModule struct {
	state     bool
	targetIds []int
}

func (f *flipFlopModule) Receive(_ int, pulse bool) (bool, bool) {
	if pulse {
		return false, false
	}

	f.state = !f.state
	return true, f.state
}

func (f *flipFlopModule) TargetIds() []int {
	return f.targetIds
}

func (f *flipFlopModule) Reset() {
	f.state = false
}

func (f *flipFlopModule) State() bool {
	return f.state
}

type broadcasterModule struct {
	targetIds []int
}

func (b *broadcasterModule) TargetIds() []int {
	return b.targetIds
}
