# Lixy - Docs

## 概要

- 本プロジェクトでは[vitepress](https://vitepress.dev/)を用いて資料を管理し、GitHub Pagesに公開しています。
- 一般ユーザー向けのドキュメントは[src/guid](./src/guide)に配置し、開発者向けのドキュメントは[src/development](./src/development/)に配置しています。

## ドキュメントの編集

`src/`配下に配置したファイルがvitepressによる公開の対象となります。

```bash
# devサーバーの起動
pnpm run docs:dev

# ビルド
pnpm run docs:build

# プレビュー
pnpm run docs:preview

# フォーマット
pnpm oxfmt --write ./
```

## カスタムアンカー

見出しの前に、次のコメントを記述します。

```md
<!-- anchor: getting-started -->

## はじめ方
```

この見出しには `#getting-started` でリンクできます。
見出しのリンクと目次にも同じアンカーが使われます。

- コメントと見出しの間には空行を入れられます。
- アンカー名には空白や `<`、`>` を含めず、ページ内で一意にしてください。
- 同じ見出しで `{#id}` と併用しないでください。
- コードブロック内のコメントや、見出しとの間に本文などがあるコメントは処理されません。
