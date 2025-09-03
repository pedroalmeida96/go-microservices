# Go Microservices Project

This project was developed while following the course **[Build a Microservice with Go](https://www.linkedin.com/learning/build-a-microservice-with-go/go-for-microservices)** on LinkedIn Learning.  
The goal is to learn how to create and run a REST API in Go within a microservices environment using Docker and Docker Compose.

---

## 🚀 Prerequisites

- [Go](https://go.dev/) (1.18+)
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
