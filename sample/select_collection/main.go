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
    //arg2 := flag.Arg(1)

    if _, err := os.Stat(*dbpath); os.IsNotExist(err) {
        fmt.Println(fmt.Sprintf("Database not found: %s", *dbpath))
    }

    var err error

    db, err = bolt.Open(*dbpath, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    selectData(*bucketname, arg1)

}


func selectData(bucketname string, arg1 string) {

    db.Update(func(tx *bolt.Tx) error {

        b := tx.Bucket([]byte(bucketname))

        if b == nil {
            fmt.Println(fmt.Sprintf("bucket:%s is not found.",bucketname))
            return nil
        }

        v := b.Get([]byte(arg1))

        if v != nil {
            fmt.Println(fmt.Sprintf("key: %s is %s", arg1, v))
            return nil
        } else {
            fmt.Println(fmt.Sprintf("key:%s is not set value.", arg1))
            return nil
        }

    })

}

