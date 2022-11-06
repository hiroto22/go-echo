package main

import (
	"log"

	"github.com/hiroto22/go-echo/handler"
	"github.com/hiroto22/go-echo/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := store.New()
	if err != nil {
		log.Fatalln("db connection error")
	}

	// ルートを設定
	e.GET("/tasks", handler.GetTask(db))
	e.POST("/tasks", handler.AddTask(db))

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}
