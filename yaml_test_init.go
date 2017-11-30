package main

import (
    "fmt"
    "io/ioutil"
    "log"

    "gopkg.in/yaml.v2"
)

type ConfigInit struct {
    Access Access `yaml:"access"`
    Env    Env    `yaml:"env"`
}

type Access struct {
    Default string    `yaml:"default"`
    Program []Program `yaml:"program"`
}

type Program struct {
    Name string `yaml:"name"`
    Uri  string `yaml:"uri"`
}

type Env struct {
    VirtualBox VirtualBox `yaml:"virtualbox"`
}

type VirtualBox struct {
    PathKey []string `yaml:"path_key"`
}

func (config *ConfigInit) getConf(filename string) *ConfigInit {
    yamlFile, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("# Config file get error\n--> %v", err)
    }
    //err = yaml.UnmarshalStrict(yamlFile, c)
    err = yaml.Unmarshal(yamlFile, config)
    if err != nil {
        log.Fatalf("# Unmarshal error\n--> %v", err)
    }

    return config
}

func main() {
    var conf ConfigInit
    conf.getConf("init.yml")

    fmt.Println(conf)
    fmt.Printf("Default Access Program: \"%s\"\n", conf.Access.Default)
    for _, val := range conf.Access.Program {
        if val.Name == conf.Access.Default {
            fmt.Printf("Program to launch: \"%s\"\n", val.Uri)
        } else {
            fmt.Printf("Name: '%s' - Program: %s\n", val.Name, val.Uri)
        }
    }

    /*
    VBoxPathKeys := make([]string, 0, len(conf.Env.VirtualBox.PathKey))
    for i, val := range conf.Env.VirtualBox.PathKey {
        VBoxPathKeys[i] = val
    }
    */

    VBoxPathKeys := make([]string, 0)
    for _, val := range conf.Env.VirtualBox.PathKey {
        VBoxPathKeys = append(VBoxPathKeys, val)
    }

    for _, val := range VBoxPathKeys {
        fmt.Printf("Key: '%s'\n", val)
    }
}
