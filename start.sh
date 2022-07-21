go env -w GOOS=linux
go build
docker-compose up --build -d
go clean