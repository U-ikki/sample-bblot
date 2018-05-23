package backend

import (
    "fmt"
    "os"
    "log"

    bolt "github.com/coreos/bbolt"
)

var (
    db    *bolt.DB
    dbpath string = "/tmp/test.db"
    bucketname string = "testbucket"
)

func OpenDb() (*bolt.DB, error){
    if _, err := os.Stat(dbpath); os.IsNotExist(err) {
        fmt.Println(fmt.Sprintf("Database not found: %s", dbpath))
    }
    var err error
    db, err = bolt.Open(dbpath, 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    return db, nil
}

func CreateBucket(db *bolt.DB) (error){
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }

    defer tx.Rollback()

    _, err = tx.CreateBucketIfNotExists([]byte(bucketname))
    if err != nil {
        return err
    }

    // Commit the transaction and check for error.
    if err := tx.Commit(); err != nil {
        return err
    }
    return nil
}


func PutCollection(db *bolt.DB, key string, interface{}) (error) {
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }

    defer tx.Rollback()

    b := tx.Bucket([]byte(bucketname))
    err := b.Put([]byte(key), []byte("42"))


}

