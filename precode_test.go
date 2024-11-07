package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Lenf(t, list, totalCount, "expected cafe count %d, got %d", totalCount, len(list))
}

func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка кода ответа
	expected := http.StatusOK
	result := responseRecorder.Code
	require.Equal(t, expected, result, "status code must be 200")

	// Проверка тела ответа
	body := responseRecorder.Body.String()
	require.NotEmpty(t, body, "body must be not empty")
}

func TestMainHandlerWhenNoCorrectCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=lipetsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка города
	expected := http.StatusBadRequest
	result := responseRecorder.Code
	require.Equal(t, expected, result, "status code must be 400")

	// Проверка тела ответа
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body)

}
