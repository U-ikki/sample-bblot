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

    if _, err := os.Stat(*dbpath); os.IsNotExist(err) {
        fmt.Println(fmt.Sprintf("Database not found: %s", *dbpath))
    }

    var err error

    db, err = bolt.Open(*dbpath, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    deleteData(*bucketname, arg1)

}


func deleteData(bucketname string, arg1 string) {
    db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucketname))
        if b == nil {
            fmt.Println(fmt.Sprintf("bucket:%s does not found.",bucketname))
            return nil
        }
        v := b.Delete([]byte(arg1))
        if v == nil {
            fmt.Println(fmt.Sprintf("key: %s was deleted.", arg1))
            return nil
        } else {
            fmt.Println(fmt.Sprintf("key:%s does not exist.", arg1))
            return nil
        }
    })
}


