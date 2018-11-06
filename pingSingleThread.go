/*
https://zaiste.net/executing_commands_via_ssh_using_go/
https://godoc.org/golang.org/x/crypto/ssh#example-PublicKeys
*/
package main

import (
    "bytes"
    "fmt"
    "golang.org/x/crypto/ssh"
    "io/ioutil"
    "log"
    "github.com/influxdata/influxdb/client/v2"
    "strings"
    "regexp"
)

type result struct {
    target string
    output []string
}

type body struct {
  target string
  region string
  transmitted string
  received string
  loss string
  min string
  avg string
  max string
}

func main() {
  // define variables
  cmd := "ping -c 1 10.78.65.1"
  port := "22"
  user := "core"
  hosts := "app1.net.awsieprod2.linsys.tmcs"

  // read the private key
  key, err := ioutil.ReadFile("/Users/federicoolivieri/.ssh/id_rsa")
  if err != nil {
    log.Fatalf("unable to read private key: %v", err)
    }
  // parse the prive key and create a ssh caller
  singer, err := ssh.ParsePrivateKey(key)
  if err != nil {
      log.Fatalf("unable to parse private key: %v", err)
  }
  // set-up ssh connection
  config := &ssh.ClientConfig{
    User: user,
    Auth: []ssh.AuthMethod{
      ssh.PublicKeys(singer),},
    // var hostKey ssh.PublicKey
    // HostKeyCallback: ssh.FixedHostKey(hostKey)
    HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
  execCmd := runPing(cmd, port, hosts, config)
  returned := jsonBody(execCmd.output)
  fmt.Println(
    "target", returned.target,
    "region", returned.region,
    "transmitted", returned.transmitted,
    "received", returned.received,
    "loss", returned.loss,
    "min", returned.min,
    "avg", returned.avg,
    "max", returned.max,
  )
  influxdb(jsonBody(execCmd.output))
}

func runPing(command, port, hostname string, config *ssh.ClientConfig) result {
  client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
  if err != nil {
      log.Fatalf("unable to connect: %v", err)
  }
  session, err := client.NewSession()
  if err != nil {
      log.Fatal("Failed to create session: %v", err)
  }
  defer session.Close()

  var stdoutBuf bytes.Buffer
  session.Stdout = &stdoutBuf
  session.Run(command)

  return result {
    target: hostname,
    output: strings.Split(stdoutBuf.String(), " "),
  }
}

func jsonBody(splittedValues []string) body {
  re := regexp.MustCompile(`\d+\.?\d?`)
  rttValues := strings.Split(splittedValues[29],"/")
  return body {
    target: splittedValues[1],
    region: "foo",
    transmitted: re.FindString(splittedValues[17]),
    received: re.FindString(splittedValues[20]),
    loss: re.FindString(splittedValues[22]),
    min: re.FindString(rttValues[0]),
    avg: re.FindString(rttValues[1]),
    max: re.FindString(rttValues[2]),
  }
}

func influxdb(r body) {
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
	tags := map[string]string{"target": r.target, "region": r.region}
  fields := map[string]interface{}{
    "target", r.target,
    "region", r.region,
    "transmitted", r.transmitted,
    "received", r.received,
    "loss", r.loss,
    "min", r.min,
    "avg", r.avg,
    "max", r.max,
	}

	pt, err := client.NewPoint("test_test", tags, fields, time.Now())
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
