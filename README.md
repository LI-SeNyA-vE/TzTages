# TzTages

**TzTages** — это gRPC-сервис для загрузки, скачивания и просмотра изображений с поддержкой хранения в SQLite и ограничением количества запросов по IP.

## Структура

- `api/proto/v1/images/` — .proto файл для генерации gRPC API
- `cmd/agent/` — CLI-клиент для взаимодействия с сервером
- `cmd/server/` — gRPC сервер
- `internal/agent/action/` — логика клиента
- `internal/delivery/grpc/` — хэндлеры gRPC + middleware
- `internal/repository/database/liteSQL/` — SQLite реализация storage
- `internal/domain/` — доменные структуры
- `pgk/` — утилиты (проверка порта)

---

## Функциональности

- Загрузка изображений (UploadImage)
- Скачивание изображения по ID (DownloadImage)
- Список всех изображений (ListImages)
- Middleware-ограничение на число запросов от IP

---

## Старт gRPC-сервера

```bash
cd cmd/server

go run main.go
```

- Введите порт, на котором запустится gRPC-сервер

---

## Старт CLI-клиента

```bash
cd cmd/agent

go run main.go
```

- Введите порт сервера
- Выберите одну из доступных команд:
    - `1` — Загрузка файла
    - `2` — Скачивание по ID
    - `3` — Список изображений
    - `0` — Выход

---

## Команда для генерации .pb.go

```bash
protoc \
  --proto_path=. \
  --go_out=. \
  --go-grpc_out=. \
  api/proto/v1/images/images.proto
```

---

## Зависимости

- Go 1.20+
- gRPC + Protocol Buffers
- SQLite (`github.com/mattn/go-sqlite3`)
- `github.com/sqweek/dialog` для GUI-выбора файла
- `logrus` для логгирования

---

## TODO

- [ ] Dockerfile + docker-compose
- [ ] JWT-авторизация
- [ ] Тесты (юнит/интеграционные)
- [ ] Web UI (React/Vue)

---

## Автор

**LI-SeNyA-vE**

