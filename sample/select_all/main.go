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
    bucketname := flag.String("bucket","testBucket","Bucket name")

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

    selectData(*bucketname)

}


func selectData(bucketname string) {

    db.View(func(tx *bolt.Tx) error {

        b := tx.Bucket([]byte(bucketname))

        c := b.Cursor()

        for k, v := c.First(); k != nil; k, v = c.Next() {
            fmt.Print(fmt.Sprintf("key=%s, value=%s\n", k, v))
        }

        return nil

    })

}

