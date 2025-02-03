.PHONY: build up down logs clean run

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

clean:
	docker-compose down -v --remove-orphans
	docker system prune -f
	docker volume prune -f

run: build up