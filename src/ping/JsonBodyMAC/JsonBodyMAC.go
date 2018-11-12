/* MAC OS
### Ping successful ###
0 PING
1 ss.shared.ams1.websys.tmcs
2 (10.75.239.173):
3 56
4 data
5 bytes
6 64
7 bytes
8 from
9 10.75.239.173:
10 icmp_seq=0
11 ttl=250
12 time=9.194
13 ms
14 ---
15 ss.shared.ams1.websys.tmcs
16 ping
17 statistics
18 ---
19 1
20 packets
21 transmitted,
22 1
23 packets
24 received,
25 0.0%
26 packet
27 loss
28 round-trip
29 min/avg/max/stddev
30 =
31 9.194/9.194/9.194/0.000
32 ms

### Ping failed ###
0 PING
1 1.2.3.4
2 (1.2.3.4):
3 56
4 data
5 bytes
6 ---
7 1.2.3.4
8 ping
9 statistics
10 ---
11 1
12 packets
13 transmitted,
14 0
15 packets
16 received,
17 100.0%
18 packet
19 loss
*/
package JsonBodyMAC

import (
    "strings"
    "regexp"
    "strconv"
)
// json body struct
type Body struct {
  Target string
  Region string
  Transmitted float64
  Received float64
  Loss float64
  Min float64
  Avg float64
  Max float64
  Stddev float64
}

/* json API body for MAC OS ping comman
based on the below struct:

type Body struct {
  Target string
  Region string
  Transmitted float64
  Received float64
  Loss float64
  Min float64
  Avg float64
  Max float64
  Stddev float64
}
*/
func JsonBody(splittedValues []string, region string) Body {
  re := regexp.MustCompile(`\d+\.?\d?`)
  if strings.Contains(splittedValues[12], "time=") {
    rttValues := strings.Split(splittedValues[31],"/")
    transmitted, _ := strconv.ParseFloat(re.FindString(splittedValues[19]),64)
    received, _ := strconv.ParseFloat(re.FindString(splittedValues[22]),64)
    loss, _ := strconv.ParseFloat(re.FindString(splittedValues[25]),64)
    min, _ := strconv.ParseFloat(re.FindString(rttValues[0]),64)
    avg, _ := strconv.ParseFloat(re.FindString(rttValues[1]),64)
    max, _ := strconv.ParseFloat(re.FindString(rttValues[2]),64)
    stddev, _ := strconv.ParseFloat(re.FindString(rttValues[3]),64)
    return Body {
      Target: string(splittedValues[1]),
      Region: string(region),
      Transmitted: transmitted,
      Received: received,
      Loss: loss,
      Min: min,
      Avg: avg,
      Max: max,
      Stddev: stddev,
      }
  } else {
    transmitted, _ := strconv.ParseFloat(re.FindString(splittedValues[11]),64)
    received, _ := strconv.ParseFloat(re.FindString(splittedValues[14]),64)
    loss, _ := strconv.ParseFloat(re.FindString(splittedValues[17]),64)
    return Body {
      Target: string(splittedValues[1]),
      Region: string(region),
      Transmitted: transmitted,
      Received: received,
      Loss: loss,
      Min: float64(0),
      Avg: float64(0),
      Max: float64(0),
      Stddev: float64(0),
    }
  }
}
