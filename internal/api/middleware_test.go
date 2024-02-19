package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"totoconfigapi/internal/api"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCountryLookupMiddleware(t *testing.T) {
	tests := []struct {
		name            string
		setupRequest    func() *http.Request
		expectedCountry string
	}{
		{
			name: "Header Set",
			setupRequest: func() *http.Request {
				req, _ := http.NewRequest("GET", "/test", nil)
				req.Header.Set("X-Appengine-Country", "US")
				return req
			},
			expectedCountry: "US",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.Default()
			router.Use(api.CountryLookupMiddleware())

			router.GET("/test", func(c *gin.Context) {
				country, exists := c.Get("country")
				assert.True(t, exists)
				assert.Equal(t, tc.expectedCountry, country)
			})

			w := httptest.NewRecorder()
			req := tc.setupRequest()
			router.ServeHTTP(w, req)
		})
	}
}
