package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
)

func G() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gs := runtime.ForEachG()
		goroutines := make([]*Goroutine, 0, len(gs))
		for _, g := range gs {
			goroutines = append(goroutines, &Goroutine{
				Goid: g.Goid,
				M: &Machine{
					Procid: g.M.ProcID(),
					ID:     g.M.ID(),
				},
				Waitreason:  runtime.WaitReasonStrings[g.Waitreason],
				Annotations: g.Annotations,
			})
		}
		res, err := json.Marshal(goroutines)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
