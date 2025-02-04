.PHONY: build up down logs clean run

export DB_USER
export DB_PASS
export API_KEY

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

run: DB_USER=$(DB_USER) DB_PASS=$(DB_PASS) API_KEY=$(API_KEY) make build up