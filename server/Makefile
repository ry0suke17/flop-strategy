# Commands for development {
# ------------------------------------------------------------------
# ------------------------------------------------------------------

## dev/fmt は強制的に適切なコードに書き変える。
dev/fmt:
	# goimports で書き換える。
	goimports -w `find . -name vendor -prune -type f -o -name '*.go'`
	# gofmt のコードフォーマットに書き換えた上で、よりシンプルに書けるコードを書き換える。
	gofmt -s -w `find . -name vendor -prune -type f -o -name '*.go'`
	# 不要な型変換を取り除く。
	unconvert -apply ./...

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

## DEV_POSTGRES_DATABASE_URL は開発時に起動する PostgreSQL データベース URL を表す。
DEV_POSTGRES_DATABASE_URL=postgres://postgres:docker@localhost:5432/flop-strategy?sslmode=disable

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

## dev/api は API サーバーを立ち上げる。
dev/api:
	go run cmd/server/api/main.go \
		-postgres_database_url "$(DEV_POSTGRES_DATABASE_URL)"

## dev/openapi/generate は API コードを生成する。
dev/openapi/generate:
	npx @openapitools/openapi-generator-cli generate \
		-i ../api/flopStrategy.yaml \
		-g go-server \
		-o ./genopenapi \
		--additional-properties=featureCORS=true,packageName=api,sourceFolder=api \
		--global-property=models,supportingFiles
	rm -rf ./transport/openapi/api
	cp -r ./genopenapi/api ./transport/openapi/
	rm -rf ./genopenapi
	make dev/fmt

# ------------------------------------------------------------------
# ------------------------------------------------------------------
# }

# テストで利用するコマンド {
# ------------------------------------------------------------------
# ------------------------------------------------------------------

## test/lint は linter を実行する。
## TODO: folint などで openapi をとりあえず除外してるのでもっといい方法で除外してあげた方がいい。（エディタ側の golint は除外されていない）
test/lint:
	# コーディングスタイルの問題をチェックする
	(! gofmt -s -d `find . -name vendor -prune -type f -o -name '*.go'` | grep '^')
	# コーディングスタイルの問題をチェックする
	golint -set_exit_status `go list ./... | grep -v transport/openapi`
	# import 文のフォーマットをチェックする
	(! goimports -l `find . -name vendor -prune -type f -o -name '*.go'` | grep 'go')
	# コンパイラが検出しないエラーをチェックする。
	go vet ./...
	# shadowed 変数をチェックする。
	go vet -vettool=$(which shadow) ./...
	# 関数のエラー戻り値をハンドリングしているかをチェックする。
	errcheck `go list ./... | grep -v transport/openapi`
	# 未使用のグローバル変数と定数がないかチェックする。
	varcheck ./...
	# 代入した値を無視していないかチェックする
	ineffassign .
	# 不要な型変換をしていないかチェックする。
	unconvert -v ./...

# ------------------------------------------------------------------
# ------------------------------------------------------------------
# }
