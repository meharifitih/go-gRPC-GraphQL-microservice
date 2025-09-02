.PHONY: help build up down logs clean rebuild

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build all Docker images
	docker-compose build

up: ## Start all services
	docker-compose up -d

down: ## Stop all services
	docker-compose down

logs: ## Show logs for all services
	docker-compose logs -f

clean: ## Remove all containers, networks, and volumes
	docker-compose down -v --remove-orphans
	docker system prune -f

rebuild: ## Rebuild and restart all services
	docker-compose down
	docker-compose build --no-cache
	docker-compose up -d

status: ## Show status of all services
	docker-compose ps

restart: ## Restart all services
	docker-compose restart

# Individual service commands
account-up: ## Start only account service
	docker-compose up -d account account_db

catalog-up: ## Start only catalog service
	docker-compose up -d catalog catalog_db

order-up: ## Start only order service
	docker-compose up -d order order_db

graphql-up: ## Start only graphql service
	docker-compose up -d graphql
