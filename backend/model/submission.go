package model

// SubmitRequest represents an incoming code submission.
type SubmitRequest struct {
	Code      string `json:"code" binding:"required"`
	Language  string `json:"language" binding:"required"`
	ProblemID string `json:"problem_id" binding:"required"`
}

// SubmitResponse represents the result of a submission.
type SubmitResponse struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	ProblemID string `json:"problem_id"`
	Language  string `json:"language"`
	Message   string `json:"message,omitempty"`
}
