# Get the current user's UID and GID
DOCKERUID := $(shell id -u)
DOCKERGID := $(shell id -g)

.PHONY: setup down clean shell-frontend shell-backend shell-ml shell-mq

setup:
	if [ ! -f .env ]; then cp .env.dist .env; fi
	build
	up

# Build and start containers
--up:
	DOCKERUID=$(DOCKERUID) DOCKERGID=$(DOCKERGID) docker compose up -d

# Build containers with no cache
--build:
	DOCKERUID=$(DOCKERUID) DOCKERGID=$(DOCKERGID) docker compose build --no-cache

# Stop containers
down:
	docker compose down

# Stop containers and remove volumes
clean:
	docker compose down -v
	docker system prune -f

# Shell access to containers
shell-frontend:
	docker compose exec frontend /bin/bash

shell-backend:
	docker compose exec backend /bin/bash

shell-ml:
	docker compose exec ml /bin/bash

shell-mq:
	docker compose exec rabbitmq /bin/bash
