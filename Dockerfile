# Этап 1: сборка бинарника
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Установите ca-certificates для TLS
RUN apk add --no-cache ca-certificates

# Явно установите GOPROXY
ENV GOPROXY=https://goproxy.io,direct

# Копируем go.mod и go.sum (для кэширования зависимостей)
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник (статически, без CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gw-exchanger ./cmd

# Этап 2: финальный образ
FROM alpine:3.23

# Устанавливаем ca-certificates для TLS (на всякий случай)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник и конфиг
COPY --from=builder /app/gw-exchanger .
COPY config.docker.yml config.yml

# Порт gRPC (опционально, для документации)
EXPOSE 50052

# Запуск
CMD ["./gw-exchanger", "-c", "config.yml"]