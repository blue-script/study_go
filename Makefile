.PHONY: server compose-up compose-build compose-down docker-build logs

server:
	set -a && . ./.env.local && set +a && go run main.go

compose-build:
	docker compose build

postgres-up:
	docker compose up -d postgres

postgres-down:
	docker compose down postgres

postgres-run:
	docker compose start -d postgres

postgres-stop:
	docker compose stop postgres

compose-up:
	docker compose up --build

compose-down:
	docker compose down

logs:
	docker compose logs -f