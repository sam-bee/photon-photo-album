# Get the current user's UID and GID
DOCKERUID := $(shell id -u)
DOCKERGID := $(shell id -g)

.PHONY: setup down clean shell-frontend shell-backend shell-ml shell-mq

setup:
	if [ ! -f .env ]; then cp .env.dist .env; fi
	DOCKERUID=$(DOCKERUID) DOCKERGID=$(DOCKERGID) docker compose build --no-cache
	DOCKERUID=$(DOCKERUID) DOCKERGID=$(DOCKERGID) docker compose up -d

# Stop containers
down:
	docker compose down

# Stop containers and remove volumes
clean:
	docker compose down -v
	docker compose kill
	docker compose rm

# Shell access to containers
shell-frontend:
	docker compose exec frontend /bin/bash

shell-backend:
	docker compose exec backend /bin/bash

shell-ml:
	docker compose exec ml /bin/bash

shell-mq:
	docker compose exec rabbitmq /bin/bash
