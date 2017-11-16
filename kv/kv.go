/*
Package kv implements a library to provide CRUDL operations on an internal
Bolt KVDB.
A service to execute the same operations from CLI is also implemented.
*/
package kv

import (
	"fmt"
	"log"
	"github.com/boltdb/bolt"
)

var db *bolt.DB

// Create and open a Bolt database at the given path
func Open(path string) {
	var err error

	db, err = bolt.Open(path, 0600, nil)
	if err != nil {
		log.Printf("%s", err)
		log.Fatal("Failed to open internal KVDB")
	}
}

// Release all database resources.
func Close() {
	db.Close()
}

// Retrieve the value for a key in the bucket
func Get(bucket, key string) (string, error) {
	var res []byte

	// Execute a function within a managed read-only transaction
	err := db.View(func(tx *bolt.Tx) error {
    	// Retrieve bucket by name
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("Bucket '%s' does not exist", bucket)
		}
		// Retrieve value from bucket
    	res = b.Get([]byte(key))
		return nil
	})

    if err != nil {
        log.Printf("%s", err)
        return string(res), fmt.Errorf("Failed to lookup entry '%s'", key)
    }

	// Cast value to string
	return string(res), nil
}

// Retrieve all values in the bucket
func List(bucket string) (map[string]string, error) {
  var res map[string]string
	res = make(map[string]string)

	// Execute a function within a managed read-only transaction
	err := db.View(func(tx *bolt.Tx) error {
		// Retrieve bucket by name
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("Bucket '%s' does not exist", bucket)
		}

    // Iterate over items in sorted key order
    if err := b.ForEach(func(k, v []byte) error {
    	// Cast value to string
			res[string(k)] = string(v)
			return nil
    }); err != nil {
    	return err
    }
		return nil
	})

	if err != nil {
		log.Printf("%s", err)
		return res, fmt.Errorf("Failed to lookup entries '%s'", bucket)
	}

    return res, nil
}

// Set the value for a key in the bucket
func Insert(bucket, key, value string) error {
	// Execute a function within a read-write managed transaction
    err := db.Update(func(tx *bolt.Tx) error {
		// Create and/or retrieve a new bucket if it doesn't already exist
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		// Insert key-value pair into bucket
    	return b.Put([]byte(key), []byte(value))
	})

    if err != nil {
        log.Printf("%s", err)
        return fmt.Errorf("Failed to insert entry '%s'", key)
    }

    return nil
}

// Remove a key from the bucket
func Remove(bucket, key string) error {
	// Execute a function within a read-write managed transaction
    err := db.Update(func(tx *bolt.Tx) error {
		// Retrieve bucket by name
    	b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("Bucket '%s' does not exist", bucket)
		}
		// Delete key from bucket
    	return b.Delete([]byte(key))
	})

    if err != nil {
        log.Printf("%s", err)
        return fmt.Errorf("Failed to remove entry '%s'", key)
    }

    return nil
}

// Delete a bucket at the given key
func RemoveBucket(bucket string) error {
	// Execute a function within a read-write managed transaction
    err := db.Update(func(tx *bolt.Tx) error {
		// Drop bucket
    	return tx.DeleteBucket([]byte(bucket))
	})

    if err != nil {
        log.Printf("%s", err)
        return fmt.Errorf("Failed to remove bucket '%s'", bucket)
    }

    return nil
}
