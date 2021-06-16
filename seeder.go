package main

func symbolsSeeder(symbols Symbols) Symbols {
	if len(symbols.Symbols) == 0 {
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "USD",
			Name: "United States Dollar",
		})
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "CAD",
			Name: "Canadian Dollar",
		})
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "IDR",
			Name: "Indonesian Rupiah",
		})
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "GBP",
			Name: "British Pound Sterling",
		})
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "CHF",
			Name: "Swiss Franc",
		})
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "SGD",
			Name: "Singapore Dollar",
		})
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "INR",
			Name: "Indian Rupee",
		})
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "MYR",
			Name: "Malaysian Ringgit",
		})
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "JPY",
			Name: "Japanese Yen",
		})
		symbols.Symbols = append(symbols.Symbols, Symbol{
			ID:   "KRW",
			Name: "South Korean Won",
		})
	}

	return symbols
}

func currenciesSeeder(currencies Currencies) Currencies {

	if len(currencies.Rates) == 0 {
		currencies.Base = "USD"
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "USD",
			Rate:     1,
		})
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "CAD",
			Rate:     1.22,
		})
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "IDR",
			Rate:     14259.35,
		})
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "GBP",
			Rate:     0.71,
		})
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "CHF",
			Rate:     0.90,
		})
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "SGD",
			Rate:     1.33,
		})
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "INR",
			Rate:     73.37,
		})
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "MYR",
			Rate:     4.12,
		})
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "JPY",
			Rate:     110.12,
		})
		currencies.Rates = append(currencies.Rates, Rates{
			Currency: "KRW",
			Rate:     1118.44,
		})
	}
	return currencies
}
