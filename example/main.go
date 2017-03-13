package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func firstHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"A": map[string][]string{
			"B": []string{"foo", "bar"},
		},
	}

	body, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func secondHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"key": 1,
		"foo": "bar",
		"stats": map[string]interface{}{
			"updated_at": time.Now().Unix(),
			"created_at": 1489345012,
		},
	}

	body, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func main() {
	http.HandleFunc("/first", firstHandler)
	http.HandleFunc("/second", secondHandler)
	http.ListenAndServe(":8080", nil)
}
