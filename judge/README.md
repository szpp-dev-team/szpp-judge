# Judge Server (for SZPP Judge)

## Dev commands

:bulb: 詳細は [Makefile](./Makefile) を確認してください。

###  `make lint`
ソースコードの静的解析やフォーマットチェックをします。

実行には `golangci-lint` を使うので以下のコマンド等でインストールしてください。

```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### `make fmt`
ソースコードをフォーマットします。


### `make test`
テストを実行します (テストケースのシャッフル、実行時の競合チェック付き)。
