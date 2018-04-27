package main

import (
    "log"
    "fmt"
    "os"
    "flag"

    bolt "github.com/coreos/bbolt"
)

var (
    db    *bolt.DB
)

func main() {
    dbpath := flag.String("database", "/tmp/testdb", "Database path")
    rootbucket := flag.String("rootbucket","1234567","Root Bucket name")
    subbucket := flag.String("subbucket","project01","Sub Bucket name")

    flag.Parse()

    if _, err := os.Stat(*dbpath); os.IsNotExist(err) {
        fmt.Println(fmt.Sprintf("Database not found: %s", *dbpath))
    }

    var err error

    db, err = bolt.Open(*dbpath, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    selectData(*rootbucket, *subbucket)

}


func selectData(rootbucket string, subbucket string) {

    db.View(func(tx *bolt.Tx) error {

        root := tx.Bucket([]byte(rootbucket))
        bkt := root.Bucket([]byte(subbucket))
        c := bkt.Cursor()

        fmt.Println(fmt.Sprintf("rootbucket: %s, subbucket: %s", rootbucket, subbucket))

        for k, v := c.First(); k != nil; k, v = c.Next() {
            fmt.Print(fmt.Sprintf("userID=%s, name=%s\n", k, v))
        }

        return nil

    })

}

