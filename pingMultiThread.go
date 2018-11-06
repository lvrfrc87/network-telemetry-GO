/*
https://zaiste.net/executing_commands_via_ssh_using_go/
https://godoc.org/golang.org/x/crypto/ssh#example-PublicKeys
*/

/* Linux Alpine
### Ping successful ###
0 PING
1 10.78.65.1
2 (10.78.65.1)
3 56(84)
4 bytes
5 of
6 data.
64
7 bytes
8 from
9 10.78.65.1:
10 icmp_seq=1
11 ttl=245
12 time=17.3
13 ms

---
14 10.78.65.1
15 ping
16 statistics
17 ---
1
18 packets
19 transmitted,
20 1
21 received,
22 0%
23 packet
24 loss,
25 time
26 0ms
rtt
27 min/avg/max/mdev
28 =
29 17.344/17.344/17.344/0.000
30 ms

### Ping failed ###
0 PING
1 1.2.3.4
2 (1.2.3.4)
3 56(84)
4 bytes
5 of
6 data.

---
7 1.2.3.4
8 ping
9 statistics
10 ---
1
11 packets
12 transmitted,
13 0
14 received,
15 100%
16 packet
17 loss,
18 time
19 0ms
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
    "time"
)

// json body structur
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
  targetIps := [] string {
    "10.63.65.1",
    "10.66.65.1",
    "10.70.65.1",
    "10.78.65.1",
    "10.226.234.116",
    "10.227.236.4",
    }
  port := "22"
  user := "core"
  hostname := "app1.net.awsieprod2.linsys.tmcs"

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

  for {
    for _, ip := range targetIps {
      go influxdb(jsonBody(runPing(ip, port, hostname, config))) // {
      }
    time.Sleep(3 * time.Second)
    }
}

func runPing(ip, port, hostname string, config *ssh.ClientConfig) []string {
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
  session.Run(fmt.Sprintf("ping -c 1 -w 1 %s", ip))

  return strings.Split(stdoutBuf.String(), " ")
}

func jsonBody(splittedValues []string) body {
  re := regexp.MustCompile(`\d+\.?\d?`)
  if strings.Contains(splittedValues[12], "time=") {
  rttValues := strings.Split(splittedValues[29],"/")
  return body {
    target: splittedValues[1],
    region: "bar",
    transmitted: re.FindString(splittedValues[17]),
    received: re.FindString(splittedValues[20]),
    loss: re.FindString(splittedValues[22]),
    min: re.FindString(rttValues[0]),
    avg: re.FindString(rttValues[1]),
    max: re.FindString(rttValues[2]),
    }
  } else {
    return body {
      target: splittedValues[1],
      region: "bar",
      transmitted: re.FindString(splittedValues[10]),
      received: re.FindString(splittedValues[13]),
      loss: re.FindString(splittedValues[15]),
      min: "0",
      avg: "0",
      max: "0",
    }
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
    "transmitted": r.transmitted,
    "received": r.received,
    "loss": r.loss,
    "min": r.min,
    "avg": r.avg,
    "max": r.max,
	}

	pt, err := client.NewPoint("go_rtt_m", tags, fields, time.Now())
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
