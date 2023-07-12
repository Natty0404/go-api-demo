.PHONY: init up down stop ps mod-tidy

build:
	docker-compose build

init:
	@make up
	@make ps

up:
	docker compose up -d

down:
	docker compose down
	
stop:
	docker-compose stop
	
ps:
	docker compose ps

serve:
	docker-compose exec go-api-demo-app go run main.go

mod-tidy:
	docker-compose exec go-api-demo-app go mod tidy