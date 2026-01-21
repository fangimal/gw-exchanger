package main

import (
	"gw-exchanger/internal/config"
	"gw-exchanger/internal/proto/proto/exchange"
	"gw-exchanger/internal/service"
	"gw-exchanger/internal/storages"
	"gw-exchanger/pkg/logging"
	"net"

	"google.golang.org/grpc"
)

func main() {

	//0. Подготовка логгера
	logger := logging.GetLogger()
	logger.Info("Start...")

	//1. Загрузка конфига
	cfg := config.GetConfig()
	logger.Info("Load config ...", cfg)

	//2. Подключение к бд
	storage := storages.NewRepository("mock")

	//3. Создание сервера
	grpcServer := grpc.NewServer()
	exchangeService := service.NewExchangeService(storage)
	exchange.RegisterExchangeServiceServer(grpcServer, exchangeService)

	//4. Запуск сервера на заданном порту
	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		logger.Fatalf("gRPC server failed: %v", err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		logger.Fatalf("gRPC server failed: %v", err)
	}
	logger.Printf("gRPC server starting on port %s", cfg.GRPC.Port)
}
