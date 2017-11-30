package main

import (
    "./vmjson"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

/*
type ConfigVM struct {
    Access        string          `json:"access.default.program"`
    Ssh           Ssh             `json:"access.ssh"`
    Putty         Putty           `json:"access.putty"`
    VBoxManage    VBoxManage      `json:"VBoxManage"`
    VirtualBoxVMS []VirtualBoxVMS `json:"VirtualBoxVMS"`
}

type Ssh struct {
    Program string `json:"program"`
}

type Putty struct {
    Program string `json:"program"`
    Default string `json:"default.session"`
    Sessions []string `json:"sessions"`
}

type VBoxManage struct {
    DefaultVM string `json:"default_vm"`
    KeyIP     string `json:"key_ip"`
}

type VirtualBoxVMS struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
}
*/

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("LaunchVM.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened File as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our configVM array
    var configVM vmjson.ConfigVM

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'configVM' which we defined above
    json.Unmarshal(byteValue, &configVM)

    // we iterate through configVM.sessions
    fmt.Println("access.default.program:", configVM.Access)
    fmt.Println("Program:", configVM.Putty.Program)
    fmt.Println("Default Session:", configVM.Putty.Default)

    countPutty := len(configVM.Putty.Sessions)

    fmt.Printf("There is %d sessions:\n", countPutty)
    for i := 0; i < countPutty; i++ {
        fmt.Println(configVM.Putty.Sessions[i])
    }

    countVM := len(configVM.VirtualBoxVMS)
    fmt.Printf("There is %d VMs:\n", countVM)
    for i := 0; i < countVM; i++ {
        fmt.Println("    Name:", configVM.VirtualBoxVMS[i].Name)
        fmt.Println("    UUID:", configVM.VirtualBoxVMS[i].Uuid)
        fmt.Println("    User:", configVM.VirtualBoxVMS[i].Username)
        fmt.Println("    Pass:", configVM.VirtualBoxVMS[i].Password)
        fmt.Println("----------")
    }
}
