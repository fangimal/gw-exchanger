package service

import (
	"context"
	"gw-exchanger/internal/proto/proto/exchange"
	"gw-exchanger/internal/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExchangeService struct {
	exchange.UnimplementedExchangeServiceServer
	storage repository.Repository
}

func NewExchangeService(storage repository.Repository) *ExchangeService {
	return &ExchangeService{storage: storage}
}

// Реализуем методы из .proto
func (s *ExchangeService) GetExchangeRates(ctx context.Context, req *exchange.Empty) (*exchange.ExchangeRatesResponse, error) {
	exchangeRates, err := s.storage.GetRates()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get rates: %v", err)
	}

	rates := make(map[string]float32)
	for _, rate := range exchangeRates {
		key := rate.FromCurrency + "_" + rate.ToCurrency
		rates[key] = rate.Rate
	}

	return &exchange.ExchangeRatesResponse{Rates: rates}, nil
}

func (s *ExchangeService) GetExchangeRateForCurrency(ctx context.Context, req *exchange.CurrencyRequest) (*exchange.ExchangeRateResponse, error) {
	exchangeRate, err := s.storage.GetRate(req.FromCurrency, req.ToCurrency)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid currency pair: %v", err)
	}
	return &exchange.ExchangeRateResponse{
		FromCurrency: exchangeRate.FromCurrency,
		ToCurrency:   exchangeRate.ToCurrency,
		Rate:         exchangeRate.Rate,
	}, nil
}
