package main

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
)

func save(temp Temperature) error {

	db, err := bolt.Open("temperature.db", 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := db.Update(func(tx *bolt.Tx) error{
		tempBucket, err := tx.CreateBucketIfNotExists([]byte("temperature"))
		if err != nil {
			return err
		}

		return tempBucket.Put([]byte(fmt.Sprintf("%d", temp.ID)), []byte(fmt.Sprintf("%f", temp.Temperature)))

	}); err != nil {
		panic(err)
	}
	return nil
}

func get(id int) ( result []byte, err error ) {

	db, err := bolt.Open("temperature.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	defer db.Close()
	if err := db.View(func(tx *bolt.Tx) error{
		dbresult := tx.Bucket([]byte("temperature")).Get([]byte(fmt.Sprintf("%d", id)))
		result = make([]byte, len(dbresult))
		copy(result, dbresult)
		return nil
	}); err != nil {
		panic(err)
	}

	return result, nil
}
