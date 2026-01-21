package storages

type ExchangeStorage interface {
	GetRates() (map[string]float32, error)
	GetRate(fromCurrency, toCurrency string) (float32, error)
}
