package main

import (
	client "github.com/influxdata/influxdb/client/v2"
)

//InfluxClient ...
type InfluxClient struct {
	client.Client
}

// NewInfluxClient ...
func NewInfluxClient() ( *InfluxClient, error ) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
	})

	if err != nil {
		panic(err)
	}

	return &InfluxClient{c}, nil
}

func (c *InfluxClient) save(temp Temperature) error {

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "wurst",
		Precision: "s",
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *InfluxClient) get(id int) ( result []byte, err error ) {


	return result, nil
}
