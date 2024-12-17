.PHONY: services-up
services-up:
	@echo "=== RUN DEVELOPMENT ENVIRONMENT (AUTO BUILD) - $$(date)"
	@echo "Fill DOCKER_SERVICE_NAME env with specific service name if want to run only that service in docker"
	@echo "=== Running CHOWN so that docker data can be overwritten"
	@sudo chown -R $$(id -u):$$(id -g) .docker
	@echo "=== Docker UP!"
	@docker-compose -f docker-compose.yml up --build $(DOCKER_SERVICE_NAME)

.PHONY: services-down
services-down:
	@docker-compose -f docker-compose.yml down

.PHONY: clear-postgres-docker-data
clear-postgres-docker-data:
	@sudo rm -rf .docker/postgres-data

.PHONY: download
download:
	@echo "📡 Starting download package dependencies"
	@GOSUMDB=off go mod download -x
	@sleep 1
	@echo "📁 Setup vendor directory"
	@go mod vendor
	@sleep 1
	@echo "👌 Download package completed"