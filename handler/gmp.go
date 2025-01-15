package handler

type Goroutine struct {
	M           *Machine `json:"m"`
	Goid        uint64   `json:"goid"`
	Waitreason  string   `json:"waitreason"`
	Annotations []string `json:"annotations"`
	Status      string   `json:"status"`
}

type Machine struct {
	Procid uint64     `json:"procid"`
	Curg   *Goroutine `json:"curg"`
	P      *Processor `json:"p"`
	ID     int64      `json:"id"`
}

type Processor struct {
	ID    int32       `json:"id"`
	M     *Machine    `json:"m"`
	Runq  []Goroutine `json:"runq"`
	GFree []Goroutine `json:"gfree"`
}

func String(i uint32) string {
	switch i {
	case 0:
		return "idle"
	case 1:
		return "runnable"
	case 2:
		return "running"
	case 3:
		return "syscall"
	case 4:
		return "wating"
	case 5:
		return "unused"
	case 6:
		return "dead"
	case 7:
		return "unused"
	case 8:
		return "copy stack"
	case 9:
		return "preempted"
	}
	return "unhandled"
}
