package main

type Symbol struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Symbols struct {
	Symbols []Symbol `json:"symbols"`
}

type Currencies struct {
	Base  string  `json:"base"`
	Rates []Rates `json:"rates"`
}

type Rates struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}
