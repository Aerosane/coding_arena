package middleware

import (
	"testing"

	"github.com/gin-gonic/gin"
)

// TestAPIKeyAuth_ValidKey tests that valid API key allows request through.
func TestAPIKeyAuth_ValidKey(t *testing.T) {
	// Setup: Create auth middleware with valid keys
	validKeys := map[string]bool{
		"test-key-123": true,
		"another-key":  true,
	}
	
	router := setupTestRouter()
	router.Use(APIKeyAuth(validKeys))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request with valid API key
	w := makeRequest(router, "GET", "/test", map[string]string{
		"X-API-Key": "test-key-123",
	})

	// Assert: Should succeed
	assertStatus(t, w, 200)
}

// TestAPIKeyAuth_InvalidKey tests that invalid API key returns 403 Forbidden.
func TestAPIKeyAuth_InvalidKey(t *testing.T) {
	// Setup: Create auth middleware with valid keys
	validKeys := map[string]bool{
		"test-key-123": true,
	}
	
	router := setupTestRouter()
	router.Use(APIKeyAuth(validKeys))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request with invalid API key
	w := makeRequest(router, "GET", "/test", map[string]string{
		"X-API-Key": "wrong-key",
	})

	// Assert: Should return 403 Forbidden
	assertStatus(t, w, 403)
}

// TestAPIKeyAuth_MissingKey tests that missing API key returns 401 Unauthorized.
func TestAPIKeyAuth_MissingKey(t *testing.T) {
	// Setup: Create auth middleware
	validKeys := map[string]bool{
		"test-key-123": true,
	}
	
	router := setupTestRouter()
	router.Use(APIKeyAuth(validKeys))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request without API key header
	w := makeRequest(router, "GET", "/test", nil)

	// Assert: Should return 401 Unauthorized
	assertStatus(t, w, 401)
}

// TestAPIKeyAuth_WhitespacePadding tests that whitespace-padded keys are trimmed.
func TestAPIKeyAuth_WhitespacePadding(t *testing.T) {
	// Setup: Create auth middleware
	validKeys := map[string]bool{
		"test-key-123": true,
	}
	
	router := setupTestRouter()
	router.Use(APIKeyAuth(validKeys))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request with whitespace-padded key
	w := makeRequest(router, "GET", "/test", map[string]string{
		"X-API-Key": "  test-key-123  ",
	})

	// Assert: Should succeed (key is trimmed)
	assertStatus(t, w, 200)
}

// TestAPIKeyAuth_EmptyKey tests that empty string key is rejected.
func TestAPIKeyAuth_EmptyKey(t *testing.T) {
	// Setup: Create auth middleware
	validKeys := map[string]bool{
		"test-key-123": true,
	}
	
	router := setupTestRouter()
	router.Use(APIKeyAuth(validKeys))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request with empty key
	w := makeRequest(router, "GET", "/test", map[string]string{
		"X-API-Key": "",
	})

	// Assert: Should return 401 Unauthorized
	assertStatus(t, w, 401)
}

// TestAPIKeyAuth_CaseSensitive tests that API key validation is case-sensitive.
func TestAPIKeyAuth_CaseSensitive(t *testing.T) {
	// Setup: Create auth middleware
	validKeys := map[string]bool{
		"test-key-123": true,
	}
	
	router := setupTestRouter()
	router.Use(APIKeyAuth(validKeys))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Execute: Make request with wrong case
	w := makeRequest(router, "GET", "/test", map[string]string{
		"X-API-Key": "TEST-KEY-123", // Wrong case
	})

	// Assert: Should return 403 Forbidden (case-sensitive)
	assertStatus(t, w, 403)
}
