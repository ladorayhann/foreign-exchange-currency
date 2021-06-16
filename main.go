package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

var symbols Symbols
var currencies Currencies
var filterCurrencies Currencies

func Server() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/currencies", getCurrencies).Methods("GET")
	r.HandleFunc("/api/symbols", getSymbols).Methods("GET")
	return r
}

func main() {
	currencies = currenciesSeeder(currencies)
	symbols = symbolsSeeder(symbols)
	r := Server()
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getCurrencies(writer http.ResponseWriter, request *http.Request) {
	parameter := GetQueryParam(request, "symbols")
	if len(parameter) != 0 {
		parameterSplit := strings.Split(parameter, ",")
		filterCurrencies.Base = "USD"
		for _, symbols := range parameterSplit {
			for _, item := range currencies.Rates {
				if symbols == item.Currency {
					filterCurrencies.Rates = append(filterCurrencies.Rates, item)
				}
			}
		}
		renderJSON(writer, filterCurrencies)
		filterCurrencies.Rates = []Rates{}
	} else {
		renderJSON(writer, currencies)
	}

}

func getSymbols(writer http.ResponseWriter, request *http.Request) {
	renderJSON(writer, symbols)
}
