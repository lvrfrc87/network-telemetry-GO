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
package jsonBodyAlpine

// json body struct
type Body struct {
  target string
  region string
  transmitted float64
  received float64
  loss float64
  min float64
  avg float64
  max float64
}

func jsonBody(splittedValues []string, region string) body {
  re := regexp.MustCompile(`\d+\.?\d?`)
  rttValues := strings.Split(splittedValues[29],"/")
  transmitted, _ := strconv.ParseFloat(re.FindString(splittedValues[17]),64)
  received, _ := strconv.ParseFloat(re.FindString(splittedValues[20]),64)
  loss, _ := strconv.ParseFloat(re.FindString(splittedValues[22]),64)
  min, _ := strconv.ParseFloat(re.FindString(rttValues[0]),64)
  avg, _ := strconv.ParseFloat(re.FindString(rttValues[1]),64)
  max, _ := strconv.ParseFloat(re.FindString(rttValues[2]),64)
  if strings.Contains(splittedValues[12], "time=") {
  return body {
    target: string(splittedValues[1]),
    region: string(region),
    transmitted: transmitted,
    received: received,
    loss: loss,
    min: min,
    avg: avg,
    max: max,
    }
  } else {
    return body {
      target: string(splittedValues[1]),
      region: string(region),
      transmitted: transmitted,
      received: received,
      loss: loss,
      min: float64(0),
      avg: float64(0),
      max: float64(0),
    }
  }
}
