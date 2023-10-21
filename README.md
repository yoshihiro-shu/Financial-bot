# financial-bot

金融に関する情報をLine Botで定期的に送信するサービス

![DALL·E 2023-10-21 16 56 36 - Illustration of a user receiving notifications on their wide smartphone screen  Icons representing financial news, stocks, FX, and crypto assets hover](https://github.com/yoshihiro-shu/financial-bot/assets/84740493/d56c809c-40f8-4298-b47c-085feb60ccae)

## Design Docs

[こちら](./docs/design_docs.md)

## 環境構築

[こちら](./docs/setting_environment.md)

## System Architecture

<img width="432" alt="image" src="https://github.com/yoshihiro-shu/financial-bot/assets/84740493/7b2edc84-c5fc-4b24-9c29-79df19d40b97">

## ディレクトリ構成

```zsh
project-name/
├── api/                  # API定義とプロトコルファイル
├── build/                # ビルドとデプロイのスクリプト
├── cmd/                  # プロジェクトのエントリーポイント (main.goファイル)
├── configs/              # 設定ファイル
├── deployments/          # KubernetesやDockerのデプロイメント設定
├── docs/                 # ドキュメント
├── images/               # Dockerfile
├── internal/             # 内部パッケージ
│   ├── app/              # アプリケーションロジック
│   ├── dao/              # データアクセスオブジェクト (DAO)
│   ├── middleware/       # ミドルウェア
│   ├── model/            # データモデル
│   └── service/          # サービスレイヤ
├── pkg/                  # 外部で利用可能なライブラリやパッケージ
├── scripts/              # スクリプト (データベースマイグレーション、セットアップスクリプトなど)
├── test/                 # 追加の外部テストアプリケーションとテストデータ
├── third_party/          # サードパーティのユーティリティ
├── web/                  # Webサーバと関連ファイル
├── .gitignore            # git ignoreファイル
├── Makefile              # ビルド、テスト、デプロイメントを自動化するためのMakefile
├── README.md             # プロジェクトのREADME
└── go.mod                # Goのモジュール依存関係
```
