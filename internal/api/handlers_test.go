package api_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"totoconfigapi/internal/api"
	"totoconfigapi/internal/db"
	test "totoconfigapi/internal/test/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetConfigHandler(t *testing.T) {
	tests := []struct {
		name               string
		packageID          string
		country            string
		setupMock          func(*test.MockConfigDB)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:      "Success",
			packageID: "testPackage",
			country:   "US",
			setupMock: func(m *test.MockConfigDB) {
				m.On("GetConfigs", "testPackage", "ZZ", mock.AnythingOfType("int")).Return([]*db.Config{{MainSKU: "SKU123"}}, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[{"main_sku":"SKU123"}]`,
		},
		{
			name:      "Database Error",
			packageID: "testPackage",
			country:   "US",
			setupMock: func(m *test.MockConfigDB) {
				m.On("GetConfigs", "testPackage", "ZZ", mock.AnythingOfType("int")).Return([]*db.Config{}, errors.New("database error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Failed to retrieve configuration"}`,
		},
		{
			name:      "No configs found",
			packageID: "testPackage",
			country:   "",
			setupMock: func(m *test.MockConfigDB) {

				m.On("GetConfigs", "testPackage", "ZZ", mock.Anything).Return([]*db.Config{}, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `null`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := new(test.MockConfigDB)
			if tc.setupMock != nil {
				tc.setupMock(mockDB)
			}

			r := gin.New()
			apiHandler := api.NewHandler(mockDB)
			r.GET("/api/configs", apiHandler.GetConfigHandler)

			req, _ := http.NewRequest("GET", "/api/configs?package="+tc.packageID, nil)

			// Manually set country if provided (simulating middleware behavior)
			if tc.country != "" {
				req.Header.Set("X-Appengine-Country", tc.country)
			}

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.Contains(t, w.Body.String(), tc.expectedResponse)
		})
	}
}
