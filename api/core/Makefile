# =================================================================
# 実行環境設定する時用のやつ
# =================================================================
.PHONY: setup-env install-tools
setup-env:
	cp -p .env_sample .env

install-tools: setup-env
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install golang.org/x/tools/cmd/godoc@latest
	npm i -g @redocly/cli@latest


.PHONY: gen-api bundle-openapi
gen-api: bundle-openapi
	oapi-codegen -package gen -generate types -o internal/gen/types.gen.go ./spec/openapi.yaml
	oapi-codegen -package gen -generate strict-server,gin -templates ./spec/templates -o internal/gen/server.gen.go ./spec/openapi.yaml

bundle-openapi:
	redocly bundle ./spec/openapi_base.yaml --output ./spec/openapi.yaml


# =================================================================
# ローカル起動する時用のやつ
# =================================================================
.PHONY: lint deps build run
lint:
	golangci-lint run --timeout 3m

deps:
	go mod download && go mod tidy

build: deps
	go build -o build/core_api cmd/main.go

run: build
	./build/core_api


# =================================================================
# コンテナ起動する時用のやつ
# =================================================================
.PHONY: up down logs
up:
	docker compose up -d core-api

down:
	docker compose down core-api

logs:
	docker compose logs -f core-api


# =================================================================
# apiリクエストする時用のやつ（dev）
# =================================================================
.PHONY: api-all-request
sleep:
	@sleep 0.5

api-all-request: healthcheck sleep \
	create-user sleep \
	get-me sleep

.PHONY: healthcheck create-user get-me
healthcheck:
	curl -i -X 'GET' \
		'http://localhost:8080/core/v1/healthcheck' \
		-H 'accept: application/json'

create-user:
	curl -i -sX 'POST' \
		'http://localhost:8080/core/v1/users' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '{"provider_type": "firebase"}'

get-me:
	curl -i -X 'GET' \
		'http://localhost:8080/core/v1/users/me' \
		-H 'accept: application/json'
