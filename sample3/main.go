package main

import (
	//"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"sample-bblot/sample3/backend"
	"sample-bblot/sample3/handler"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := backend.OpenDb()
	if err != nil {
		log.Fatal("open error.")
	}
	defer db.Close()

	if err = backend.CreateBucket(db);  err != nil {
		log.Fatal("open error.")
	}

	// ルーティング
	e.GET("/hello", handler.HelloWorld())
	e.GET("/hello/:username", handler.SamplePage())


    // curl -X GET http://localhost:1323/db/:key
    e.GET("/db/:key", handler.GetCollectionService(db))
	// curl -X PUT http://localhost:1323/db/:key -H 'Content-Type: application/json' -d '{"profile":"sample_data"}'
	e.PUT("/db/:key", handler.PutCollectionService(db))


	// サーバー起動
	e.Start("localhost:1323")
}
