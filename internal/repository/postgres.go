package repository

import (
	"context"
	"fmt"

	"gw-exchanger/internal/config"
	"gw-exchanger/internal/db"
	"gw-exchanger/pkg/logging"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
	log  *logging.Logger
}

func NewPostgresRepository(ctx context.Context, cfg *config.StorageConfig, log *logging.Logger) (r *PostgresRepository, closeConnection func()) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("unable to create pgxpool: %v", err)
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}

	log.Info("connected to PostgreSQL")

	return &PostgresRepository{
		pool: pool,
		log:  log,
	}, pool.Close
}

func (r *PostgresRepository) GetRates() ([]db.ExchangeRate, error) {
	rows, err := r.pool.Query(context.Background(), "SELECT from_currency, to_currency, rate FROM exchange_rates")
	if err != nil {
		r.log.Errorf("failed to query rates: %v", err)
		return nil, fmt.Errorf("failed to query rates: %w", err)
	}
	defer rows.Close()

	var rates []db.ExchangeRate
	for rows.Next() {
		var rate db.ExchangeRate
		if err := rows.Scan(&rate.FromCurrency, &rate.ToCurrency, &rate.Rate); err != nil {
			r.log.Errorf("failed to scan row: %v", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		rates = append(rates, rate)
	}

	if err = rows.Err(); err != nil {
		r.log.Errorf("rows iteration error: %v", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return rates, nil
}

func (r *PostgresRepository) GetRate(fromCurrency, toCurrency string) (db.ExchangeRate, error) {
	var rate db.ExchangeRate
	sql := "SELECT from_currency, to_currency, rate FROM exchange_rates WHERE from_currency = $1 AND to_currency = $2"
	err := r.pool.QueryRow(context.Background(),
		sql,
		fromCurrency, toCurrency).Scan(&rate.FromCurrency, &rate.ToCurrency, &rate.Rate)

	if err != nil {
		r.log.Errorf("failed to get rate for %s -> %s: %v", fromCurrency, toCurrency, err)
		return db.ExchangeRate{}, fmt.Errorf("rate not found for %s -> %s", fromCurrency, toCurrency)
	}

	return rate, nil
}
