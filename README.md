# go-echo-simple-example

## 最初にやること

自身のGitHubリポジトリに合わせて初期化する
```
go mod init github.com/kgwny/go-echo-simple-example
```

echo フレームワークのインストール
```
go get -u github.com/labstack/echo/v4
```

## Hello, Echo! を実行する

以下のコマンドを実行する
```
go run main.go
```

echo サーバーが起動する
```
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.13.4
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```

ブラウザで `localhost:8080` にアクセスすると `Hello, Echo!` が表示される
