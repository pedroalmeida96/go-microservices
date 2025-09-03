# Go Microservices Project

This project was developed while following the course **[Build a Microservice with Go](https://www.linkedin.com/learning/build-a-microservice-with-go/go-for-microservices)** on LinkedIn Learning.  
The goal is to learn how to create and run a simple **CRUD REST API** in Go within a microservices environment using Postgres, Docker and Docker Compose.

The service is built with:
- [Echo](https://echo.labstack.com/) — Web framework for building APIs
- [GORM](https://gorm.io/) — ORM library for database operations
- [PostgreSQL](https://www.postgresql.org/) — Relational database for persistence

---

## 🚀 Prerequisites

- [Go](https://go.dev/)
- GORM and Echo
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 📦 Setup

```bash
sudo chown $USER /var/run/docker.sock
docker-compose pull && docker-compose up -d
go run main.go
docker stop $(docker ps -aq) && docker rm $(docker ps -aq)
```


## 📚 Reference

- Course: [Build a Microservice with Go (LinkedIn Learning)](https://www.linkedin.com/learning/build-a-microservice-with-go/go-for-microservices)
