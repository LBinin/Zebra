.PHONY: dev

NOW = $(shell date -u '+%Y%m%d%I%M%S')
RELEASE_VERSION = v0.0.0
APP 			= shop

dev:
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s && ./bin/air

docker-compose:
	@docker-compose up -d --remove-orphans --build

build:
	@go build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(SERVER_BIN) ./cmd/${APP}

swagger:
	@swag init --parseDependency --generalInfo ./cmd/${APP}/main.go --output ./internal/app/swagger
