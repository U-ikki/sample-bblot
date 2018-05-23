package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

func HelloWorld() echo.HandlerFunc {
    return func(c echo.Context) error {     //c をいじって Request, Responseを色々する 
        return c.String(http.StatusOK, "Hello World\n")
    }
}

func SamplePage() echo.HandlerFunc {
    return func(c echo.Context) error {
        username := c.Param("username")    //プレースホルダusernameの値取り出し
        return c.String(http.StatusOK, "Sample Page : " + username + "\n")
    }
}


