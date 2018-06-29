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

// PutCollection レコード登録
func PutCollection(db *bolt.DB, key string, value string) (error) {
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }

    b := tx.Bucket([]byte(bucketname))

    if err := b.Put([]byte(key), []byte(value)); err != nil {
        return err
    }

    _ = tx.Commit()

    return nil
}

// GetCollection レコード取得
func GetCollection(db *bolt.DB, key string)  *string {
    tx, err := db.Begin(false)
    if err != nil {
        return nil
    }

    defer tx.Rollback()

    b := tx.Bucket([]byte(bucketname))
    value := b.Get([]byte(key))
    v := string(value)
    return &v
}
