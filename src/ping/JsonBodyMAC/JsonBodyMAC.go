/* MAC OS
### Ping successful ###
0 PING
1 ss.shared.ams1.websys.tmcs
2 (10.75.239.173):
3 56
4 data
5 bytes
64
6 bytes
7 from
8 10.75.239.173:
9 icmp_seq=0
10 ttl=250
11 time=24.537
12 ms

---
13 ss.shared.ams1.websys.tmcs
14 ping
15 statistics
16 ---
1
17 packets
18 transmitted,
19 1
20 packets
21 received,
22 0.0%
23 packet
24 loss
round-trip
25 min/avg/max/stddev
26 =
27 24.537/24.537/24.537/0.000
28 ms

### Ping failed ###
0 PING
1 1.2.3.4
2 (1.2.3.4):
3 56
4 data
5 bytes

---
6 1.2.3.4
7 ping
8 statistics
9 ---
1
10 packets
11 transmitted,
12 0
13 packets
14 received,
15 100.0%
16 packet
17 loss
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

/* json API body for MAC ping comman
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
  if strings.Contains(splittedValues[11], "time=") {
    rttValues := strings.Split(splittedValues[27],"/")
    transmitted, _ := strconv.ParseFloat(re.FindString(splittedValues[16]),64)
    received, _ := strconv.ParseFloat(re.FindString(splittedValues[19]),64)
    loss, _ := strconv.ParseFloat(re.FindString(splittedValues[22]),64)
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
    transmitted, _ := strconv.ParseFloat(re.FindString(splittedValues[9]),64)
    received, _ := strconv.ParseFloat(re.FindString(splittedValues[12]),64)
    loss, _ := strconv.ParseFloat(re.FindString(splittedValues[15]),64)
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
