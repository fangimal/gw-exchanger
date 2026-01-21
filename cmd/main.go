package main

import (
	"fmt"
	"gw-exchanger/internal/config"
	"log"
)

func main() {

	//0. Подготовка логгера
	fmt.Println("Start...")

	//1. Загрузка конфига
	cfg := config.GetConfig()

	//2. Подключение к бд

	//3. Создание сервера

	//4. Запуск сервера на заданном порту

	log.Printf("Config: %+v", cfg)
}
