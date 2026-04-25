package middleware

import (
	"testing"

	"github.com/gin-gonic/gin"
)

// TestCORS_AllowedOrigin tests that requests from allowed origins get CORS headers.
func TestCORS_AllowedOrigin(t *testing.T) {
	// Setup: Create CORS config with allowed origin
	config := CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		MaxAge:       3600,
	}
	
	router := setupTestRouter()
	router.Use(CORS(config))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request with allowed origin
	w := makeRequest(router, "GET", "/test", map[string]string{
		"Origin": "http://localhost:3000",
	})

	// Assert: CORS headers should be present
	assertHeader(t, w, "Access-Control-Allow-Origin", "http://localhost:3000")
	assertHeaderExists(t, w, "Access-Control-Allow-Methods")
	assertHeaderExists(t, w, "Access-Control-Allow-Headers")
	assertHeaderExists(t, w, "Vary")
}

// TestCORS_DisallowedOrigin tests that requests from disallowed origins don't get CORS headers.
func TestCORS_DisallowedOrigin(t *testing.T) {
	// Setup: Create CORS config with specific allowed origin
	config := CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		MaxAge:       3600,
	}
	
	router := setupTestRouter()
	router.Use(CORS(config))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request with disallowed origin
	w := makeRequest(router, "GET", "/test", map[string]string{
		"Origin": "http://evil.com",
	})

	// Assert: CORS headers should NOT be present
	assertHeaderNotExists(t, w, "Access-Control-Allow-Origin")
}

// TestCORS_PreflightRequest tests that OPTIONS preflight requests return 204 with CORS headers.
func TestCORS_PreflightRequest(t *testing.T) {
	// Setup: Create CORS config
	config := CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		MaxAge:       3600,
	}
	
	router := setupTestRouter()
	router.Use(CORS(config))
	router.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make OPTIONS preflight request
	w := makeRequest(router, "OPTIONS", "/test", map[string]string{
		"Origin": "http://localhost:3000",
	})

	// Assert: Should return 204 No Content
	assertStatus(t, w, 204)
	
	// Assert: CORS headers should be present
	assertHeader(t, w, "Access-Control-Allow-Origin", "http://localhost:3000")
}

// TestCORS_MissingOrigin tests that requests without Origin header are handled gracefully.
func TestCORS_MissingOrigin(t *testing.T) {
	// Setup: Create CORS config
	config := CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		MaxAge:       3600,
	}
	
	router := setupTestRouter()
	router.Use(CORS(config))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request without Origin header
	w := makeRequest(router, "GET", "/test", nil)

	// Assert: Request should still succeed
	assertStatus(t, w, 200)
	
	// Assert: No CORS headers should be added
	assertHeaderNotExists(t, w, "Access-Control-Allow-Origin")
}

// TestCORS_EmptyAllowedOrigins tests that empty allowed origins list blocks all origins.
func TestCORS_EmptyAllowedOrigins(t *testing.T) {
	// Setup: Create CORS config with empty allowed origins
	config := CORSConfig{
		AllowOrigins: []string{}, // Empty list
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		MaxAge:       3600,
	}
	
	router := setupTestRouter()
	router.Use(CORS(config))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request with any origin
	w := makeRequest(router, "GET", "/test", map[string]string{
		"Origin": "http://localhost:3000",
	})

	// Assert: No CORS headers should be added
	assertHeaderNotExists(t, w, "Access-Control-Allow-Origin")
}

// TestCORS_VaryHeader tests that Vary header is set for caching purposes.
func TestCORS_VaryHeader(t *testing.T) {
	// Setup: Create CORS config
	config := CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		MaxAge:       3600,
	}
	
	router := setupTestRouter()
	router.Use(CORS(config))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request with allowed origin
	w := makeRequest(router, "GET", "/test", map[string]string{
		"Origin": "http://localhost:3000",
	})

	// Assert: Vary header should be set to Origin
	assertHeader(t, w, "Vary", "Origin")
}
