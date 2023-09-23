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

    generate the orm file

    ```zsh
    sqlc generate
    ```

2. Migrationツール

    [goose](https://github.com/pressly/goose)

    ```zsh
    brew install goose
    ```

    generate the migration file

    ```zsh
    goose --dir ./db/migrations create add_some_column sql
    Created new file: 20170506082420_add_some_column.sql
    ```

3. go-zero

    generate template `.api`file

    ```zsh
    goctl api -o internal/server/api/some.api
    ```

    generate go file by `.api`file

    ```zsh
    goctl api go -api internal/server/api/main.api -dir internal/server
    ```
