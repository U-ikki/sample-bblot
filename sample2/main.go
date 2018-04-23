package main

import (
    "fmt"
    "log"
    "os"
    "net/http"

    "github.com/gorilla/mux"

    bolt "github.com/coreos/bbolt"

)

var (
    db      *bolt.DB
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Hello)
    // create db
    router.HandleFunc("/bboltdb/{dbname}", CreateDb)
    // select * from db
    router.HandleFunc("/bboltdb/{dbname}/{bucketname}/all", SelectAll)

    log.Fatal(http.ListenAndServe(":8080", router))

}

func Hello(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Hello bbolt!")
}

//--------------

func Open(w http.ResponseWriter, dbname string) {
    dbpath := "/tmp/"+ dbname

    if _, err := os.Stat(dbpath); os.IsNotExist(err) {
        fmt.Fprintln(w,"Database not found: ", dbpath)
    }

    var err error

    db, err = bolt.Open(dbpath, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }

}

func Close() {
    db.Close()
}

func SelectData(w http.ResponseWriter, bucketname string) {
    db.View(func(tx *bolt.Tx) error {

        b := tx.Bucket([]byte(bucketname))

        c := b.Cursor()

        for k, v := c.First(); k != nil; k, v = c.Next() {
            fmt.Fprintln(w, fmt.Sprintf("key=%s, value=%s", k, v))
        }

        return nil

    })
}

//------------------------

func CreateDb(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    dbname := vars["dbname"]
    Open(w, dbname)
    Close()
    fmt.Fprintln(w, "Database file created.")
}

func SelectAll(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    dbname := vars["dbname"]
    bucketname := vars["bucketname"]
    Open(w, dbname)
    fmt.Fprintln(w, "bucket :", bucketname)
    SelectData(w,bucketname)
    Close()
}



