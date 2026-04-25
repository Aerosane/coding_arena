package middleware

import (
	"testing"

	"github.com/gin-gonic/gin"
)

// TestMiddlewareIntegration tests multiple middleware components working together.
// This ensures that the middleware stack behaves correctly when combined,
// matching the production configuration in main.go.
func TestMiddlewareIntegration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("RateLimiter_And_Auth", func(t *testing.T) {
		// Setup: Create router with rate limiter (burst 1) and auth
		limiter := NewRateLimiter(1.0, 1)
		validKeys := map[string]bool{"test-key": true}
		
		router := setupTestRouter()
		router.Use(limiter.Middleware())
		router.Use(APIKeyAuth(validKeys))
		router.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": true})
		})

		// Execute: First request with valid key should succeed
		w1 := makeRequest(router, "GET", "/test", map[string]string{
			"X-API-Key": "test-key",
		})
		assertStatus(t, w1, 200)

		// Execute: Second request should be rate limited BEFORE reaching auth
		w2 := makeRequest(router, "GET", "/test", map[string]string{
			"X-API-Key": "test-key",
		})
		assertStatus(t, w2, 429) // Rate limited, never reached auth
	})

	t.Run("CORS_And_SecurityHeaders", func(t *testing.T) {
		// Setup: Create router with CORS and security headers
		corsConfig := CORSConfig{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"GET", "POST"},
			AllowHeaders: []string{"Content-Type"},
			MaxAge:       3600,
		}
		
		router := setupTestRouter()
		router.Use(CORS(corsConfig))
		router.Use(SecurityHeaders())
		router.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": true})
		})

		// Execute: Make request with allowed origin
		w := makeRequest(router, "GET", "/test", map[string]string{
			"Origin": "http://localhost:3000",
		})

		// Assert: Both CORS and security headers should be present
		assertHeader(t, w, "Access-Control-Allow-Origin", "http://localhost:3000")
		assertHeaderExists(t, w, "X-Frame-Options")
		assertHeaderExists(t, w, "X-Content-Type-Options")
		assertHeaderExists(t, w, "Strict-Transport-Security")
	})

	t.Run("FullMiddlewareStack", func(t *testing.T) {
		// Setup: Create router with full middleware stack like main.go
		// Order: Recovery -> Logger -> SecurityHeaders -> MaxBodySize -> CORS -> RateLimiter -> Auth
		limiter := NewRateLimiter(10.0, 20)
		validKeys := map[string]bool{"prod-key": true}
		corsConfig := CORSConfig{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"GET", "POST"},
			AllowHeaders: []string{"Content-Type", "X-API-Key"},
			MaxAge:       3600,
		}
		
		router := setupTestRouter()
		router.Use(gin.Recovery())
		router.Use(RequestLogger())
		router.Use(SecurityHeaders())
		router.Use(MaxBodySize(1024))
		router.Use(CORS(corsConfig))
		router.Use(limiter.Middleware())
		router.Use(APIKeyAuth(validKeys))
		router.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": true})
		})

		// Execute: Make request with all required headers
		w := makeRequest(router, "GET", "/test", map[string]string{
			"Origin":    "http://localhost:3000",
			"X-API-Key": "prod-key",
		})

		// Assert: Request should succeed with all headers
		assertStatus(t, w, 200)
		assertHeader(t, w, "Access-Control-Allow-Origin", "http://localhost:3000")
		assertHeaderExists(t, w, "X-Frame-Options")
		assertHeaderExists(t, w, "Cache-Control")
	})

	t.Run("ErrorPropagation", func(t *testing.T) {
		// Setup: Create router with multiple middleware
		// Auth will abort the request, subsequent middleware shouldn't execute
		validKeys := map[string]bool{"valid-key": true}
		
		router := setupTestRouter()
		router.Use(SecurityHeaders())
		router.Use(APIKeyAuth(validKeys))
		router.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": true})
		})

		// Execute: Make request without API key (auth will abort)
		w := makeRequest(router, "GET", "/test", nil)

		// Assert: Should return 401 from auth middleware
		assertStatus(t, w, 401)
		
		// Assert: Security headers should still be present (ran before auth)
		assertHeaderExists(t, w, "X-Frame-Options")
		assertHeaderExists(t, w, "X-Content-Type-Options")
	})
}

// TestMiddlewareOrder tests that middleware execution order matters.
func TestMiddlewareOrder(t *testing.T) {
	// Test 1: Rate limiter before auth
	t.Run("RateLimiter_Before_Auth", func(t *testing.T) {
		limiter := NewRateLimiter(1.0, 1)
		validKeys := map[string]bool{"test-key": true}
		
		router := setupTestRouter()
		router.Use(limiter.Middleware()) // Rate limiter first
		router.Use(APIKeyAuth(validKeys)) // Auth second
		router.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": true})
		})

		// Exhaust rate limit
		makeRequest(router, "GET", "/test", map[string]string{"X-API-Key": "test-key"})
		
		// Next request should be rate limited even with valid key
		w := makeRequest(router, "GET", "/test", map[string]string{"X-API-Key": "test-key"})
		assertStatus(t, w, 429) // Rate limited, never reached auth
	})

	// Test 2: Auth before rate limiter
	t.Run("Auth_Before_RateLimiter", func(t *testing.T) {
		limiter := NewRateLimiter(1.0, 1)
		validKeys := map[string]bool{"test-key": true}
		
		router := setupTestRouter()
		router.Use(APIKeyAuth(validKeys)) // Auth first
		router.Use(limiter.Middleware()) // Rate limiter second
		router.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": true})
		})

		// Request without key should fail at auth (before rate limiter)
		w := makeRequest(router, "GET", "/test", nil)
		assertStatus(t, w, 401) // Auth failed, never reached rate limiter
	})
}

// BenchmarkMiddlewareStack benchmarks the performance of the full middleware stack.
func BenchmarkMiddlewareStack(b *testing.B) {
	// Setup: Create router with full middleware stack
	limiter := NewRateLimiter(1000.0, 2000) // High limits for benchmarking
	validKeys := map[string]bool{"bench-key": true}
	corsConfig := CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type", "X-API-Key"},
		MaxAge:       3600,
	}
	
	router := setupTestRouter()
	router.Use(gin.Recovery())
	router.Use(RequestLogger())
	router.Use(SecurityHeaders())
	router.Use(MaxBodySize(1024))
	router.Use(CORS(corsConfig))
	router.Use(limiter.Middleware())
	router.Use(APIKeyAuth(validKeys))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		makeRequest(router, "GET", "/test", map[string]string{
			"Origin":    "http://localhost:3000",
			"X-API-Key": "bench-key",
		})
	}
}
