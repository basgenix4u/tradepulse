<div align="center">

# 📈 TradePulse

**A real-time market streaming platform — Go backend, WebSockets, and a reactive trading dashboard.**

[![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![React](https://img.shields.io/badge/React-18-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)](https://react.dev)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)](./LICENSE)

</div>

---

## ✨ Overview

TradePulse is a real-time trading platform built for low-latency market data. A Go service streams price ticks over **WebSockets** to a React dashboard, with JWT-secured accounts and PostgreSQL persistence underneath.

It is deliberately small and readable — a clean reference for streaming architecture rather than a production exchange.

---

## 🚀 Features

- **Live price streaming** — a WebSocket hub broadcasts ticks to every connected client
- **JWT authentication** — signed access tokens with HTTP-only cookie support
- **Market endpoints** — REST routes for instruments, candles, and quotes
- **PostgreSQL persistence** — accounts, sessions, and market history
- **Containerised stack** — one `docker compose up` brings the whole system online
- **Typed Go internals** — clean separation across `config`, `db`, `handlers`, `middleware`, `models`, and `websocket`

---

## 🛠 Tech Stack

| Layer | Technology |
| --- | --- |
| Backend | Go 1.22, standard library + `gorilla/websocket` |
| Frontend | React 18, Vite, Tailwind CSS |
| Database | PostgreSQL 15 |
| Auth | JWT |
| Packaging | Docker + Docker Compose |
| CI | GitHub Actions (build, vet, test) |

---

## ⚡ Quick Start

### With Docker (recommended)

```bash
docker compose up --build
# → frontend on :5173, API on :8080
```

### Manual

**Backend**

```bash
cd backend
go mod download
go run ./cmd/server
# → http://localhost:8080
```

**Frontend**

```bash
cd frontend
npm install
npm run dev
# → http://localhost:5173
```

---

## 🔐 Environment Variables

Every secret is **required** — the stack refuses to start with placeholder credentials.

```bash
# .env (read by docker-compose.yml)
POSTGRES_USER=tradepulse
POSTGRES_PASSWORD=<generate-me>
POSTGRES_DB=tradepulse
JWT_SECRET=<generate-me>
```

Generate strong values with:

```bash
openssl rand -base64 32
```

Then:

```bash
docker compose up
```

---

## 📁 Project Structure

```text
tradepulse/
├── backend/
│   ├── cmd/server/         # Application entrypoint
│   └── internal/
│       ├── auth/           # JWT issuing & verification
│       ├── config/         # Environment configuration
│       ├── db/             # Postgres connection + migrations
│       ├── handlers/       # Auth and market HTTP handlers
│       ├── middleware/     # Request authentication
│       ├── models/         # Domain structs
│       └── websocket/      # Hub — client registry + broadcast
├── frontend/
│   ├── src/
│   │   ├── App.jsx         # Dashboard shell
│   │   └── main.jsx
│   └── vite.config.js
├── docker-compose.yml
└── .github/workflows/ci.yml
```

---

## 🧪 Testing & CI

```bash
cd backend
go vet ./...
go test ./...
```

GitHub Actions runs `go build`, `go vet`, and `go test` on every push and pull request.

---

## 🗺 Roadmap

- [ ] Order placement and portfolio tracking
- [ ] Historical candle backfill
- [ ] Rate limiting on the WebSocket hub
- [ ] Horizontal scaling via Redis pub/sub

---

## 📄 License

Released under the [MIT License](./LICENSE).

---

<div align="center">

Built by [Abdulbasit Abdulalim](https://github.com/basgenix4u)

</div>
