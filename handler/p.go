package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
)

func P() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
					G0: &Goroutine{
						Goid: m.G0.GoID(),
					},
				}
			}

			runq := make([]Goroutine, 0, len(p.Runq))
			for _, g := range p.Runq {
				gp := g.Ptr()
				if gp == nil {
					continue
				}
				runq = append(runq, Goroutine{
					Goid:        gp.Goid,
					Waitreason:  runtime.WaitReasonStrings[gp.Waitreason],
					Annotations: gp.Annotations,
				})
			}

			var gFree []Goroutine
			for range p.GFree.N {
				gp := p.GFree.GList.Pop()
				if gp == nil {
					break
				}
				gFree = append(gFree, Goroutine{
					Goid:        gp.Goid,
					Waitreason:  runtime.WaitReasonStrings[gp.Waitreason],
					Annotations: gp.Annotations,
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
