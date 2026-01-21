package main

import (
	"fmt"
	"gw-exchanger/internal/config"
	"gw-exchanger/internal/proto/proto/exchange"
	"gw-exchanger/internal/service"
	"gw-exchanger/internal/storages"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	//0. Подготовка логгера
	fmt.Println("Start...")

	//1. Загрузка конфига
	cfg := config.GetConfig()

	//2. Подключение к бд
	storage := storages.NewRepository("mock")

	//3. Создание сервера
	grpcServer := grpc.NewServer()
	exchangeService := service.NewExchangeService(storage)
	exchange.RegisterExchangeServiceServer(grpcServer, exchangeService)

	//4. Запуск сервера на заданном порту
	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
	log.Printf("gRPC server starting on port %s", cfg.GRPC.Port)
}
