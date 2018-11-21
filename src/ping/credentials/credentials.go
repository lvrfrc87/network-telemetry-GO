package credpass

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
)

/*
extract username and password from ~/.credentials.json
hostname must be passad as argument - type  string
*/
type Load string

func (h Load) CredPass(u string) string {
  // variable initialization
  var result map[string]interface{}

  //read file
	byteValue, err:= ioutil.ReadFile("/root/.credentials.json")
  if err != nil {fmt.Println(err)}

  // unmarshal json file and assign to result variable
  json.Unmarshal(byteValue, &result)

  // return username and passwd
  host := result[string(h)].(map[string]interface{})
  return string(host[u].(string))
}

// func main() {
//   instance := Load("influxdb")
//   fmt.Println(instance.credPass("username"))
//   fmt.Println(instance.credPass("password"))
// }
