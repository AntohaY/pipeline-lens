package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AntohaY/pipeline-lens/backend/workflow"
)

func WorkflowRouter(summary workflow.Summary) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/summary", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(summary); err != nil {
			log.Printf("write summary response: %v", err)
		}
	})

	return mux
}
