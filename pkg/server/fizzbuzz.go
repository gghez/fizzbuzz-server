package server

import (
	"encoding/json"
	"fizzbuzz/pkg/fizzbuzz"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Payload struct {
	Int1, Int2, Limit int
	Str1, Str2        string
}

// HandleFizzBuzz handles fizzbuzz requests
func HandleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var p Payload
	if err := json.Unmarshal(data, &p); err != nil {
		w.Write([]byte(fmt.Sprintf("failed to parse JSON body: %v", err)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	callHistory[p]++

	results := fizzbuzz.Compute(p.Int1, p.Int2, p.Limit, p.Str1, p.Str2)
	jsonResp, err := json.Marshal(results)
	if err != nil {
		log.Printf("failed to marshal results: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonResp)
}
