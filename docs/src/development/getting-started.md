# Getting Started

## Prerequisite

本プロジェクトで開発を行うには以下のツールを事前に導入しておく必要があります。以下に記載の無いツールに関してはmiseやpnpm経由でインストールされます。

- [Go](https://go.dev/)（`go.mod` に指定されたバージョン。現在は 1.27.4）
- [mise](https://mise.jdx.dev/)
- コンテナの実行環境（[Docker](https://www.docker.com) / [OrbStack](https://orbstack.dev/) / [colima](https://colima.run/)など）
  - Dockerfileのビルドには別途BuildKitが必要なため、必要に応じて導入してください。

また、本プロジェクトではサプライチェーン攻撃対策として[AikidoSec/safe-chain](https://github.com/AikidoSec/safe-chain)の導入を推奨しています。

## Quick start

1. ローカルにリポジトリをクローンします。

   ```bash
   git clone https://github.com/aqyuki/lixy.git
   ```

1. ディレクトリに移動した後、開発に必要なツール類をインストールします。

   ```bash
   mise install
   ```

   `mise install`ではmiseで管理されているツールがインストールされます。また、`mise install`終了後に自動で`pnpm install` および `lefthook install` が実行されます。

1. テストおよびビルドが出来ることを確認する

   ```bash
   go test ./...
   go build ./cmd/server
   ```

以降は[ローカルでの開発](./local-development.md)を参照
