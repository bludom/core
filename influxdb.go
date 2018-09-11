package main

import (
	client "github.com/influxdata/influxdb/client/v2"
	"time"
	"fmt"
	"strconv"
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
		Database:  "test",
		Precision: "s",
	})
	if err != nil {
		panic(err)
	}

		// Create a point and add to batch
	tags := map[string]string{"measurement": "temp", "id": strconv.Itoa(temp.ID)}
	fields := map[string]interface{}{
		"temperature":   temp.Temperature,
	}

	pt, err := client.NewPoint("temperature", tags, fields, time.Now())
	if err != nil {
		return err
	}

	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		return err
	}
	
	return nil
}

func (c *InfluxClient) get(id int) ( result []byte, err error ) {

	res, err := queryDB(c, fmt.Sprintf("SELECT \"id\"::%d",id))
	if err != nil {
		return nil, err
	}

	result = []byte(fmt.Sprintf("%v", res))

	return result, nil
}

func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "test",
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
