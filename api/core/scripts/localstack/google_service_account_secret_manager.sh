#!/bin/bash

# .env_localstack の値を読み込む
set -a # 以降の環境変数の自動エクスポートを有効化
source /etc/localstack/init/ready.d/.env_localstack
set +a # 自動エクスポートを無効化

# Firebase シークレットの作成
awslocal secretsmanager create-secret \
  --name "core/dev/firebase-service-account" \
  --secret-string "$FIREBASE_SERVICE_ACCOUNT" \
  --region ap-northeast-1
