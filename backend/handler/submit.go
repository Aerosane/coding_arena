package handler

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/Aerosane/coding_arena/backend/model"
	"github.com/gin-gonic/gin"
)

var supportedLanguages = map[string]bool{
	"python": true,
	"cpp":    true,
	"c":      true,
	"java":   true,
	"go":     true,
}

// Submit handles POST /submit — accepts code, language, and problem ID.
func Submit(c *gin.Context) {
	var req model.SubmitRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing or invalid fields: code, language, and problem_id are required",
		})
		return
	}

	if !supportedLanguages[req.Language] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("unsupported language: %s", req.Language),
		})
		return
	}

	// Generate a collision-safe submission ID
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate submission ID"})
		return
	}
	submissionID := "sub_" + hex.EncodeToString(b)

	// TODO: Task 5 — send submission to DMOJ judge-server
	// Adapter must map: code→source, language→DMOJ executor ID (PY3, CPP17, etc.)
	// and supply: time-limit, memory-limit, short-circuit, meta from problem config
	resp := model.SubmitResponse{
		ID:        submissionID,
		Status:    "queued",
		ProblemID: req.ProblemID,
		Language:  req.Language,
		Message:   "submission received, pending judge execution",
	}

	c.JSON(http.StatusAccepted, resp)
}
