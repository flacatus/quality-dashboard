package v1alpha1

import "time"

// Indicate the test suites specs for a job. Need to be in XUNIT format
type JobSuites struct {
	// Job ID identification
	JobID string `json:"job_id"`

	// Test name for a specific case
	TestCaseName string `json:"test_name"`

	// Indicate the status for a specific tests
	TestCaseStatus string `json:"test_status"`

	// Return the total time for a test
	TestTiming float64 `json:"test_timing"`

	// Return if is postsubmit, presubmit or periodic
	JobType string `json:"job_type"`

	CreatedAt time.Time `json:"created_at"`

	ErrorMessage string `json:"error_message"`
}

type SuitesFailureFrequency struct {
	Name string `json:"name"`

	Status string `json:"status"`

	ErrorMessage string `json:"error_message"`

	Count int `json:"count"`
}
