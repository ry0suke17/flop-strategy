# server

サーバー側のコードを管理しています。

## 開発

API サーバーの立ち上げまでの流れは以下の通りです。

```bash
# dev/setup は初回の開発に必要なセットアップを行う
make dev/setup

# API サーバーを立ち上げる。
make dev/api
```

API の定義を更新した時は以下のコマンドでコードを生成してあげる必要があります。

```bash
# API コードを生成する。
make dev/openapi/generate
```
