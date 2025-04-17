package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/admin/get-replica", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"data": []map[string]interface{}{
				{"id": 1, "name": "r1", "url": "http://localhost:9001", "status": "active"},
				{"id": 2, "name": "r2", "url": "http://localhost:9002", "status": "active"},
				{"id": 3, "name": "r3", "url": "http://localhost:9003", "status": "active"},
				{"id": 4, "name": "r4", "url": "http://localhost:9004", "status": "active"},
			},
		})
	})

	http.HandleFunc("/admin/get-prequal-parameters", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"id":                  1,
				"max_life_time":       1,
				"pool_size":           16,
				"probe_factor":        1.2,
				"probe_remove_factor": 1,
				"mu":                  1,
			},
		})
	})

	http.ListenAndServe(":8080", nil)
}
