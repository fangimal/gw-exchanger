package repository

import (
	"context"
	"gw-exchanger/internal/config"
	"gw-exchanger/internal/db"
	"gw-exchanger/pkg/logging"
)

type Repository interface {
	GetRates() ([]db.ExchangeRate, error)
	GetRate(fromCurrency, toCurrency string) (db.ExchangeRate, error)
}

func NewRepository(ctx context.Context, cfg *config.StorageConfig, log *logging.Logger) (r Repository, closeConnection func()) {
	return NewPostgresRepository(ctx, cfg, log)
}
