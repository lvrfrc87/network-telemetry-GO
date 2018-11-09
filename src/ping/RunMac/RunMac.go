package RunMac

import (
  "os/exec"
  "os"
  "strings"
)

/*
local ping commamnd for MAC OS
Send 1 packet every second. Wait 1 second before to timeout
*/
func RunPing(ip string) []string {
  output, err := exec.Command("ping", "-c", "1", "-t", "1", ip).CombinedOutput()
  if err != nil {
    os.Stderr.WriteString(err.Error())
  }
  return strings.Split(string(output), " ")
}
