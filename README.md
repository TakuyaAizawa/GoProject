# Cloud Run タスク管理アプリケーション

このプロジェクトは、GoバックエンドとNext.jsフロントエンドを組み合わせたモダンなタスク管理アプリケーションです。
マイクロサービスアーキテクチャを採用し、Google Cloud Runを活用したスケーラブルな構成となっています。

## デプロイ済みURL

- フロントエンド: https://nextjs-frontend-896425711976.asia-northeast1.run.app/

## 技術スタック

### バックエンド (`backend/`)
- **言語**: Go 1.21
- **アーキテクチャ**: クリーンアーキテクチャベース
- **データベース**: PostgreSQL (Supabase)
- **主要パッケージ**:
  - `lib/pq`: PostgreSQLドライバー
  - `godotenv`: 環境変数管理
- **デプロイ**: Google Cloud Run

### フロントエンド (`task-frontend/`)
- **フレームワーク**: Next.js 15.1.6
- **言語**: TypeScript
- **UIコンポーネント**:
  - Radix UI
  - shadcn/ui
  - Tailwind CSS
- **HTTP クライアント**: Axios
- **デプロイ**: Google Cloud Run

## アーキテクチャ設計

### バックエンド構成
```
backend/
├── cmd/            # エントリーポイント
├── internal/       # 内部パッケージ
│   ├── config/     # 設定管理
│   ├── handler/    # HTTPハンドラ
│   ├── middleware/ # ミドルウェア
│   ├── model/      # ドメインモデル
│   └── repository/ # データアクセス層
└── pkg/            # 外部公開可能なパッケージ
```

### フロントエンド構成
```
task-frontend/
├── src/
│   ├── app/       # ページコンポーネント
│   ├── components/# 共通コンポーネント
│   └── lib/       # ユーティリティ
├── public/        # 静的ファイル
└── .env.local     # 環境変数設定
```

## 主要機能

### タスク管理機能
- タスクの作成・編集・削除
- タスク一覧の表示
- タスクの状態管理
- タスクの優先度設定

### TODOリスト機能
- TODOの作成・編集・削除
- TODOリストの表示
- 完了状態の管理

## セットアップ手順

### バックエンド
1. 環境変数の設定
```bash
cd backend
cp .env.example .env
# .envファイルを編集して必要な値を設定
```

2. ローカル開発環境の起動
```bash
go mod download
go run cmd/api/main.go
```

### フロントエンド
1. 環境変数の設定
```bash
cd task-frontend
cp .env.local.example .env.local
# .env.localファイルを編集して必要な値を設定
```

2. 依存パッケージのインストール
```bash
npm install
```

3. 開発サーバーの起動
```bash
npm run dev
```

## デプロイメント

### Cloud Runへのデプロイ
各ディレクトリの `deploy.sh` スクリプトを使用してデプロイを実行できます：

```bash
# バックエンドのデプロイ
cd backend
./deploy.sh

# フロントエンドのデプロイ
cd task-frontend
./deploy.sh
```

## 開発ステータス

### 実装済み機能
- ✅ タスク管理のCRUD操作
- ✅ TODOリストのCRUD操作
- ✅ モダンなUI実装
- ✅ Cloud Runへの自動デプロイ
- ✅ CORSサポート

### 開発中の機能
- 🚧 ユーザー認証（JWT）
- 🚧 テストコードの整備
- 🚧 CI/CDパイプライン

### 今後の計画
1. **品質向上**
   - テストカバレッジの向上
   - エラーハンドリングの強化
   - パフォーマンス最適化

2. **機能拡張**
   - ユーザー認証・認可の実装
   - APIドキュメントの自動生成
   - タスクの共有機能

3. **インフラ整備**
   - 監視・アラートの設定
   - ログ収集の整備
   - インフラのコード化



