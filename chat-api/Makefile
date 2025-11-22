.PHONY: test up down logs

test:
	go test ./... -v

up:
	docker compose up --build -d

down:
	docker compose down -v

logs:
	docker compose logs -f app

