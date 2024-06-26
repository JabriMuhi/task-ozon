# Указываем базовый образ
FROM golang:1.22-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы
COPY . .

# Устанавливаем переменные окружения
ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG STORAGE_TYPE
ARG PORT

ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}
ENV STORAGE_TYPE=${STORAGE_TYPE}
ENV PORT=${PORT}

# Компилируем Go-приложение
RUN go build -o server ./server.go

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Копируем скрипт запуска
COPY ./docker/entrypoint.sh /app/entrypoint.sh

# Делаем скрипт исполняемым
RUN chmod +x /app/entrypoint.sh

# Указываем точку входа
ENTRYPOINT ["/app/entrypoint.sh"]
