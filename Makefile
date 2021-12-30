# ==============================================================================
# Docker

docker_dev:
	@echo Starting local docker dev compose
	docker-compose -f docker-compose.yaml up --build

local:
	@echo Starting local docker compose
	docker-compose -f docker-compose.local.yaml up -d --build

