sudo chown $USER /var/run/docker.sock
docker-compose pull && docker-compose up -d
go run main.go
docker stop $(docker ps -aq) && docker rm $(docker ps -aq)

