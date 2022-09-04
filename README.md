
コンテストを解き始める
```bash
make init https://atcoder.jp/contests/abc267
# 引数なしの場合、コンテストのURLを入力する
# make init
```

サンプルデータでテストをおこなう
```bash
make test
# 過去解いたコンテストを動かす場合、引数にディレクトリを渡す
# ./shell/test.sh contests/atcoder/abc267/a
```

解き終わったら、テストデータをmain.goのファイルをcontests配下に移動する
```bash
make save
```

問題を新しく解き始める
```bash
# B問題に進む場合
make new-b
```

今解いている問題を表示する
```bash
make current
```