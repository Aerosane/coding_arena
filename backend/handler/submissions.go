package handler

import (
	"net/http"

	"github.com/GCET-Open-Source-Foundation/coding_arena/backend/db"
	"github.com/gin-gonic/gin"
)

// ListSubmissions handles GET /api/submissions?problem=...&verdict=...
func ListSubmissions(c *gin.Context) {
	problem := c.Query("problem")
	verdict := c.Query("verdict")

	submissions, err := db.GetSubmissions(c.Request.Context(), problem, verdict)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve submissions",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"submissions": submissions,
	})
}