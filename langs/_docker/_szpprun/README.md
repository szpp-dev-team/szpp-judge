# SZPP Run

## これは何？

以下の機能を持った、コマンド実行用のコマンドです。

- コマンドを時間・メモリ・ファイル書き出しサイズなどの制限付きでサブプロセスとして実行する
- 実行したコマンドの実行時間を計測する
- 実行したコマンドのメモリ使用量を計測する
- コマンド実行後、`.exec-result.txt` に {実行時間[ms]}, {メモリ使用量[kiB]}, {exitCode} を空白区切りで保存する (末尾の改行はナシ)

:warn: rlimit を用いている & /proc/{pid}/status を読む関係で、おそらく Linux でしか動作しません。

## 実行書式

```sh
szpprun timeLimit_ms memLimit_MiB fileSizeLimit_MiB numOpenFileLimit CMD...
```

例：
```sh
# 実行時間制限 1000 ms
# メモリ制限 256 MiB
# ファイル書き出し制限 1 MiB
# 同時ファイルオープン数上限 128
szpprun 1000 256 1 128 python3 main.py
```

## 注意

以下の項目は、szpprun 自身にも制約を課すことになるので注意してください。
理由は、以下の制約は szpprun のプロセス自身に rlimit をかけることで実現しているからです。
(rlimit 制約はサブプロセスにも受け継がれることを利用している)

- ファイル書き出し制限
- 同時ファイルオープン数
