package workflow

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

type Summary struct {
	TotalRuns   int `json:"total_runs"`
	SuccessRuns int `json:"success_runs"`
}

func LoadRuns(path string) ([]WorkflowRun, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open workflow file: %w", err)
	}
	defer file.Close()

	var runs []WorkflowRun
	if err := json.NewDecoder(file).Decode(&runs); err != nil {
		return nil, fmt.Errorf("decode workflow runs: %w", err)
	}

	return runs, nil
}

func SummarizeRuns(runs []WorkflowRun) Summary {
	successCount := 0
	for _, run := range runs {
		if run.Conclusion == "success" {
			successCount++
		}
	}

	return Summary{
		TotalRuns:   len(runs),
		SuccessRuns: successCount,
	}
}
