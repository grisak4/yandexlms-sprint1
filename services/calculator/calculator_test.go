package calculator_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sprintone/services/calculator"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCalculatorGet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/calculator", calculator.CalculatorGet)

	tests := []struct {
		name         string
		input        map[string]string
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name: "валидное выражение",
			input: map[string]string{
				"expression": "2+2",
			},
			expectedCode: 200,
			expectedBody: map[string]interface{}{
				"result": 4.0,
			},
		},
		{
			name: "деление на ноль",
			input: map[string]string{
				"expression": "1/0",
			},
			expectedCode: 500,
			expectedBody: map[string]interface{}{
				"error": "Internal server error",
			},
		},
		{
			name: "некорректное выражение",
			input: map[string]string{
				"expression": "1++2",
			},
			expectedCode: 500,
			expectedBody: map[string]interface{}{
				"error": "Internal server error",
			},
		},
		{
			name:         "отсутствует выражение",
			input:        map[string]string{},
			expectedCode: 422,
			expectedBody: map[string]interface{}{
				"error": "Expression is not valid",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)

			req, _ := http.NewRequest(http.MethodPost, "/calculator", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			var responseBody map[string]interface{}
			_ = json.Unmarshal(w.Body.Bytes(), &responseBody)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, responseBody)
		})
	}
}
