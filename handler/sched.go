package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

type Scheduler struct {
	Runq    []Goroutine `json:"runq"`
	Stack   []Goroutine `json:"stack"`
	NoStack []Goroutine `json:"noStack"`
}

func Sched() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sched := runtime.Sched()
		var scheduler Scheduler
		var runq []Goroutine
		for !sched.Runq.Empty() {
			gp := sched.Runq.Pop()
			if gp == nil {
				break
			}
			runq = append(runq, Goroutine{
				Goid:        gp.Goid,
				Waitreason:  runtime.WaitReasonStrings[gp.Waitreason],
				Annotations: gp.Annotations,
			})
		}
		scheduler.Runq = runq

		var stack, noStack []Goroutine
		for range sched.GFree.N {
			sgp := sched.GFree.Stack.Pop()
			if sgp == nil {
				break
			}
			stack = append(stack, Goroutine{
				Goid:        sgp.Goid,
				Waitreason:  runtime.WaitReasonStrings[sgp.Waitreason],
				Annotations: sgp.Annotations,
			})

			nsgp := sched.GFree.NoStack.Pop()
			if nsgp == nil {
				break
			}
			noStack = append(noStack, Goroutine{
				Goid:        nsgp.Goid,
				Waitreason:  runtime.WaitReasonStrings[nsgp.Waitreason],
				Annotations: nsgp.Annotations,
			})
		}
		scheduler.Stack = stack
		scheduler.NoStack = noStack

		res, err := json.Marshal(scheduler)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal response: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
