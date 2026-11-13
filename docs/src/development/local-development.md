# ローカルでの開発

## アプリケーションの実行

```bash
go run ./cmd/server
```

### Dockerでの実行

Dockerfileのビルドには[BuildKit](https://docs.docker.com/build/buildkit/)が必要です。

```bash
docker buildx build --load -t lixy:local .
docker run --rm lixy:local
```

### Airによるホットリロード

本プロジェクトでは[air](https://github.com/air-verse/air)によるホットリロードが可能です。
プロジェクトルートで以下のコマンドを実行するとair経由でDiscord Botが起動され、コードを編集するとホットリロードされます。

::: warning
`config/env/<env>.yaml`が変更されてもairはホットリロードを行わないため、必要に応じてair自体を再起動してください。
:::

```bash
air
```
