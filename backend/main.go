package main

import (
	"fmt"
	"os"

	"example.com/pipeline-lens/backend/workflow"
)

func main() {
	runs, err := workflow.LoadRuns("testdata/runs.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	summary := workflow.SummarizeRuns(runs)

	fmt.Printf("Total runs: %d\n", summary.TotalRuns)
	fmt.Printf("Successful runs: %d\n", summary.SuccessRuns)
}
