
package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type Conf struct {
    Hits int64 `yaml:"hits"`
    Time int64 `yaml:"time"`
}

func (c *Conf) getConf() *Conf {

    yamlFile, err := ioutil.ReadFile("configure.yml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {
    var c Conf
    c.getConf()

    fmt.Println(c)
}
