package server

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	callHistory map[fizzBuzzParams]int
)

func init() {
	callHistory = make(map[fizzBuzzParams]int)
}

type statsResponse struct {
	Payload fizzBuzzParams `json:"parameters"`
	Calls   int            `json:"calls"`
}

// HandleStats handles stats requests
func HandleStats(w http.ResponseWriter) {
	max := -1
	var mostFrequent fizzBuzzParams
	for p, c := range callHistory {
		if c > max {
			max = c
			mostFrequent = p
		}
	}

	resp := statsResponse{mostFrequent, max}

	data, err := json.Marshal(resp)
	if err != nil {
		// silent server errors to client
		log.Printf("failed to marshal response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
