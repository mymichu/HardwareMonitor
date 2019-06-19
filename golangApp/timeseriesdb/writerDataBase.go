package timeseriesdb

import (
	"errors"
	"fmt"
	"time"

	"../common"
	"github.com/influxdata/influxdb1-client/v2"
)

type InfluxConnection struct {
	Address      string
	Username     string
	Password     string
	DataBaseName string
	client       *client.Client
}

func (r *InfluxConnection) initializeInfluxConnection() error {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     r.Address,
		Username: r.Username,
		Password: r.Password,
	})
	fmt.Println("CONNECT " + r.Address)
	if err != nil {
		fmt.Println(err)
		r.client = nil
		return errors.New("Client not connected")
	}
	r.client = &c
	return nil
}

func (r *InfluxConnection) writeToInfluxDB(
	dataBase string,
	point string,
	tags map[string]string,
	fields map[string]interface{},
	timestamp time.Time,
) error {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  dataBase,
		Precision: "s",
	})

	if err != nil {
		return err
	}
	// Create a new point and add to batch
	pt, err := client.NewPoint(point, tags, fields, timestamp)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	err = (*r.client).Write(bp)

	return err

}

func (r *InfluxConnection) writeCPUStateEntry(
	devID string,
	cpuState common.CpuState,
) error {
	// Create a point and add to batch
	tags := map[string]string{"device_id": devID}
	fields := map[string]interface{}{
		"usage": cpuState.Usage,
		"temp":  cpuState.Temp,
	}
	return r.writeToInfluxDB(r.DataBaseName, "cpu_state", tags, fields, cpuState.Timestamp)
}
