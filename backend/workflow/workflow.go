package workflow

type WorkflowRun struct {
	ID         int    `json:"id"`
	Workflow   string `json:"workflow"`
	Conclusion string `json:"conclusion"`
}

type workflowRuns interface {
	func loadRuns(path string) ([]WorkflowRun, error)
	func summarizeRuns(runs []WorkflowRun) Summary
}
func loadRuns(path string) ([]WorkflowRun, error)
func summarizeRuns(runs []WorkflowRun) Summary
