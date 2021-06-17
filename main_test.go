package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCurrencies(t *testing.T) {
	currencies = currenciesSeeder(currencies)
	request, _ := http.NewRequest("GET", "/api/currencies", nil)
	response := httptest.NewRecorder()

	Server().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)
	t.Log(response.Body.String())
}

func TestCurrenciesFilter(t *testing.T) {
	currencies = currenciesSeeder(currencies)
	request, _ := http.NewRequest("GET", "/api/currencies?symbols=USD", nil)
	response := httptest.NewRecorder()

	Server().ServeHTTP(response, request)
	assert.Equal(t, "{\"base\":\"USD\",\"rates\":[{\"currency\":\"USD\",\"name\":\"United States Dollar\",\"rate\":1}]}\n", response.Body.String())
}

func TestSymbol(t *testing.T) {
	symbols = symbolsSeeder(symbols)
	request, _ := http.NewRequest("GET", "/api/symbols", nil)
	response := httptest.NewRecorder()
	Server().ServeHTTP(response, request)
	assert.Equal(t, "{\"symbols\":[{\"id\":\"USD\",\"name\":\"United States Dollar\"},{\"id\":\"EUR\",\"name\":\"Euro\"},{\"id\":\"CAD\",\"name\":\"Canadian Dollar\"},{\"id\":\"IDR\",\"name\":\"Indonesian Rupiah\"},{\"id\":\"GBP\",\"name\":\"British Pound Sterling\"},{\"id\":\"CHF\",\"name\":\"Swiss Franc\"},{\"id\":\"SGD\",\"name\":\"Singapore Dollar\"},{\"id\":\"INR\",\"name\":\"Indian Rupee\"},{\"id\":\"MYR\",\"name\":\"Malaysian Ringgit\"},{\"id\":\"JPY\",\"name\":\"Japanese Yen\"},{\"id\":\"KRW\",\"name\":\"South Korean Won\"}]}\n", response.Body.String())
	t.Log(response.Body.String())
}
