package yaml

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "log"
)

type YamlTargets struct {
  Region string
  Targets []string
}

/*
YAML file parser based on the below struct:

type yamlTargets struct {
  Region string
  Targets []string
}
*/
func YamlReader() YamlTargets {
  // define variables
  var y YamlTargets
  // open YAML file
  yamlFile, err := ioutil.ReadFile("var/targets.yaml")
  if err != nil {
      log.Printf("yaml file get err %v ", err)
  }
  // read YAML file
  err = yaml.Unmarshal(yamlFile, &y)
  if err != nil {
      log.Fatalf("unmarshal error: %v", err)
  }
  return y
}
