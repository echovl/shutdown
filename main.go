package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var mustShutdownNow = false

type Response struct {
	MustShutdownNow bool `json:"must_shutdown_now"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			mustShutdownNow = true
		} else if r.Method == http.MethodDelete {
			mustShutdownNow = false
		}

		resp := Response{mustShutdownNow}

		fmt.Println(resp)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

    fmt.Println("Server running at port 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
