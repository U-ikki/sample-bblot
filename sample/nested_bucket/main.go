package main

import (
    "log"
    "fmt"
    "os"
    "flag"
    "strconv"

    bolt "github.com/coreos/bbolt"
)

var (
    db    *bolt.DB
)

func main() {
    dbpath := flag.String("database", "/tmp/testdb", "Database path")
    username := flag.String("username", "none", "User name")

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

    createBucket(1234567)
    createNestedBucket(1234567, "project01")
    insertUserInfo(1234567, "project01", *username)
}


func createBucket(accountid uint64) {
    db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte(strconv.FormatUint(accountid, 10)))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        } else {
            fmt.Println(fmt.Sprintf("bucket: %s created.", accountid))
        }
        return nil
    })
}


func createNestedBucket(rootbucket uint64, subbucket string) error {
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    root := tx.Bucket([]byte(strconv.FormatUint(rootbucket, 10)))
    if root == nil {
        fmt.Println(fmt.Sprintf("accountid:%s is not found.",rootbucket))
        return nil
    }

    _, err = root.CreateBucketIfNotExists([]byte(subbucket))
    if err != nil {
        return err
    }

    if err = tx.Commit(); err != nil {
        return err
    }
    return nil
}

func insertUserInfo(rootbucket uint64, subbucket string, name string) error {
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    root := tx.Bucket([]byte(strconv.FormatUint(rootbucket, 10)))
    bkt := root.Bucket([]byte(subbucket))

    userID, err := bkt.NextSequence()
    if err != nil {
        return err
    }

    err = bkt.Put([]byte(strconv.FormatUint(userID, 10)), []byte(name) )
    if err != nil {
        return err
    }

    if err := tx.Commit(); err != nil {
        return err
    }
    return nil
}
