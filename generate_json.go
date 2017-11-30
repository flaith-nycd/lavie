// See https://play.golang.org/p/jKkb86p5SG
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"

    vm "./vmbox"
    "./vmjson"
    putty "./vmputty"
)

func main() {
    // Get putty
    sessions, count := putty.GetPuttySession()
    //onePuttySession := make([]vmjson.Session, count)
    onePuttySession := make([]string, count)
    if count > 0 {
        for i, session := range sessions {
            //onePuttySession[i].Name = session
            onePuttySession[i] = session
        }
    } else {
        fmt.Println("There is no Putty's sessions available.")
    }

    // Get Vms
    listVMS, _ := vm.ListVMS()
    oneVM := make([]vmjson.VirtualBoxVMS, len(listVMS))
    for i, VM := range listVMS {
        oneVM[i].Name = VM.Name
        oneVM[i].Uuid = VM.UUID
    }

    // Set main structure ConfigVM
    allVMs := vmjson.ConfigVM{
        Access: "ssh",
        Ssh:    vmjson.Ssh{Program: "C:\\app\\TOOLS\\git\\usr\\bin\\ssh.exe"},
        Putty: vmjson.Putty{
            Program:  "C:\\app\\TOOLS\\putty\\PUTTY.EXE",
            Sessions: onePuttySession,
        },
        VBoxManage: vmjson.VBoxManage{
            Program: vm.VBM,
        },
        VirtualBoxVMS: oneVM,
    }

    // Generate JSON
    configVMJson, err := json.MarshalIndent(allVMs, "", "    ")
    if err != nil {
        log.Fatal("Cannot encode to JSON ", err)
    }

    jsonFile, _ := os.Create("testJson.json")
    //fmt.Printf("%s", configVMJson)
    jsonFile.WriteString(fmt.Sprintf("%s", configVMJson))
    jsonFile.Close()
}
