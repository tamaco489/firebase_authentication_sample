#!/bin/bash

# .env_localstack の値を読み込む
set -a # 以降の環境変数の自動エクスポートを有効化
source /etc/localstack/init/ready.d/.env_localstack
set +a # 自動エクスポートを無効化

# Secrets Manager に Redis のシークレットを作成
awslocal secretsmanager create-secret \
  --name 'core/dev/redis-cluster' \
  --region ap-northeast-1 \
  --secret-string "{
    \"host\": \"${REDIS_HOST}\",
    \"port\": \"${REDIS_POOL}\",
    \"password\": \"${REDIS_PASSWORD}\",
    \"pool_size\": ${REDIS_POOL_SIZE}
  }"
