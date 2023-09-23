# 環境構築手順

## Docker環境の構築

1. Docker Networkの作成

    ```zsh
    docker network create financial-bot-network
    ```

2. Docker Composeのビルド

    ```zsh
    docker compose up --build
    ```

3. レスポンスチェック

    ```zsh
    curl http://127.0.0.1:8080/
    ```

## Tool Install

1. ORM生成ツール

    [sqlc](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html)

    ```zsh
    brew install sqlc
    ```

2. Migrationツール

    [goose](https://github.com/pressly/goose)

    ```zsh
    brew install goose
    ```

    migration fileの生成方法

    ```zsh
    goose create ./db/migrations/add_some_column sql
    Created new file: 20170506082420_add_some_column.sql
    ```
