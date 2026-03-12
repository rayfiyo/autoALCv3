# autoALCv3

- ALC を自動でやる 3 版
- This is the 3rd edition of ALC with automation
- 依存関係: Google Chrome

# Download

事前準備として、`Google Chrome` をインストールしておく。

## Windows

[Release](https://github.com/rayfiyo/autoALCv3/releases/latest) から、
対応するアーキテクチャのバイナリをダウンロードする。
（`autoALCv3_windows_アーキテクチャ.exe`）

## MacOS

MacOS 10.12 Sierra 以降では
[Release](https://github.com/rayfiyo/autoALCv3/releases/latest) から、
対応するアーキテクチャのバイナリをダウンロードする。
（`autoALCv3_darwin_アーキテクチャ`）

## Linux

[Release](https://github.com/rayfiyo/autoALCv3/releases/latest) から、
対応するアーキテクチャのバイナリをダウンロードする。
（`autoALCv3_linux_アーキテクチャ`）

## その他/動かない場合

このレポジトリをクローン後 `go build main.go` を行い、生成されたバイナリを実行する。

# Usage

## ログイン

### ID

- s を付けた 学籍番号 を入力する
- ex: s54321

### Password

- **ALC の**ログインに使うパスワードを入力する
- `-`でマスクされる
- ローカルなどに保存はされない
- 統合認証とは異なる可能性あり

### コース

- 表示に従う
- PWH は PowerWords Hybrid コース のこと
- TC1 は TOEIC(R) L&R テスト 500 点突破コース のこと

### サブコース

- PWH では LEVEL01 などのこと
- TC1 では Stage 1 「狙い目」攻略 などのこと
- 上から順に数字を振っている

# About skill point

- ユニットを解くとスキルポイントが付与される
- ポイントは `L,S,R,W,G,V` の６種類（６項目）がある
- 現在以下のように実装しているが このポイント付与は値や種類が正しくない場合がある
- それ以外の場合 と判定されたときのみ標準出力に その旨が出力される

## ポイントの実装

- PWH の場合: `V` に 10 ポイント
- TC 系（TOEIC）: `L` に 10 ポイント
- JT（実力テスト）: `V` に 30 ポイント
- KT（確認テスト）: 0 ポイント // 本来はスキルポイントのペイロードがない？
- それ以外の場合: 0 ポイント
