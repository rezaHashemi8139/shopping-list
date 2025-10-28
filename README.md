# 🧩 Go Modular API

A clean, modular boilerplate for building APIs using **Go + Gin + PostgreSQL + JWT + WebSocket**.  
Modular architecture with clear separation of layers (Handler, Service, Repository).

---

## 🚀 Features

- Modular & scalable architecture (Module Pattern)
- Layered design: Handler / Service / Repository / Model
- JWT Authentication ready
- Swagger for API documentation
- WebSocket support for real-time communication
- PostgreSQL integration with GORM
- Docker support for easy development

---

## 📂 Project Structure

```
shopping-list/
│
├── cmd/                  # CLI commands (future use)
├── config/               # App configuration
│   └── config.go
│
├── modules/              # Feature modules (auth, user, chat, ...)
│   ├── auth/
│   │   ├── handlers/
│   │   ├── services/
│   │   ├── repositories/
│   │   ├── models/
│   │   └── routes.go
│   ├── user/
│   └── chat/
│
├── routes/               # API route registration
│   └── router.go
│
├── utils/                # Utilities (JWT, logger, helpers)
├── docs/                 # Swagger documentation
├── main.go               # Application entrypoint
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Setup & Run

### 1️⃣ Prerequisites

- [Go](https://go.dev/dl/) 1.20+
- [Docker](https://www.docker.com/) (optional for PostgreSQL)

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

### 3️⃣ Start Database (optional via Docker)

```bash
docker compose up -d
```

You can also use your own PostgreSQL instance and set the `DATABASE_URL` environment variable.

### 4️⃣ Run the App

```bash
go run main.go
```

### 5️⃣ Test Health Endpoint

Visit:

```
http://localhost:5555/health
```

Response:

```json
{ "status": "ok" }
```

---

## 🔑 Environment Variables

| Variable       | Description                  | Default                                                            |
| -------------- | ---------------------------- | ------------------------------------------------------------------ |
| `PORT`         | Server port                  | `5555`                                                             |
| `DATABASE_URL` | PostgreSQL connection string | `postgres://postgres:password@localhost:5432/mydb?sslmode=disable` |
| `JWT_SECRET`   | JWT secret key               | `supersecretkey`                                                   |

---

## 🧠 Architecture & Design Patterns

This project follows several key software architecture patterns:

- **Layered Architecture** — separation of Handler, Service, and Repository layers.
- **Dependency Injection** — loose coupling via struct injection.
- **Repository Pattern** — abstracted data access.
- **Module Pattern (NestJS style)** — self-contained feature modules.
- **Observer Pattern** — used for WebSocket hub & real-time events.

---

## 🧩 Modules Overview

| Module | Description                        |
| ------ | ---------------------------------- |
| `auth` | JWT authentication & authorization |
| `user` | User management                    |
| `chat` | WebSocket-based real-time chat     |

---

## 🔮 Roadmap

1. ✅ Basic modular skeleton (this step)
2. 🧱 Add PostgreSQL integration with GORM
3. 🔐 Implement JWT authentication (login/register)
4. 📜 Add Swagger documentation
5. 🔌 Add WebSocket chat module
6. 🧩 Add unit tests & CI setup

---

## 🧑‍💻 Author

**[Reza Hashemi](https://github.com/rezaHashemi8139/)**  
Base stack: Go 1.20+, Gin, GORM, PostgreSQL

---

## 📄 License

MIT License — free to use, learn, and modify.
