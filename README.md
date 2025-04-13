# 📚 Bookstore API (Go + MongoDB/PostgreSQL + gRPC + JWT)

A production-ready Bookstore backend built with Golang using Clean Architecture principles. Supports both REST and gRPC, secured with JWT authentication, and ready for Docker/Kubernetes deployments.

---

## ✨ Features

- 🛠 REST & gRPC APIs
- 🔐 JWT Authentication
- 🧹 Clean Architecture
- 💾 PostgreSQL (via `gorm` or `sqlc`)
- 🧪 Unit & Integration Tests
- 📦 Docker + Docker Compose
- 📊 Swagger Docs + Observability Hooks
- 🚀 Kubernetes-ready

---

## 🧱 Project Structure

```bash
bookstore-api/
│
├── cmd/                # Main entry point (HTTP/gRPC server)
├── internal/           # Business logic modules (user, book, auth)
├── api/                # Protobuf definitions & gRPC code
├── pkg/                # Shared packages (logger, utils)
├── config/             # Environment & config management
├── test/               # Tests
├── migrations/         # SQL migrations
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
