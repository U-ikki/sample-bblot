package main

import(
    //"fmt"
    "log"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "sample-bblot/sample3/handler"
    "sample-bblot/sample3/backend"
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

    err = backend.CreateBucket(db)
    if err != nil {
        log.Fatal("open error.")
    }

    // ルーティング
    e.GET("/hello", handler.HelloWorld())
    e.GET("/hello/:username", handler.SamplePage())


    //e.POST("/db/:key", handler.PutHandlerfunc(db))
    // curl -X POST -d "{postgresql:sample}"

    //e.GET("/db/:key", GetCollection())

    // サーバー起動
    e.Start("localhost:1323")
}

