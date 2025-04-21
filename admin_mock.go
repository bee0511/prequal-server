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
			{"id": 1, "name": "u1", "url": "http://user1", "status": "active"},
			{"id": 2, "name": "u2", "url": "http://user2", "status": "active"},
			{"id": 3, "name": "u3", "url": "http://user3", "status": "active"},
			{"id": 4, "name": "u4", "url": "http://user4", "status": "active"},
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
