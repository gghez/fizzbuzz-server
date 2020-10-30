package server

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	callHistory map[Payload]int
)

func init() {
	callHistory = make(map[Payload]int)
}

type statsResponse struct {
	Payload Payload `json:"parameters"`
	Calls   int     `json:"calls"`
}

// HandleStatS handles stats requests
func HandleStats(w http.ResponseWriter) {
	max := -1
	var mostFrequent Payload
	for p, c := range callHistory {
		if c > max {
			max = c
			mostFrequent = p
		}
	}

	resp := statsResponse{mostFrequent, max}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("failed to marshal response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
