package RunMac

import (
  "os/exec"
  "strings"
)

/*
local ping commamnd for MAC OS
Send 1 packet every second. Wait 1 second before to timeout
*/
func RunPing(ip string) []string {
  output, _ := exec.Command("ping", "-c", "1", "-t", "1", ip).CombinedOutput()
  return strings.Fields(string(output))
}
