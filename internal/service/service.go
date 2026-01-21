package service

import (
	"context"
	"gw-exchanger/internal/proto/proto/exchange"
	"gw-exchanger/internal/storages"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExchangeService struct {
	storage                                     storages.ExchangeStorage
	exchange.UnimplementedExchangeServiceServer // для совместимости
}

func NewExchangeService(storage storages.ExchangeStorage) *ExchangeService {
	return &ExchangeService{storage: storage}
}

// Реализуем методы из .proto
func (s *ExchangeService) GetExchangeRates(ctx context.Context, req *exchange.Empty) (*exchange.ExchangeRatesResponse, error) {
	rates, err := s.storage.GetRates()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get rates: %v", err)
	}
	return &exchange.ExchangeRatesResponse{Rates: rates}, nil
}

func (s *ExchangeService) GetExchangeRateForCurrency(ctx context.Context, req *exchange.CurrencyRequest) (*exchange.ExchangeRateResponse, error) {
	rate, err := s.storage.GetRate(req.FromCurrency, req.ToCurrency)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid currency pair: %v", err)
	}
	return &exchange.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         rate,
	}, nil
}
