package main

import (
    "bytes"
    "fmt"
    "golang.org/x/crypto/ssh"
    "io/ioutil"
    "log"
)

func main() {
  cmd := "ping -c 1 8.8.8.8"
  hosts := []string {
    "ec2-3-8-77-175.eu-west-2.compute.amazonaws.com",
  }
  results := make(chan string, 10)

  // read the private key
  key, err := ioutil.ReadFile("/Users/federicoolivieri/Downloads/olivierif.pem")
  if err != nil {
    log.Fatalf("unable to read private key: %v", err)
    }
  // parse the prive key and create a ssh caller
  singer, err := ssh.ParsePrivateKey(key)
  if err != nil {
      log.Fatalf("unable to parse private key: %v", err)
  }
  config := &ssh.ClientConfig{
    Auth: []ssh.AuthMethod{
      ssh.PublicKeys(singer),},
    // var hostKey ssh.PublicKey
    // HostKeyCallback: ssh.FixedHostKey(hostKey)
    HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

  for _, hostname := range hosts {
    go func(hostname string) {
      results <- executeCmd(cmd, hostname, config)
    }(hostname)
  }

  for i := 0; i < len(hosts); i++ {
      select {
      case res := <-results:
          fmt.Print(res)
        }
    }
}

func executeCmd(command, hostname string, config *ssh.ClientConfig) string {
  client, err := ssh.Dial("tcp", "22", hostname, config)
  if err != nil {
      log.Fatalf("unable to connect: %v", err)
  }
  session, err := client.NewSession()
  if err != nil {
      log.Fatal("Failed to create session: ", err)
  }
  defer session.Close()

  var stdoutBuf bytes.Buffer
  session.Stdout = &stdoutBuf
  session.Run(command)

  return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String())
}
