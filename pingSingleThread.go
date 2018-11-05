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
    // "github.com/influxdata/influxdb/client/v2"
    "strings"
)

type result struct {
    target string
    output string
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

  s := result.target
  fmt.Println(s)
}

func executeCmd (command, port string, hostname string, config *ssh.ClientConfig) result{
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
