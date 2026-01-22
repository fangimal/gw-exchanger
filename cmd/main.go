package main

import (
	"context"
	"gw-exchanger/internal/config"
	"gw-exchanger/internal/proto/proto/exchange"
	"gw-exchanger/internal/repository"
	"gw-exchanger/internal/service"
	"gw-exchanger/pkg/logging"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	//0. Подготовка логгера
	logger := logging.GetLogger()
	logger.Info("Start...")

	//1. Загрузка конфига
	cfg := config.GetConfig()
	logger.Info("Config loaded")

	//2. Подключение к бд
	repo, closeDB := repository.NewRepository(ctx, &cfg.Storage, logger)
	defer closeDB()

	//3. Создание сервера
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	exchangeService := service.NewExchangeService(repo)
	exchange.RegisterExchangeServiceServer(grpcServer, exchangeService)

	//4. Запуск сервера на заданном порту
	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		logger.Fatalf("gRPC server failed to listen: %v", err)
	}

	logger.Infof("gRPC server starting on port %s", cfg.GRPC.Port)
	if err = grpcServer.Serve(lis); err != nil {
		logger.Fatalf("gRPC server failed: %v", err)
	}

	//5. Ожидание сигнала завершения
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	logger.Info("Shutting down gRPC server...")
	grpcServer.GracefulStop()
}
