package main

import (
	"github.com/aocm/vue-go-spa-sample/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/aocm/vue-go-spa-sample/database"
)

func main() {

	db := database.Connect()
    defer db.Close()

    err := db.Ping()

    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }
	
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.POST("/yamabiko", handler.YamabikoAPI())
	e.OPTIONS("/yamabiko", handler.OptionsCheck())

	// サーバー起動
	e.Start(":8000")
}
