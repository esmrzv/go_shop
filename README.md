# Go Shop 🛒

Учебный backend-проект интернет-магазина на Go.  
Проект сделан для практики **чистой архитектуры**, **HTTP-сервисов**, **PostgreSQL**,  
а также **конкурентности (goroutines, channels, mutex)** на реальных кейсах.

---

## 🚀 Стек технологий

- Go 1.22+
- net/http
- PostgreSQL
- database/sql + pgx
- JWT (authentication)
- Context
- Channels / Goroutines
- Mutex
- Graceful shutdown
- .env конфигурация

---

## 📦 Архитектура

Проект построен по слоям:

cmd/
└── app/ // main.go

internal/
├── config/ // конфигурация приложения
├── db/ // подключение к PostgreSQL
├── model/ // доменные модели
├── repository/ // работа с БД
├── service/ // бизнес-логика
├── http/
│ ├── handler/ // HTTP handlers
│ ├── routers/ // регистрация роутов
│ └── middleware/ // JWT middleware
├── cartworker/ // cart через channels (actor model)
└── orderworker/ // worker pool для заказов



**Handler → Service → Repository**  
HTTP слой не знает о БД, сервисы не знают о HTTP.

---

## 🧠 Основные сущности

### User
- регистрация
- логин
- JWT авторизация

### Product
- создание продукта
- получение списка

### Category
- категории продуктов

### Cart
- корзина пользователя
- реализована **двумя способами**:
  - через `sync.Mutex`
  - через `channels + worker` (actor model)

### Order
- создание заказа из корзины
- асинхронная обработка через worker pool

---

## 🔄 Конкурентность

В проекте реализованы и сравнены два подхода:

### Mutex
- прямой доступ к данным
- защита через `sync.Mutex`
- проще, но требует аккуратности

### Channels (Actor Model)
- один goroutine владеет состоянием
- все операции — через сообщения
- отсутствие data race по дизайну

Корзина (`Cart`) реализована через оба подхода для практики и сравнения.

---

## 🔐 Аутентификация

- JWT
- middleware для защиты эндпоинтов
- user_id передаётся через `context.Context`

---

## ⚙️ Конфигурация

Используется `.env` файл:

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_shop
DB_SSLMODE=disable

JWT_SECRET=supersecretkey

▶️ Запуск проекта

go mod tidy
go run cmd/app/main.go


📡 Основные эндпоинты

POST /register
POST /login

Products
POST /products        (JWT)
GET  /products        (JWT)

Categories
POST /categories      (JWT)
GET  /categories      (JWT)

Cart
POST /cart/add        (JWT)
GET  /cart            (JWT)

🧭 Статус проекта

Проект завершён как учебный.
Дальнейшее развитие не планируется — используется как база знаний и reference.