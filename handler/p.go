package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"fmt"
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
				processor.M =  &Machine{
					ID: m.ID,
				}
			}
			runq := make([]Goroutine, 0, len(p.Runq))
			for _, g := range p.Runq {
				if g != 0 {
					gp := g.Ptr()		
					runq = append(runq, Goroutine{
						Goid: gp.Goid,
						Waitreason: runtime.WaitReasonStrings[gp.Waitreason],
						Annotations: gp.Annotations,
			})
				}
			}
			processor.Runq = runq
			processors = append(processors, processor)
		}
		res, err := json.Marshal(processors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(string(res))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
