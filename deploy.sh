#!/bin/bash

# .envファイルが存在する場合に読み込む
if [ -f .env ]; then
    export $(cat .env | grep -v '#' | xargs)
fi

# 環境変数が設定されているか確認
if [ -z "$DATABASE_URL" ]; then
    echo "Error: DATABASE_URL is not set"
    echo "Please create .env file from .env.example"
    exit 1
fi

# プロジェクトIDを既存のものに設定
PROJECT_ID="striking-decker-449414-a6"
# サービス名
SERVICE_NAME="go-service"
# リージョン
REGION="asia-northeast1"
# Artifact Registryのリポジトリ名
REPO_NAME="go-apps"

# リポジトリが存在しない場合のみ作成
if ! gcloud artifacts repositories describe $REPO_NAME --location=$REGION &>/dev/null; then
    gcloud artifacts repositories create $REPO_NAME \
        --repository-format=docker \
        --location=$REGION
fi

# Dockerの認証設定
gcloud auth configure-docker $REGION-docker.pkg.dev

# Dockerイメージをビルド（プラットフォームを明示的に指定）
docker build --platform linux/amd64 -t $REGION-docker.pkg.dev/$PROJECT_ID/$REPO_NAME/$SERVICE_NAME .

# イメージをプッシュ
docker push $REGION-docker.pkg.dev/$PROJECT_ID/$REPO_NAME/$SERVICE_NAME

# Cloud Runにデプロイ（環境変数を含める）
gcloud run deploy $SERVICE_NAME \
    --image $REGION-docker.pkg.dev/$PROJECT_ID/$REPO_NAME/$SERVICE_NAME \
    --platform managed \
    --region $REGION \
    --allow-unauthenticated \
    --set-env-vars "DATABASE_URL=${DATABASE_URL}" \
    --port 8080