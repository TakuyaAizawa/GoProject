# Go Cloud Run Backend

Cloud Run上で動作するGo言語のバックエンドAPIサービスです。クリーンアーキテクチャの原則に従い、保守性と拡張性を重視した設計になっています。

## 技術スタック

- **言語**: Go 1.21
- **データベース**: PostgreSQL (Supabase)
- **インフラ**: Google Cloud Platform (Cloud Run)
- **その他**:
  - JWT認証（予定）
  - GitHub Actions（予定）

## アーキテクチャ

クリーンアーキテクチャの原則に基づき、以下のような構造で実装しています：

```
.
├── cmd/
│   └── api/          # アプリケーションのエントリーポイント
├── internal/         # 内部パッケージ
│   ├── config/       # 設定管理
│   ├── handler/      # HTTPハンドラー
│   ├── middleware/   # ミドルウェア
│   ├── model/        # データモデル
│   └── repository/   # データアクセス層
└── pkg/             # 外部から利用可能なパッケージ
    └── database/     # データベース接続管理
```

### レイヤー構造

1. **Handlers** (`internal/handler/`)
   - HTTPリクエストの受付
   - リクエストのバリデーション
   - レスポンスの整形

2. **Repositories** (`internal/repository/`)
   - データベースアクセスの抽象化
   - CRUDオペレーションの実装

3. **Models** (`internal/model/`)
   - ドメインモデルの定義
   - バリデーションルール

## 機能一覧

### 現在の機能
- タスク管理 (CRUD)
  - タスクの作成
  - タスクの取得（単一/一覧）
  - タスクの更新
  - タスクの削除
- TODOリスト管理 (CRUD)
  - TODOの作成
  - TODOの取得（単一/一覧）
  - TODOの更新
  - TODOの削除
- ヘルスチェックエンドポイント
- CORSサポート

### 今後追加予定の機能
- [ ] JWT認証
- [ ] バリデーションの強化
- [ ] テストコードの整備
- [ ] CI/CD (GitHub Actions)
- [ ] OpenAPI (Swagger) ドキュメント
- [ ] エラーハンドリングの改善
- [ ] ロギングの強化
- [ ] メトリクス収集

## 技術選定の理由

### クリーンアーキテクチャの採用
- **理由**: 
  - コードの責務を明確に分離し、保守性を向上
  - テストが書きやすい構造
  - 依存関係が明確で理解しやすい

### PostgreSQL (Supabase)
- **理由**:
  - 信頼性の高いRDBMS
  - Supabaseによる管理の容易さ
  - 無料枠での開発が可能

### Cloud Run
- **理由**:
  - コンテナベースで簡単にスケール
  - 従量課金で開発コストを抑制
  - デプロイが容易

## セットアップ

### 必要条件
- Go 1.21以上
- PostgreSQL
- Google Cloud SDKとアカウント

### ローカル開発環境のセットアップ

1. リポジトリのクローン
```bash
git clone [repository-url]
cd go-cloud-run
```

2. 環境変数の設定
```bash
cp .env.example .env
# .envファイルを編集して必要な値を設定
```

3. アプリケーションの起動
```bash
go run cmd/api/main.go
```

### デプロイ

```bash
# Cloud Runへのデプロイ（準備中）
```

## API エンドポイント

### タスク管理
- `GET /api/tasks` - タスク一覧の取得
- `POST /api/tasks` - 新規タスクの作成
- `GET /api/task?id={id}` - 特定のタスクの取得
- `PUT /api/task?id={id}` - タスクの更新
- `DELETE /api/task?id={id}` - タスクの削除

### TODOリスト
- `GET /api/todos` - TODO一覧の取得
- `POST /api/todos` - 新規TODOの作成
- `GET /api/todo?id={id}` - 特定のTODOの取得
- `PUT /api/todo?id={id}` - TODOの更新
- `DELETE /api/todo?id={id}` - TODOの削除

### その他
- `GET /health` - ヘルスチェック

## 今後の展望

1. **品質向上**
   - テストカバレッジの向上
   - エラーハンドリングの強化
   - パフォーマンス計測と改善

2. **機能拡張**
   - ユーザ認証・認可の実装
   - APIドキュメントの自動生成
   - バッチ処理の実装

3. **インフラ整備**
   - CI/CDパイプラインの構築
   - 監視・アラートの設定
   - IaCの導入

## ライセンス

MIT License 