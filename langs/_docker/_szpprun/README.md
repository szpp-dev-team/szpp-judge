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
szpprun TIME_LIMIT_ms MEM_LIMIT_MiB CMD...
```

例：
```sh
# 1000 ms , 256 MiB の制限つきで実行する場合:
szpprun 1000 256 python3 main.py
```

内部で rlimit による制限をしていますが、それは実行時には設定できません。
プログラム内で定数として設定しています。