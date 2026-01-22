package db

type ExchangeRate struct {
	FromCurrency string  `json:"from_currency"`
	ToCurrency   string  `json:"to_currency"`
	Rate         float32 `json:"rate"`
}
