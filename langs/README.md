# SZPP Judge Languages

## コマンド

※ 詳細は Makefile を確認してください。

### langs.go のチェック

```sh
make lint
```

上記のコマンドを実行することで、例えば以下の項目をチェックできます
- LangID の重複チェック
- 言語名称の重複チェック
- 未設定のフィールドがないか
- 言語名称に余分な空白が含まれていないか

詳細は langs_test.go を確認してください。

### langs.go の整形

```sh
make fmt
```

### 各言語の Docker イメージのビルド

```sh
make docker/build/langs
```

### イメージの一覧 (SZPP Judge 用のものだけ表示)

```sh
make docker/image/ls
```


### 古いイメージの削除 (SZPP Judge 用のものだけ削除)

GC (ガーベジコレクション) します

```sh
make docker/image/gc
```
