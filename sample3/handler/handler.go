package handler

import (
    "net/http"
    //"encoding/json"

    "sample-bblot/sample3/backend"

    "github.com/labstack/echo"
    bolt "github.com/coreos/bbolt"
)

type Profile struct {
    TestProfile string `json:"profile"`
    TestConf string `json:"conf"`
}

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

func GetCollectionService(db *bolt.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        collectionKey := c.Param("key")

        v := backend.GetCollection(db, collectionKey)

        return c.String(http.StatusOK, *v)
    }
}

func PutCollectionService(db *bolt.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        p := new(Profile)
        if err := c.Bind(p); err != nil {
            return err
        }

        return c.JSON(http.StatusOK, p)

        /*
        m, err := json.Marshal(p)
        if err != nil {
            return err
        }

        jsonString := string(m)
        return c.JSON(http.StatusOK, jsonString)
        */
    }
}
