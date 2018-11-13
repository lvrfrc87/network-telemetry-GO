package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
	json "ping/JsonBodyMAC"
	ping "ping/RunMac"
	"ping/yaml"
	"time"
  "fmt"
)

func main() {
	for {
		for _, ip := range yaml.YamlReader().Targets {
			// go routine
			go influxdb(json.JsonBody(ping.RunPing(ip), yaml.YamlReader().Region))
		}
	}
}

func influxdb(r json.Body) {
	// Create a new HTTPClient
	dbList := []string{"app1.net.awsieprod2.linsys.tmcs", "db1.telemetry.netams1.netsys.tmcs",}
	for _, db := range dbList {
		// go routine
		go writeDb(db, r)
  }
}

func writeDb(db string, r json.Body) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:8086", db),
		Username: "root",
		Password: "supersecretpassword",
	})
	if err != nil {
		log.Fatalf("Failed to create %s client session: %v", db, err)
	}
	defer c.Close()
	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "network_telemetry",
		Precision: "s",
	})
	// Create a point and add to batch
	tags := map[string]string{"target": r.Target, "region": r.Region}
	fields := map[string]interface{}{
		"transmitted": r.Transmitted,
		"received":    r.Received,
		"loss":        r.Loss,
		"min":         r.Min,
		"avg":         r.Avg,
		"max":         r.Max,
		"stddev":      r.Stddev,
	}
	pt, _ := client.NewPoint("ping_rtt_go", tags, fields, time.Now())
	bp.AddPoint(pt)
	// Write the batch
	c.Write(bp)
	// Close client resources
	c.Close();
	return
}
