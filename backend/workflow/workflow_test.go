package workflow

import "testing"

func TestSummarizeRuns(t *testing.T) {
	tests := []struct {
		name string
		runs []WorkflowRun
		want Summary
	}{
		{
			name: "mixed outcomes",
			runs: []WorkflowRun{
				{Conclusion: "success"},
				{Conclusion: "failure"},
				{Conclusion: "success"},
				{Conclusion: "cancelled"},
			},
			want: Summary{TotalRuns: 4, SuccessRuns: 2},
		},
		{
			name: "empty input",
			runs: []WorkflowRun{},
			want: Summary{TotalRuns: 0, SuccessRuns: 0},
		},
		{
			name: "all successful",
			runs: []WorkflowRun{
				{Conclusion: "success"},
				{Conclusion: "success"},
				{Conclusion: "success"},
			},
			want: Summary{TotalRuns: 3, SuccessRuns: 3},
		},
		{
			name: "no successful runs",
			runs: []WorkflowRun{
				{Conclusion: "failure"},
				{Conclusion: "cancelled"},
			},
			want: Summary{TotalRuns: 2, SuccessRuns: 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SummarizeRuns(tc.runs)

			if got != tc.want {
				t.Errorf("SummarizeRuns() = %+v; want %+v", got, tc.want)
			}
		})
	}
}
