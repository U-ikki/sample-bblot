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
    arg1 := flag.Arg(0)
    arg2 := flag.Arg(1)

    if _, err := os.Stat(*dbpath); os.IsNotExist(err) {
        fmt.Println(fmt.Sprintf("Database not found: %s", *dbpath))
    }

    var err error

    db, err = bolt.Open(*dbpath, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    upsertData(*bucketname, arg1, arg2)

}


func upsertData(bucketname string, arg1 string, arg2 string) {
    db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucketname))
        if b == nil {
            fmt.Println(fmt.Sprintf("bucket:%s is not found.",bucketname))
            return nil
        }

        v := b.Get([]byte(arg1))
        if v != nil {
            err := b.Put([]byte(arg1),[]byte(arg2))
            fmt.Println(fmt.Sprintf("key:%s was updated.", arg1))
            return err
        } else {
            err := b.Put([]byte(arg1),[]byte(arg2))
            fmt.Println(fmt.Sprintf("key:%s was inserted.", arg1))
            return err
        }
    })
}


