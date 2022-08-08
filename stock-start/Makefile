GRAPH_BINARY=graphApp
MESSAGE_BINARY=messageApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: build_graph build_message
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_graph: builds the graph binary as a linux executable
build_graph:
	@echo "Building graph binary..."
	cd ../graph-gateway && env GOOS=linux CGO_ENABLED=0 go build -o ${GRAPH_BINARY} .
	@echo "Done!"
	
## build_message: builds the message binary as a linux executable
build_message:
	@echo "Building message binary..."
	cd ../message-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MESSAGE_BINARY} .
	@echo "Done!"

