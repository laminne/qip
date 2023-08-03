# Qip

## 設計

依存性逆転の法則(DIP)を利用したレイヤードアーキテクチャ(もどき)を採用しています。

## ディレクトリ構成

```
- /
 - main.go # プログラム全体のエントリーポイント
 - /cmd    # サーバーの起動などのスクリプト群
 - /pkg
    - /activitypub  # ActivityPub関連
    - /application  # CRUD系処理 (Application Service)
    - /controller   # コントローラー
        - /models   # APIのRequest/Responseの型
    - /domain       # ドメインモデル
        - /service  # Domain Service
    - /errorType    # システム全体で使うユーザー定義エラー型
    - /repository   # Repository
    - /server       # HTTPサーバー関連
        - /handler  # ハンドラ関数群
        - /router   # ルーティング関数群
    - /utils
        - /config   # 設定読み込み関連
        - /id       # id生成関連
        - /key      # RSA鍵生成関連
        - /logger   # ロガー(Zapのラッパー)
        - /password # パスワードハッシュ関連
        - /token    # アクセストークン関連
```
