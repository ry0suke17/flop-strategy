# Commands for development {
# ------------------------------------------------------------------
# ------------------------------------------------------------------

## dev/setup は初回の開発に必要なセットアップを行う
dev/setup: \
	dev/compose \
	dev/postgres/wait \
	dev/postgres/setup

## dev/compose はサービスの docker を立ち上げる
dev/compose:
	docker-compose up -d

## dev/postgres/setup は Postgres のセットアップを行う。
dev/postgres/setup: \
	dev/postgres/create/schema \
	dev/postgres/seeding

## DEV_POSTGRES_ROOT_PASSWORD は Postgres root のパスワードを表す。
DEV_POSTGRES_ROOT_PASSWORD=docker

## DEV_POSTGRES_PORT は Postgres ポート番号を表す。
DEV_POSTGRES_PORT=5432

## DEV_POSTGRES_DATABASE は Postgres データベース名を表す。
DEV_POSTGRES_DATABASE=flop-strategy

## dev/postgres/wait は Postgres サーバーの起動を待つ。
dev/postgres/wait:
	sleep 10

## dev/postgres/create/schema は Postgres データベーススキーマを作成する。
dev/postgres/create/schema:
	PGPASSWORD=$(DEV_POSTGRES_ROOT_PASSWORD) createdb $(DEV_POSTGRES_DATABASE) -U postgres -h localhost -p $(DEV_POSTGRES_PORT)
	PGPASSWORD=$(DEV_POSTGRES_ROOT_PASSWORD) psql -d $(DEV_POSTGRES_DATABASE) -U postgres -h localhost -p $(DEV_POSTGRES_PORT) -f _sql/schema.sql

## dev/postgres/recreate/schema は Postgres データベーススキーマを再作成する。
dev/postgres/recreate/schema:
	PGPASSWORD=$(DEV_POSTGRES_ROOT_PASSWORD) dropdb $(DEV_POSTGRES_DATABASE) -U postgres -h localhost -p $(DEV_POSTGRES_PORT) --if-exists
	make dev/postgres/create/schema

## dev/postgres/seeding は Postgres データベースに初期データを挿入する。
dev/postgres/seeding:
	PGPASSWORD=$(DEV_POSTGRES_ROOT_PASSWORD) psql -d $(DEV_POSTGRES_DATABASE) -U postgres -h localhost -p $(DEV_POSTGRES_PORT) -f _sql/boards.sql
	PGPASSWORD=$(DEV_POSTGRES_ROOT_PASSWORD) psql -d $(DEV_POSTGRES_DATABASE) -U postgres -h localhost -p $(DEV_POSTGRES_PORT) -f _sql/heads_up_situations.sql
	PGPASSWORD=$(DEV_POSTGRES_ROOT_PASSWORD) psql -d $(DEV_POSTGRES_DATABASE) -U postgres -h localhost -p $(DEV_POSTGRES_PORT) -f _sql/flop_situations.sql

# ------------------------------------------------------------------
# ------------------------------------------------------------------
# }