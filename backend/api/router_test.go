package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AntohaY/pipeline-lens/backend/workflow"
)

func TestGetSummary(t *testing.T) {
	// Arrange: create the router with known data.
	summary := workflow.Summary{
		TotalRuns:   4,
		SuccessRuns: 2,
	}
	router := WorkflowRouter(summary)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/summary",
		nil,
	)
	recorder := httptest.NewRecorder()

	// Act: send the request directly through the router.
	router.ServeHTTP(recorder, request)

	// Assert: inspect the recorded response.
	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d; want %d",
			response.StatusCode, http.StatusOK)
	}

	if got := response.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q; want application/json", got)
	}

	var body map[string]int
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response JSON: %v", err)
	}

	if got := body["total_runs"]; got != 4 {
		t.Errorf("total_runs = %d; want 4", got)
	}

	if got := body["success_runs"]; got != 2 {
		t.Errorf("success_runs = %d; want 2", got)
	}
}

func TestPostSummary(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/summary",
		nil,
	)
	summary := workflow.Summary{
		TotalRuns:   4,
		SuccessRuns: 2,
	}
	router := WorkflowRouter(summary)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d; want %d",
			recorder.Code, http.StatusMethodNotAllowed)
	}
}
