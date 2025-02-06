# 依存関係ステージ
FROM --platform=linux/amd64 node:18-alpine AS deps
WORKDIR /app
COPY package*.json ./
# 本番用の依存関係のみをインストール
RUN npm ci --production

# ビルダーステージ
FROM --platform=linux/amd64 node:18-alpine AS builder
WORKDIR /app
# 開発用依存関係を含めてインストール
COPY package*.json ./
RUN npm ci
# ソースコードをコピー
COPY . .
# ビルド
ENV NODE_ENV=production
RUN npm run build

# 実行ステージ
FROM --platform=linux/amd64 node:18-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production
# 本番用の依存関係のみをコピー
COPY --from=deps /app/node_modules ./node_modules
# ビルド済みのアプリケーションをコピー
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/next.config.js ./next.config.js

EXPOSE 3000
CMD ["npm", "start"]