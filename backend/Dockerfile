# ビルドステージ
FROM --platform=linux/amd64 golang:1.21-alpine AS builder

WORKDIR /app

# go.mod をコピー
COPY go.mod ./

# ソースコードをコピー
COPY . .

# AMD64向けにバイナリをビルド
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go-server

# 実行ステージ
FROM --platform=linux/amd64 alpine:latest

# CA証明書のインストール
RUN apk --no-cache add ca-certificates

WORKDIR /

# ビルドステージからバイナリをコピー
COPY --from=builder /go-server /go-server

# コンテナのポートを公開
EXPOSE 8080

# 実行コマンドの設定
ENTRYPOINT ["/go-server"]
