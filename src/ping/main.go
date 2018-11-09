package main

import (
    "log"
    "github.com/influxdata/influxdb/client/v2"
    "time"
    "ping/yaml"
    json "ping/JsonBodyMAC"
    ping "ping/RunMac"
)

func main() {
  // while loop
  for {
    // loop through target list
    for _, ip := range yaml.YamlReader().Targets {
      // go routine
      // fmt.Println(reflect.TypeOf(jm.JsonBody(runPing(ip), yaml.YamlReader().Region)))
      go influxdb(json.JsonBody(ping.RunPing(ip), yaml.YamlReader().Region))
      }
    time.Sleep(3 * time.Second)
    }
  }

func influxdb(r json.Body) {
  // Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
		Username: "root",
		Password: "supersecretpassword",
	})
  if err != nil {
    log.Fatal("Failed to create db client session: %v", err)
  }
  defer c.Close()

  // Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig {
		Database:  "network_telemetry",
		Precision: "s",
	})
	if err != nil {
		log.Fatal("Failed to create point batch: %v", err)
	}

  // Create a point and add to batch
	tags := map[string]string{"target": r.Target, "region": r.Region}
  fields := map[string]interface{}{
    "transmitted": r.Transmitted,
    "received": r.Received,
    "loss": r.Loss,
    "min": r.Min,
    "avg": r.Avg,
    "max": r.Max,
    "stddev": r.Stddev,
	}
	pt, err := client.NewPoint("ping_rtt_go", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	// Close client resources
	if err := c.Close(); err != nil {
    		log.Fatal(err)
	}
}
