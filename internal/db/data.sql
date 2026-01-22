CREATE TABLE IF NOT EXISTS exchange_rates (
                                              from_currency VARCHAR(3) NOT NULL,
                                              to_currency   VARCHAR(3) NOT NULL,
                                              rate          NUMERIC(10,4) NOT NULL,
                                              PRIMARY KEY (from_currency, to_currency)
);

INSERT INTO exchange_rates (from_currency, to_currency, rate) VALUES
                                                                  ('USD', 'RUB', 90.0),
                                                                  ('USD', 'EUR', 0.93),
                                                                  ('EUR', 'USD', 1.075),
                                                                  ('RUB', 'USD', 0.0111);