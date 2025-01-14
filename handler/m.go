package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
)

func M() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ms := runtime.ForEachM()
		machines := make([]*Machine, 0, len(ms))
		for _, m := range ms {
			machine := &Machine{
				ID: m.ID,
				Curg: &Goroutine{
					Goid: m.Curg.GoID(),
				},
			}
			machines = append(machines, machine)
		}
		res, err := json.Marshal(machines)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
