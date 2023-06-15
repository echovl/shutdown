package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

	fmt.Println(fmt.Sprintf("Server running at port %s", os.Getenv("PORT")))

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")), nil); err != nil {
		log.Fatal(err)
	}
}
