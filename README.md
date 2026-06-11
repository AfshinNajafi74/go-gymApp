# 🏋️ Go Gym Backend

Backend service for a Gym Management Application built with Go.

This project is being developed as a learning-focused, production-inspired backend using Clean Architecture principles.

---

## 🏗 Architecture

* Clean Architecture
* Repository Pattern
* Service Layer
* Dependency Injection

---

## 🛠 Tech Stack

* Go
* PostgreSQL
* GORM
* Gorilla Mux
* JWT
* Swagger

---

## 🚀 Getting Started

```bash
go mod tidy
go run ./cmd/server
```

Server will start on:

```text
http://localhost:8080
```

---

## 📖 API Documentation

Generate Swagger docs:

```bash
swag init --generalInfo ./cmd/server/main.go --output ./docs
```

Swagger UI:

```text
http://localhost:8080/swagger/index.html
```

---

## 🤝 Contributing

Contributions, ideas, and suggestions are always welcome.

If you're interested in backend development with Go, feel free to open an issue, submit a pull request, or share your feedback.

---

⭐ If you find this project useful, consider giving it a star.
