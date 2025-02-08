#!/bin/bash

# .env.local から環境変数を読み込む
export $(cat .env.local | xargs)

# プロジェクトIDを設定
PROJECT_ID="${GCP_PROJECT_ID}"
# サービス名
SERVICE_NAME="frontend"
# リージョン
REGION="asia-northeast1"
# Artifact Registryのリポジトリ名
REPO_NAME="nextjs-apps"
# バックエンドのURL（Cloud RunのURL）
BACKEND_URL="${NEXT_PUBLIC_API_URL}"  # 環境変数から読み取り

# リポジトリが存在しない場合のみ作成
if ! gcloud artifacts repositories describe $REPO_NAME --location=$REGION &>/dev/null; then
    gcloud artifacts repositories create $REPO_NAME \
        --repository-format=docker \
        --location=$REGION
fi

# Dockerの認証設定
gcloud auth configure-docker $REGION-docker.pkg.dev

# Dockerイメージをビルド
docker build \
    --platform linux/amd64 \
    --build-arg NEXT_PUBLIC_API_URL=$BACKEND_URL \
    -t $REGION-docker.pkg.dev/$PROJECT_ID/$REPO_NAME/$SERVICE_NAME .

# イメージをプッシュ
docker push $REGION-docker.pkg.dev/$PROJECT_ID/$REPO_NAME/$SERVICE_NAME

# Cloud Runにデプロイ
gcloud run deploy $SERVICE_NAME \
    --image $REGION-docker.pkg.dev/$PROJECT_ID/$REPO_NAME/$SERVICE_NAME \
    --platform managed \
    --region $REGION \
    --allow-unauthenticated \
    --port 3000 \
    --set-env-vars "NEXT_PUBLIC_API_URL=$BACKEND_URL"