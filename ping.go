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
)

type result struct {
    f1 string
    f2 string
}

func main() {
  // define variables
  cmd := "ping -c 1 10.78.65.1"
  port := "22"
  user := "core"
  hosts := []string {
    "app1.net.awsieprod2.linsys.tmcs",
    "app2.net.awsieprod2.linsys.tmcs",
    "app3.net.awsieprod2.linsys.tmcs",
    "app1.net.awsdeprod2.linsys.tmcs",
    "app2.net.awsdeprod2.linsys.tmcs",
    "app3.net.awsdeprod2.linsys.tmcs",
  }
  // create go routine channel
  // https://tour.golang.org/concurrency/2
  results := make(chan string, 10)

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


  for _, hostname := range hosts {
    go func(hostname string) {
      results <- executeCmd(cmd, port, hostname, config)
    }(hostname)
  }
  for i := 0; i < len(hosts); i++ {
      select {
      case res := <-results:
          fmt.Print(res)
        }
    }
}

func executeCmd(command, port string, hostname string, config *ssh.ClientConfig) (target string, splitOut []string) {
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

  return hostname, strings.Split(stdoutBuf.String(), " ")
}

func influxdb(splitted_values []string, target string, region string) {
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
}
  // Create a new point batch
	// bp, err := client.NewBatchPoints(client.BatchPointsConfig{
	// 	Database:  "network_telemetry",
	// 	Precision: "s",
	// })
	// if err != nil {
	// 	log.Fatal("Failed to create point batch: %v", err)
	// }
  //
  // // Create a point and add to batch
	// tags := map[string]string{"host": target, "region":region}
  // if 'time=' in self.splitted_values[12]
  // fields := map[string]interface{}{
  //   "transmitted": splitted_values[19],
  //   "received": splitted_values[22],
  //   "loss": splitted_values[25],
  //   "min": splitted_values[31].split("/")[0],
  //   "avg": splitted_values[31].split("/")[1],
  //   "max": splitted_values[31].split("/")[2],
	// }
