package db

import (
	"context"
	"encoding/json"
)

type Submission struct {
	ID           string
	ProblemID    string
	Language     string
	Source       string
	Verdict      string
	Points       float64
	TotalPoints  float64
	TotalTime    float64
	MaxMemory    int64
	CompileError string
	Cases        []byte
	IP           string
}

func InsertSubmission(ctx context.Context, s *Submission) error {
	_, err := Pool.Exec(ctx, `
		INSERT INTO submissions (id, problem_id, language, source, verdict, points, total_points, total_time, max_memory, compile_error, cases, ip)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		s.ID, s.ProblemID, s.Language, s.Source, s.Verdict,
		s.Points, s.TotalPoints, s.TotalTime, s.MaxMemory,
		s.CompileError, s.Cases, s.IP,
	)
	return err
}

func MarshalCases(cases any) []byte {
	b, err := json.Marshal(cases)
	if err != nil {
		return []byte("[]")
	}
	return b
}
