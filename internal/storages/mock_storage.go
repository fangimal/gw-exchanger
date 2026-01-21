package storages

import (
	"fmt"
)

type MockStorage struct{}

func NewMockStorage() ExchangeStorage {
	return &MockStorage{}
}

func (m *MockStorage) GetRates() (map[string]float32, error) {
	return map[string]float32{
		"USD": 1.0,
		"RUB": 90.0,
		"EUR": 0.93,
	}, nil
}

func (m *MockStorage) GetRate(from, to string) (float32, error) {
	rates := map[string]float32{
		"USD_RUB": 90.0,
		"USD_EUR": 0.93,
		"EUR_USD": 1.075,
		"RUB_USD": 0.0111,
	}
	key := from + "_" + to
	if rate, ok := rates[key]; ok {
		return rate, nil
	}
	return 0, fmt.Errorf("unsupported currency pair: %s -> %s", from, to)
}
