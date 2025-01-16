package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

func P() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: make goroutine by user action.
		runtime.GC()
		for range 10 {
			go func() {
				time.Sleep(5 * time.Second)
			}()
		}

		ps := runtime.ForEachP()
		processors := make([]*Processor, 0, len(ps))

		for _, p := range ps {
			processor := &Processor{
				ID: p.ID,
			}
			m := p.M.Ptr()
			if m != nil {
				processor.M = &Machine{
					ID: m.ID,
					Curg: &Goroutine{
						Goid: m.Curg.GoID(),
					},
				}
			}
			runq := make([]Goroutine, 0, len(p.Runq))
			for _, gp := range p.XRunq {
				runq = append(runq, Goroutine{
					Goid:        gp.Goid,
					Waitreason:  runtime.WaitReasonStrings[gp.Waitreason],
					Annotations: gp.Annotations,
					Status:      String(gp.Atomicstatus.Load()),
				})
			}

			var gFree []Goroutine
			for range p.GFree.N {
				gp := p.GFree.GList.Pop()
				if gp == nil {
					continue
				}
				gFree = append(gFree, Goroutine{
					Goid:        gp.Goid,
					Waitreason:  runtime.WaitReasonStrings[gp.Waitreason],
					Annotations: gp.Annotations,
					Status:      String(gp.Atomicstatus.Load()),
				})
			}
			processor.GFree = gFree
			processor.Runq = runq
			processors = append(processors, processor)
		}
		res, err := json.Marshal(processors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
