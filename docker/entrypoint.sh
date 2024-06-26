#!/bin/sh

# Запуск тестов и вывод результатов
echo "Running tests..."
go test ./... -cover

until nc -z -v -w30 $DB_HOST $DB_PORT
do
  echo "Ожидание запуска PostgreSQL на хосте $DB_HOST и порту $DB_PORT..."
  sleep 5
done

echo "PostgreSQL запущен и готов к работе!"

# Making DB migrations
echo "Making DB migrations..."
goose -dir migrations postgres "host=$DB_HOST port=$DB_PORT user=$DB_USER dbname=$DB_NAME password=$DB_PASSWORD sslmode=disable" up


# Запускаем сервер с нужным типом хранилища
./server --storageType="${STORAGE_TYPE}"

