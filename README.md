# Pollify

**Pollify** — backend-приложение, написанное на Go. Проект использует PostgreSQL для хранения данных, RabbitMQ для работы с очередями сообщений, Redis для кэширования, а также Prometheus и Grafana для мониторинга приложения и инфраструктуры.

## Технологии

Проект использует следующие технологии:

* **Go** — основной язык разработки
* **PostgreSQL** — реляционная база данных
* **RabbitMQ** — брокер сообщений
* **Redis** — in-memory хранилище и кэш
* **Docker Compose** — управление локальной инфраструктурой
* **Prometheus** — сбор и хранение метрик
* **Grafana** — визуализация метрик и мониторинг
* **golang-migrate** — управление миграциями базы данных

---

# Требования

Перед началом работы необходимо установить:

* Go
* Docker
* Docker Compose
* Make

Проверьте установленные инструменты:

```bash
go version
docker --version
docker compose version
make --version
```

---

# Структура проекта

Примерная структура проекта:

```text
pollify/
├── cmd/
│   └── pollify/
│       └── main.go
├── migrations/
├── out/
│   ├── logs/
│   └── pgdata/
├── prometheus.yml
├── docker-compose.yml
├── Makefile
├── go.mod
├── go.sum
└── .env
```

### Основные директории

| Директория           | Назначение                             |
| -------------------- | -------------------------------------- |
| `cmd/pollify`        | Точка входа в приложение               |
| `migrations`         | SQL-миграции базы данных               |
| `out/logs`           | Логи приложения                        |
| `out/pgdata`         | Данные PostgreSQL                      |
| `prometheus.yml`     | Конфигурация Prometheus                |
| `docker-compose.yml` | Конфигурация инфраструктурных сервисов |

---

# Конфигурация

Проект использует файл `.env` для хранения переменных окружения.

Пример:

```env
POSTGRES_USER=pollify
POSTGRES_PASSWORD=pollify_password
POSTGRES_DB=pollify

RABBIT_USER=pollify
RABBIT_PASSWORD=pollify_password
```

Переменные из `.env` автоматически используются в `Makefile`:

```makefile
include .env
export
```

---

# Запуск инфраструктуры

Для запуска всех необходимых сервисов выполните:

```bash
make env-up
```

Будут запущены:

* PostgreSQL
* RabbitMQ
* Redis
* Grafana
* Prometheus

Проверить запущенные контейнеры:

```bash
docker compose ps
```

---

# Остановка инфраструктуры

Для остановки сервисов:

```bash
make env-down
```

---

# Очистка данных окружения

PostgreSQL сохраняет данные локально в:

```text
out/pgdata
```

Для полного удаления данных выполните:

```bash
make env-cleanup
```

Команда запросит подтверждение:

```text
Clean all environment volume files? [y/N]:
```

Введите:

```text
y
```

для удаления данных PostgreSQL.

> Внимание: эта операция удаляет локальные данные базы данных.

---

# PostgreSQL

PostgreSQL запускается в Docker.

Для удобного подключения к базе данных можно использовать port forwarder.

Запуск:

```bash
make env-port-forward
```

Остановка:

```bash
make env-port-close
```

После запуска port forwarder PostgreSQL будет доступен локально:

```text
Host: localhost
Port: 5432
Database: значение POSTGRES_DB
User: значение POSTGRES_USER
Password: значение POSTGRES_PASSWORD
```

---

# RabbitMQ

RabbitMQ запускается вместе с Management UI.

После запуска инфраструктуры интерфейс управления доступен по адресу:

```text
http://localhost:15672
```

Порт AMQP:

```text
localhost:5672
```

Учётные данные задаются через переменные:

```env
RABBIT_USER=
RABBIT_PASSWORD=
```

---

# Redis

Redis доступен локально по адресу:

```text
localhost:6379
```

Для подключения можно использовать:

```bash
redis-cli
```

Или:

```bash
redis-cli -h localhost -p 6379
```

---

# Миграции базы данных

Для управления миграциями используется `golang-migrate`.

## Создание миграции

Для создания новой миграции:

```bash
make migrate-create seq=create_users
```

Будут созданы SQL-файлы в директории:

```text
migrations/
```

Например:

```text
000001_create_users.up.sql
000001_create_users.down.sql
```

---

## Применение миграций

Для применения всех доступных миграций:

```bash
make migrate-up
```

---

## Откат миграции

Для отката последней миграции:

```bash
make migrate-down
```

Команда подключается к PostgreSQL внутри Docker-сети:

```text
postgres://POSTGRES_USER:POSTGRES_PASSWORD@pollify-postgres:5432/POSTGRES_DB
```

---

# Запуск Pollify

Перед запуском приложения рекомендуется запустить инфраструктуру:

```bash
make env-up
```

Затем, если необходимо подключение к PostgreSQL через localhost:

```bash
make env-port-forward
```

После этого запустите приложение:

```bash
make pollify-run
```

Команда выполняет следующие действия:

```bash
go mod tidy
go run ./cmd/pollify/main.go
```

Также автоматически устанавливаются переменные:

```text
LOGGER_FOLDER=${PROJECT_ROOT}/out/logs
POSTGRES_HOST=localhost
```

---

# Мониторинг

Проект использует Prometheus и Grafana для мониторинга.

## Prometheus

Prometheus доступен по адресу:

```text
http://localhost:9090
```

Конфигурация находится в файле:

```text
prometheus.yml
```

Пример конфигурации Pollify:

```yaml
global:
  scrape_interval: 5s

scrape_configs:
  - job_name: "pollify"
    scrape_interval: 5s
    static_configs:
      - targets:
          - "host.docker.internal:9091"
```

Prometheus собирает метрики приложения Pollify с endpoint:

```text
http://localhost:9091/metrics
```

---

# Проверка Prometheus

После запуска приложения откройте Prometheus:

```text
http://localhost:9090
```

В интерфейсе Prometheus можно проверить доступность Pollify с помощью запроса:

```promql
up{job="pollify"}
```

Если результат равен:

```text
1
```

Prometheus успешно получает метрики.

Если результат равен:

```text
0
```

Prometheus не может подключиться к приложению.

Также можно проверить конкретные метрики:

```promql
http_requests_total
```

или:

```promql
rate(http_requests_total[1m])
```

---

# Grafana

Grafana доступна по адресу:

```text
http://localhost:3000
```

Grafana используется для визуализации:

* HTTP-запросов
* скорости обработки запросов
* количества ошибок
* производительности приложения
* пользовательских метрик Pollify

---

## Добавление Prometheus в Grafana

В Grafana добавьте новый Data Source:

```text
Type: Prometheus
```

URL:

```text
http://prometheus:9090
```

Если Grafana и Prometheus находятся в одной Docker Compose сети, необходимо использовать имя сервиса:

```text
prometheus
```

а не:

```text
localhost
```

---

# Метрики Pollify

Приложение должно предоставлять endpoint:

```text
/metrics
```

Например:

```text
http://localhost:9091/metrics
```

Проверить доступность метрик можно командой:

```bash
curl http://localhost:9091/metrics
```

Для проверки конкретной метрики:

```bash
curl http://localhost:9091/metrics | grep http_requests_total
```

---

# Проверка подключения из контейнера Prometheus

Если Prometheus не получает метрики, проверьте доступность Pollify непосредственно из контейнера:

```bash
docker exec -it prometheus wget -qO- \
  http://host.docker.internal:9091/metrics
```

Если метрики отображаются, сетевое соединение работает корректно.

---

# Полезные команды

## Запустить инфраструктуру

```bash
make env-up
```

## Остановить инфраструктуру

```bash
make env-down
```

## Очистить данные PostgreSQL

```bash
make env-cleanup
```

## Открыть PostgreSQL на localhost

```bash
make env-port-forward
```

## Закрыть PostgreSQL port forwarder

```bash
make env-port-close
```

## Создать миграцию

```bash
make migrate-create seq=название_миграции
```

Например:

```bash
make migrate-create seq=create_users
```

## Применить миграции

```bash
make migrate-up
```

## Откатить миграцию

```bash
make migrate-down
```

## Запустить приложение

```bash
make pollify-run
```

---

# Быстрый старт

Полная последовательность запуска проекта:

```bash
# 1. Клонировать репозиторий
git clone <repository-url>

# 2. Перейти в директорию проекта
cd pollify

# 3. Создать и настроить .env
cp .env.example .env

# 4. Запустить инфраструктуру
make env-up

# 5. Открыть PostgreSQL на localhost
make env-port-forward

# 6. Применить миграции
make migrate-up

# 7. Запустить Pollify
make pollify-run
```

После запуска сервисы будут доступны по следующим адресам:

| Сервис              | Адрес                              |
| ------------------- | ---------------------------------- |
| Pollify             | зависит от конфигурации приложения |
| Pollify Metrics     | `http://localhost:9091/metrics`    |
| PostgreSQL          | `localhost:5432`                   |
| RabbitMQ            | `localhost:5672`                   |
| RabbitMQ Management | `http://localhost:15672`           |
| Redis               | `localhost:6379`                   |
| Prometheus          | `http://localhost:9090`            |
| Grafana             | `http://localhost:3000`            |

---

# Разработка

Для локальной разработки рекомендуется следующий workflow:

```bash
make env-up
make env-port-forward
make migrate-up
make pollify-run
```

Для остановки:

```bash
make env-port-close
make env-down
```
