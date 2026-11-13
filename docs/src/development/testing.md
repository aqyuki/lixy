# テストと品質チェック

## 単体テスト

```bash
# 外部依存の無い単体テストのみを実行する
go test -v -short -shuffle on -skip '^TestIntegrate' ./...

# 全ての単体テストを実行する
go test -v -shuffle on -skip '^TestIntegrate' ./...
```

## 結合テスト

```bash
# 全ての結合テストを実行する
go test -v -shuffle on -run '^TestIntegrate' ./...
```

## golangci-lint

Go のリンターとして[golangci-lint](https://golangci-lint.run/)を利用しています。

```bash
# 静的検査（フォーマットの検査も含む）
golangci-lint run ./...

# 自動修正
golangci-lint run --fix ./...

# フォーマッタの適用
golangci-lint fmt ./...
```

## oxfmt

jsonやYAMLファイルのフォーマッタとして[oxfmt](https://oxc.rs/docs/guide/usage/formatter.html)を利用しています。

```bash
# 検査
pnpm oxfmt --check ./

# 修正
pnpm oxfmt --write ./
```

## Lefthookによる検査

本プロジェクトではGit Hook管理に[Lefthook](http://lefthook.dev/)を利用しています。

### pre-commit hook での検査

コミット前に以下の項目が検査されます。

- golangci-lint によるフォーマットの検査
- oxfmt によるフォーマットの検査
- betterleaks によるシークレットの検査
