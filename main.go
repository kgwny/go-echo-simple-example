package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// インスタンス生成
	e := echo.New()

	// ルートハンドラの作成
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// サーバーの起動
	e.Logger.Fatal(e.Start(":8080"))
}
