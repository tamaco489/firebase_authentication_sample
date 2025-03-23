# =================================================================
# 実行環境設定する時用のやつ
# =================================================================
.PHONY: setup-env install-tools
setup-env:
	@cp -p .env_sample .env
	@cp .env_localstack_sample .env_localstack

install-tools: setup-env
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install golang.org/x/tools/cmd/godoc@latest
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install go.uber.org/mock/mockgen@latest
	npm i -g @redocly/cli@latest


# =================================================================
# 自動生成する時用のやつ
# =================================================================
.PHONY: gen-api bundle-openapi gen-sqlc
gen-api: bundle-openapi
	oapi-codegen -package gen -generate types -o internal/gen/types.gen.go ./spec/openapi.yaml
	oapi-codegen -package gen -generate strict-server,gin -templates ./spec/templates -o internal/gen/server.gen.go ./spec/openapi.yaml

bundle-openapi:
	redocly bundle ./spec/openapi_base.yaml --output ./spec/openapi.yaml

gen-sqlc:
	sqlc generate -f internal/repository/sqlc.yaml
	mockgen -source=./internal/repository/gen_sqlc/querier.go -package mock -destination=./internal/repository/gen_mock/mock_querier.go
	mockgen -source=./internal/repository/gen_sqlc/db.go -package mock -destination=./internal/repository/gen_mock/mock_db.go

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
	docker compose up -d core-api mysql redis localstack

down:
	docker compose down core-api mysql redis localstack

logs:
	docker compose logs -f core-api

# secret managerを操作する時のやつ
.PHONY: list-secrets get-secret-value
list-secrets:
	docker exec -it core-localstack awslocal secretsmanager list-secrets --region ap-northeast-1

get-secret-value:
	docker exec -it core-localstack awslocal secretsmanager get-secret-value --secret-id core/dev/rds-cluster --region ap-northeast-1
	docker exec -it core-localstack awslocal secretsmanager get-secret-value --secret-id core/dev/redis-cluster --region ap-northeast-1
	docker exec -it core-localstack awslocal secretsmanager get-secret-value --secret-id core/dev/firebase-service-account --region ap-northeast-1


# =================================================================
# MySQL操作する時のやつ
# =================================================================
# mysqlにアクセス
.PHONY: mysql
mysql:
	MYSQL_ROOT_PASSWORD=password#0 docker compose exec mysql mysql -u root -ppassword#0 -h localhost -P 33306 dev_core

# MySQLのテーブルをマイグレーションする時のやつ
.PHONY: migrate-new migrate-up migrate-down migrate-status
migrate-new:
	sql-migrate new -env='dev' -config=./_tools/sql-migrate/config.yaml $(NAME)

migrate-up:
	sql-migrate up -env='dev' -config=./_tools/sql-migrate/config.yaml

migrate-down:
	sql-migrate down -limit=1 -env='dev' -config=./_tools/sql-migrate/config.yaml

migrate-status:
	sql-migrate status -env='dev' -config=./_tools/sql-migrate/config.yaml


# =================================================================
# Redis操作する時のやつ
# =================================================================
.PHONY: redis redis-v redis-get-all redis-clear-all redis-get-by-key
redis:
	docker exec -it core-redis redis-cli

redis-v:
	docker-compose exec redis redis-server --version

redis-get-all:
	docker-compose exec redis redis-cli KEYS '*'

redis-clear-all:
	docker-compose exec redis redis-cli FLUSHALL

# e.g. make redis-get-by-key sub=2iSI3im4bcOFJDoT7E9QLebbU9G2
redis-get-by-key:
	docker-compose exec redis redis-cli -n 0 GET session:${sub}

# e.g. make redis-get-ttl-by-key sub=2iSI3im4bcOFJDoT7E9QLebbU9G2
redis-get-ttl-by-key:
	docker-compose exec redis redis-cli -n 0 TTL session:${sub}

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
		-d '{"provider_type": "firebase"}' \
		-w "\nTotal time: %{time_total}s\n"

get-me:
	curl -i -sX 'GET' \
		'http://localhost:8080/core/v1/users/me' \
		-H 'accept: application/json' \
		-H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjMwYjIyMWFiNjU2MTdiY2Y4N2VlMGY4NDYyZjc0ZTM2NTIyY2EyZTQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vdGFtYWNvNDg5LWF1dGgtc2FtcGxlIiwiYXVkIjoidGFtYWNvNDg5LWF1dGgtc2FtcGxlIiwiYXV0aF90aW1lIjoxNzQyNzIzODA1LCJ1c2VyX2lkIjoibTNrV3RLVmlmU1BKZGRaZTcwMnBrdm9kQXJPMiIsInN1YiI6Im0za1d0S1ZpZlNQSmRkWmU3MDJwa3ZvZEFyTzIiLCJpYXQiOjE3NDI3MzE2ODMsImV4cCI6MTc0MjczNTI4MywiZW1haWwiOiJ0YW1hY280ODlAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbInRhbWFjbzQ4OUBnbWFpbC5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9fQ.bRZoNJe5OsOQymfmXZ8CXgUVkb_2VUa6ffixvH4zva_gI0kguaEFjxR6PpgHlkGHkBDdd9pLd_p6_VC0e6KoX3TouRwCbGJC1VEvZbbaLBY4tqe9cU5L3fvbmjhaa434E8LH99Qqh31EbYui0H8IUqXr4PadM6n5MxhubaBNaN3NyD8p_g47BQhnJTLItMrHB_lBOj4x-f0x_WVvFPTfrPbt9VjL-WLqueoRCZhKr-6__Z0QrZyRwj_OZT_GQp5dYxtfWGMHi11vwGOiakyh0lp3vCD3IOaThYCH6epv5xJYtb6r17hUudQ3REVEeCnlEH1aAIKvYNcxh2iT63IM2g' \
		-w "\nTotal time: %{time_total}s\n"
