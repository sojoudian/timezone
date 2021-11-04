package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/")
	q := r.URL.Query()
	m := make(map[string]interface{})

	m["time"] = time.Now().UTC()
	// if ?tz=Time/Zone in Query params, update
	if tz := q.Get("tz"); tz != "" {
		if loc, err := time.LoadLocation(tz); err == nil {
			m["time"] = time.Now().In(loc)
		} else {
			m["error"] = "unknown timezone"
			m["time"] = nil
		}
	}

	json.NewEncoder(w).Encode(m)
}
func main() {
	http.HandleFunc("/time", timeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
