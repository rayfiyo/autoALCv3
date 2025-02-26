# autoALCv3

- ALC を自動化でやる 第３版 / This is the 3rd edition of ALC with automation
- サーバーに負荷をかけないために，遅延大きめで実装
- Google Chrome を利用している（依存している）ので，Google Chrome のインストールが事前に必要

# Download

- 事前準備として、`Google Chrome` をインストールしておく

## Windows

次のどちらかの方法がある

1. [こちらの Release](https://github.com/rayfiyo/autoALCv3/releases/latest)から，一致するアーキテクチャのバイナリをダウンロードする（autoALCv3_windows_アーキテクチャ.exe）
2. `bin` のディレクトリの，一致するアーキテクチャのバイナリをダウンロードする

## MacOS

MacOS 10.12 Sierra 以降では，次のどちらかの方法がある
1. [こちらの Release](https://github.com/rayfiyo/autoALCv3/releases/latest)から，一致するアーキテクチャのバイナリをダウンロードする（autoALCv3_mac_アーキテクチャ.command）
2. `bin` のディレクトリの，一致するアーキテクチャのバイナリをダウンロードする

## Linux

次のどちらかの方法がある

1. [こちらの Release](https://github.com/rayfiyo/autoALCv3/releases/latest)から，一致するアーキテクチャのバイナリをダウンロードする（autoALCv3_linux_アーキテクチャ）
2. `bin` のディレクトリの，一致するアーキテクチャのバイナリをダウンロードする

## その他/動かない場合

- このレポジトリをクローンし，`go build main.go`を行い，生成されたバイナリを実行する

# Usage

## ログイン

### ID

- s を付けた 学籍番号 を入力する
- ex: s54321

### Password

- **ALC の**ログインに使うパスワードを入力する
- - でマスクされる
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
- ポイントは，L,S,R,W,G,V の６種類（６項目）がある
- 現在以下のように実装しているが，このポイント付与は 値や種類が正しくない場合がある
- それ以外の場合 と判定されたときのみ，標準出力に その旨が出力される

## ポイントの実装

- PWH の場合: V に 10 ポイント
- TC 系（TOEIC）: L に 10 ポイント
- JT（実力テスト）: V に 30 ポイント
- KT（確認テスト）: 0 ポイント // 本来はスキルポイントのペイロードがない？
- それ以外の場合: 0 ポイント

## 対処

- 現在開発開始予定のポイント調整プログラムを待つか，自分で対処する
