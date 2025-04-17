// backend/backend.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type PingResponse struct {
	ServerName       string `json:"serverName"`
	RequestsInFlight int    `json:"requestInFlight"`
	Latency          int    `json:"latency"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1233"
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		latency := 50
		time.Sleep(time.Duration(latency) * time.Millisecond)

		resp := PingResponse{
			ServerName:       "server-" + port,
			RequestsInFlight: 1,
			Latency:          latency,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Printf("Backend server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
