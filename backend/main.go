package main

import (
	"log"
	"net/http"

	"github.com/AntohaY/pipeline-lens/backend/api"
	"github.com/AntohaY/pipeline-lens/backend/workflow"
)

func main() {
	runs, err := workflow.LoadRuns("testdata/runs.json")
	if err != nil {
		log.Fatal(err)
	}

	summary := workflow.SummarizeRuns(runs)
	router := api.WorkflowRouter(summary)

	log.Println("Listening on http://127.0.0.1:8080")

	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		log.Fatal(err)
	}
}
