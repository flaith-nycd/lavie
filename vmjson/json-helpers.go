package vmjson

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "os"
)

type ConfigVM struct {
    Access        string          `json:"access.default.program"`
    SSH           Ssh             `json:"access.ssh"`
    Putty         Putty           `json:"access.putty"`
    VBoxManage    VBoxManage      `json:"vbox.manager"`
    VirtualBoxVMS []VirtualBoxVMS `json:"vbox.vms"`
}

type Ssh struct {
    Program string `json:"program"`
}

type Putty struct {
    Program  string   `json:"program"`
    Default  string   `json:"default.session"`
    Sessions []string `json:"sessions"`
}

type VBoxManage struct {
    Program   string `json:"program"`
    DefaultVm string `json:"default.vm"`
    KeyIp     string `json:"key.ip"`
}

type VirtualBoxVMS struct {
    Name     string `json:"name"`
    Uuid     string `json:"uuid"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func readJSON(filename string) []byte {
    // Open our jsonFile
    jsonFile, err := os.Open(filename)

    // if we os.Open returns an error then handle it
    if err != nil {
        log.Fatal(err)
    }

    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened File as a byte array.
    dataJson, _ := ioutil.ReadAll(jsonFile)

    return dataJson
}

func GetJSON(filename string) ConfigVM {
    dataJson := readJSON(filename)

    // we initialize our configVM as type ConfigVM
    var config ConfigVM

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'configVM' which we defined above
    if err := json.Unmarshal(dataJson, &config); err != nil {
        log.Fatalf("Unmarshal: %v\n", err)
    }

    return config
}
