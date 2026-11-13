# コーディング規約

## エラー処理

- Controller層までエラーを戻したうえで、Controllerが適切にエラーを処理する
- Controller層以外で発生したエラーは`errors.Wrap`などでラップしたうえで呼び出し元に返す
- アプリケーションの初期化処理中にエラーが発生し、初期化を継続できない場合は`panic`で終了させる

## ログ出力

- `log/slog` を利用する
- ログ出力には `slog.Info` ではなく `slog.InfoContext` のようにContextを引数として要求するAPIを利用する。
  - `sloglint`で検査される
- エラーはController層で記録し、返却元では重複して記録しない
