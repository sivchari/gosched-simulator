package handler

type Goroutine struct {
	M           *Machine
	Goid        uint64
	Waitreason  string
	Annotations []string
}

type Machine struct {
	G0     *Goroutine
	Procid uint64
	Curg   *Goroutine
	P      *Processor
	ID     int64
}

type Processor struct {
	ID      int32
	M       *Machine
	Runq    []Goroutine
	GFree   []Goroutine
	Runnext *Goroutine
}
