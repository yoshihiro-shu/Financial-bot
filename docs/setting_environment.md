# 環境構築手順

## Docker環境の構築

1. Docker Networkの作成

    ```zsh
    docker network create financial-bot-network
    ```

2. Docker Composeのビルド

    ```zsh
    docker compose build
    docker compose -f compose-apps.yaml build
    ```

3. Docker Composeのスタート

    ```zsh
    make start
    ```

4. レスポンスチェック

    ```zsh
    curl http://127.0.0.1:80/api/v1/batch/health-check
    curl http://127.0.0.1:80/api/v1/notification/health-check
    ```

## Tool Install

1. ORM生成ツール

    [sqlc](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html)

    ```zsh
    brew install sqlc
    ```

    generate the orm file

    ```zsh
    make gen-orm
    ```

2. Migrationツール

    [goose](https://github.com/pressly/goose)

    ```zsh
    brew install goose
    ```

    generate the migration file

    ```zsh
    make gen-migration
    ```

3. go-zero

    generate template `.api`file

    ```zsh
    goctl api -o apps/some/main.api
    ```

    generate go file by `.api`file

    ```zsh
    goctl api go -api apps/some/main.api -dir apps/some
    ```
