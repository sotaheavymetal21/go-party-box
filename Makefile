APP_CONTAINER_NAME = app
DB_CONTAINER_NAME = db
RUN_APP = docker-compose exec $(APP_CONTAINER_NAME)

prepare:
	docker-compose up -d --build

up:
	docker-compose up -d

build:
	docker-compose build

down:
	docker-compose down

app:
	docker exec -it $(APP_CONTAINER_NAME) bash

db:
	docker exec -it $(DB_CONTAINER_NAME) bash
