FROM golang:1.22 AS builder

WORKDIR /currency

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum  ./
RUN go mod download
COPY cmd/ ./cmd
COPY internal/ ./internal
COPY pkg/ ./pkg
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /build/currency ./cmd

# Финальный этап с минимальным образом
FROM alpine:3 AS currency

# Копируем скомпилированное приложение
COPY --from=builder /build/currency /bin/currency
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Europe/Minsk /etc/localtime && \
    echo "Europe/Minsk" > /etc/timezone
ENV NODE_ENV=DOCKER
# Указываем точку входа
ENTRYPOINT ["/bin/currency"]
