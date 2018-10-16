package main

import (
	"fmt"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

//InfluxClient ...
type InfluxClient struct {
	client.Client
}

// NewInfluxClient ...
func NewInfluxClient() (*InfluxClient, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://gomano.de:8086/write?db=talk",
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
	tags := map[string]string{
		"host":        temp.Device,
		"core":        strconv.Itoa(temp.Core),
	}

	fields := map[string]interface{}{
		"value": temp.Temp,
	}

	pt, err := client.NewPoint("temperatur", tags, fields, time.Now())
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

func (c *InfluxClient) get(id int) (result []byte, err error) {

	res, err := queryDB(c, fmt.Sprintf("SELECT \"temperature\"::field FROM \"test\" WHERE \"id\" = %d", id))
	if err != nil {
		return nil, err
	}

	result = []byte(fmt.Sprintf("%+v", res[0]))

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
