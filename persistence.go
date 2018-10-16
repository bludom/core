package main

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

//BoltClient ...
type BoltClient struct {
	*bolt.DB
}

// NewBoltClient ...
func NewBoltClient() (*BoltClient, error) {
	db, err := bolt.Open("temperature.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return &BoltClient{db}, nil
}

func (c *BoltClient) save(temp Temperature) error {

	if err := c.Update(func(tx *bolt.Tx) error {
		tempBucket, err := tx.CreateBucketIfNotExists([]byte("temperature"))
		if err != nil {
			return err
		}

		return tempBucket.Put([]byte(fmt.Sprintf("%d", temp.Core)), []byte(fmt.Sprintf("%f", temp.Temp)))

	}); err != nil {
		panic(err)
	}
	return nil
}

func (c *BoltClient) get(id int) (result []byte, err error) {

	if err := c.View(func(tx *bolt.Tx) error {
		dbresult := tx.Bucket([]byte("temperature")).Get([]byte(fmt.Sprintf("%d", id)))
		result = make([]byte, len(dbresult))
		copy(result, dbresult)
		return nil
	}); err != nil {
		panic(err)
	}

	return result, nil
}
