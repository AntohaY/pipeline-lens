package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type WorkflowRun struct {
	ID         int    `json:"id"`
	Workflow   string `json:"workflow"`
	Conclusion string `json:"conclusion"`
}

func main() {
	file, err := os.Open("testdata/runs.json")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var runs []WorkflowRun
	if err := json.NewDecoder(file).Decode(&runs); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	successCount := 0
	for _, run := range runs {
		if run.Conclusion == "success" {
			successCount++
		}
	}

	fmt.Printf("Total runs: %d\n", len(runs))
	fmt.Printf("Successful runs: %d\n", successCount)
}
