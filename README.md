# financial-bot

金融に関する情報をLine Botで定期的に送信するサービス

## Design Docs

[こちら](./docs/design_docs.md)

## System Design

[diagram](https://app.diagrams.net/?src=about#G1u6J3T5kn-cB-85117u00gpe_KJAJez7F)

## ニュース送信のフロー

<img width="432" alt="image" src="https://github.com/yoshihiro-shu/financial-bot/assets/84740493/7b2edc84-c5fc-4b24-9c29-79df19d40b97">

## 環境構築

1. Docker Networkの作成

    ```zsh
    docker network create financial-bot
    ```

2. Docker Composeのビルド

    ```zsh
    docker compose up --build
    ```

3. レスポンスチェック

    ```zsh
    curl http://127.0.0.1:8080/
    ```


## ディレクトリ構成

### 基本構造

```zsh
project-name/
├── api/                  # API定義とプロトコルファイル
├── assets/               # 画像、CSS、JavaScriptファイルなどのアセット
├── build/                # ビルドとデプロイのスクリプト
├── cmd/                  # プロジェクトのエントリーポイント (main.goファイル)
├── configs/              # 設定ファイル
├── deployments/          # KubernetesやDockerのデプロイメント設定
├── docs/                 # ドキュメント
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
├── Dockerfile            # Dockerfile
├── Makefile              # ビルド、テスト、デプロイメントを自動化するためのMakefile
├── README.md             # プロジェクトのREADME
└── go.mod                # Goのモジュール依存関係
```
