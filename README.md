# Auth gRPC Service

Локальный gRPC-сервер аутентификации на Go. На текущем этапе transport,
конфигурация, middleware и подключения к Postgres/Redis готовы, а бизнес-методы
`AuthService` намеренно ещё не реализованы.

## Требования

- Go 1.26+
- Docker Engine с Docker Compose v2
- `protoc`, только если требуется повторно сгенерировать protobuf-код

## Быстрый запуск

1. Проверьте значения в `.env`. Для локального запуска оставьте `POSTGRES_HOST`
   и `REDIS_HOST` закомментированными: тогда используются `localhost`.
2. Запустите инфраструктуру и сервер:

   ```bash
   make dev
   ```

   Команда дождётся healthcheck-ов Postgres и Redis, затем запустит сервер на
   адресе из `GRPC_PORT` (по умолчанию `:44044`).

3. Остановите сервер сочетанием `Ctrl+C`. Сервер перестанет принимать новые RPC,
   дождётся активных до `GRPC_SHUTDOWN_TIMEOUT`, а затем при необходимости
   принудительно остановится.

4. Остановите контейнеры, когда они больше не нужны:

   ```bash
   make docker-down
   ```

## Полезные команды

```bash
make help          # список основных команд
make docker-up     # только Postgres и Redis
make run           # только gRPC-сервер
make infra-logs    # логи Postgres и Redis
make test          # тесты
make check         # тесты и go vet
make fmt           # форматирование исходного Go-кода
make gen           # генерация Go-кода из proto
```

## Middleware

Для каждого unary RPC сервер:

- принимает `x-request-id` от клиента или генерирует новый и возвращает его в
  response metadata;
- ограничивает время выполнения значением `GRPC_TIMEOUT`;
- записывает в лог request ID, RPC-метод, gRPC status code, длительность и
  ошибку, если она есть.

Вызовы RPC пока будут возвращать `Unimplemented`: это ожидаемое поведение до
появления service-логики.
