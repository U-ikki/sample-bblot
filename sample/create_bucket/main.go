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
    //arg1 := flag.Args(0)
    //arg2 := flag.Args(1)

    if _, err := os.Stat(*dbpath); os.IsNotExist(err) {
        fmt.Println(fmt.Sprintf("Database not found: %s", *dbpath))
    }

    var err error

    db, err = bolt.Open(*dbpath, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    createBucket(*bucketname)

}


func createBucket(bucketname string) {
    db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte(bucketname))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        } else {
            fmt.Println(fmt.Sprintf("bucket: %s created.", bucketname))
        }
        return nil
    })
}



