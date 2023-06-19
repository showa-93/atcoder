# atcoder

学習用リポジトリ  

コンテストを解き始める

```bash
make init https://atcoder.jp/contests/abc267
# 引数なしの場合、コンテストのURLを入力する
# make init or make
```

サンプルデータでテストをおこなう

```bash
make test
# 過去解いたコンテストを動かす場合、引数にディレクトリを渡す
# ./shell/test.sh contests/atcoder/abc267/a
```

現在解いている問題の回答を提出する

```bash
make solve
```

解き終わったら、テストデータをmain.goのファイルをcontests配下に移動し、コミットする

```bash
make save
```

問題を新しく解き始める

```bash
# B問題に進む場合
make new-b
```

今解いている問題をブラウザで表示する

```bash
make open
```

TODO:

- [x] template/main_test.go(.tmpl)をつくる。go generateしたい。
- [ ] ランダムテストの生成
  - [x] ベース作成
  - [x] int対応
  - [ ] float対応
  - [ ] string対応
    - [ ] Fuzzing対応
    - [ ] 複数文字列パターン対応
- [x] modのライブラリを改修する
- [ ] ベクトル系ライブラリ。scrapboxにベースはあるはず
- [ ] suffix arrayを実装しなおす
