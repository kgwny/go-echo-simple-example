# go-echo-simple-example

## 最初にやること

go.mod ファイルを初期化する
```
go mod init go-echo-simple-example
```

echo フレームワークのインストール
```
go get -u github.com/labstack/echo/v4
```

gorm のインストール
```
go get -u gorm.io/gorm
```

gorm sqlite ドライバのインストール
```
go get -u gorm.io/driver/sqlite
```

config ライブラリのインストール
```
go get gopkg.in/ini.v1
```

## サーバー起動
```
go run main.go
```

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

`http://localhost:8080` にアクセスすると `Hello, Echo!` が表示される  
`http://localhost:8080/list` にアクセスするとTODOリスト作成画面が表示される
