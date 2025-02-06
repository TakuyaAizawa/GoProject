# Next.js タスク管理フロントエンド

Cloud Run上で動作するNext.jsベースのフロントエンドアプリケーションです。モダンなUIとUXを提供し、GoバックエンドAPIと連携してタスク管理機能を実現しています。

## デプロイ済みURL

- アプリケーション: https://nextjs-frontend-896425711976.asia-northeast1.run.app/

## 技術スタック

- **フレームワーク**: Next.js 15.1.6
- **言語**: TypeScript
- **UIコンポーネント**:
  - Radix UI 
  - shadcn/ui 
  - Tailwind CSS 
- **HTTP クライアント**: Axios
- **デプロイ**: Google Cloud Run
- **その他**:
  - ESLint (コード品質管理)
  - Prettier (コードフォーマット)

## アーキテクチャ

```
task-frontend/
├── src/
│   ├── app/          # ページコンポーネント (App Router)
│   ├── components/   # 共通コンポーネント
│   │   ├── ui/      # 基本UIコンポーネント
│   │   └── forms/   # フォームコンポーネント
│   └── lib/         # ユーティリティ関数
├── public/          # 静的ファイル
└── styles/         # グローバルスタイル
```

## 主要機能

### タスク管理UI
- タスクの作成・編集フォーム
- タスク一覧表示
- タスクの状態管理UI
- タスクの優先度設定UI

### TODOリストUI
- TODOの作成・編集
- TODOリストの表示
- 完了状態の切り替え
- ドラッグ&ドロップによる並び替え

### UI/UXの特徴
- レスポンシブデザイン
- ダークモード対応
- アクセシビリティ対応
- スムーズなアニメーション

## セットアップ手順

### 必要条件
- Node.js 18.0.0以上
- npm または yarn
- Google Cloud SDKとアカウント（デプロイ時）

### ローカル開発環境のセットアップ

1. 環境変数の設定
```bash
cp .env.local.example .env.local
# .env.localファイルを編集して必要な値を設定
```

2. 依存パッケージのインストール
```bash
npm install
# または
yarn install
```

3. 開発サーバーの起動
```bash
npm run dev
# または
yarn dev
```

### 本番ビルド
```bash
npm run build
npm start
# または
yarn build
yarn start
```

## デプロイメント

Cloud Runへのデプロイは `deploy.sh` スクリプトを使用します：

```bash
./deploy.sh
```

## 開発ステータス

### 実装済み機能
- ✅ タスク管理のCRUD操作UI
- ✅ TODOリストのCRUD操作UI
- ✅ レスポンシブデザイン
- ✅ Cloud Runへの自動デプロイ
- ✅ バックエンドAPIとの連携

### 開発中の機能
- 🚧 ユーザー認証UI
- 🚧 テストコードの整備
- 🚧 パフォーマンス最適化

### 今後の計画
1. **UI/UX改善**
   - アニメーションの追加
   - ダークモードの実装
   - アクセシビリティの強化

2. **機能拡張**
   - オフライン対応
   - プッシュ通知
   - ドラッグ&ドロップ

3. **品質向上**
   - E2Eテストの追加
   - パフォーマンスモニタリング
   - エラーハンドリングの改善