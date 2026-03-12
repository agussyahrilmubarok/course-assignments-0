dev:
	go run cmd/course/main.go

build:
	go build -o bin/course cmd/course/main.go

compose/postgres/up:
	docker compose -f docker-compose-postgres.yaml up -d --build

compose/postgres/down:
	docker compose -f docker-compose-postgres.yaml down
	docker volume prune -a

compose/stage/up:
	docker compose -f docker-compose-postgres.yaml up -d --build

compose/stage/down:
	docker compose -f docker-compose-postgres.yaml down
	docker volume prune -a

compose/up:
	docker compose up -d --build

compose/down:
	docker compose down
	docker volume prune -a