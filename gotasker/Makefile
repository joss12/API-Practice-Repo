run:
	go run ./cmd/server/main.go

build:
	docker build -t gotasker .

up:
	docker-compose up -d --build

down:
	docker-compose down -v

logs:
	docker-compose logs -f app

clean:
	docker system prune -f

test:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -func=coverage.out
