/**
  When we launch for the first time, we need to :
  A - Check if VirtualBox and Putty are installed
  B - Always get the list of VM
  C - json file exist?
  D - If not, ask for installation:
      List of Virtualbox's VMs:
          1. "CentOS 7"
          2. "Ubuntu 16.04.3 Server"

      Select the vm to launch: _
  D - Save new JSON file
  E - If yes, check if the default vm is still in the list we got from (B)
*/
package main

import (
    "fmt"

    "./vmbox"
    "./vmjson"
    "./vmputty"

    "io"
    "log"
    "os"
)

// JSONFILENAME : JSON Filename
const JSONFILENAME = "LaunchVM.json"

func readJSON(filename string) {
    jsonFile, err := os.Open(filename)
    if err != nil {
        fmt.Printf("File error: %v\n", err)
        os.Exit(1)
    }
    defer jsonFile.Close()

    // Get the status of the File to get its size
    fileInfo, err := jsonFile.Stat()

    // define our buffer size here
    buf := make([]byte, fileInfo.Size())

    // Loop until EOF
    for {
        n, err := jsonFile.Read(buf)

        if n > 0 {
            fmt.Print(string(buf[:n])) // your read buffer.
        }

        if err == io.EOF {
            break
        }
        if err != nil {
            log.Printf("read %d bytes: %v", n, err)
            break
        }
    }
}

func extractJSON(filename string) {
    // Open, read and get the structurized data from json file
    configVM := vmjson.GetJSON(filename)

    // Just print the data we got
    fmt.Println("Default access entity", configVM.Access)

    fmt.Println("access.ssh:")
    fmt.Println("Program:", configVM.Ssh.Program)

    fmt.Println("access.putty:")
    fmt.Println("Program:", configVM.Putty.Program)
    fmt.Println("Default Session:", configVM.Putty.Default)

    // we iterate through JSONconfig.sessions
    count := len(configVM.Putty.Sessions)
    fmt.Printf("There is %d sessions:\n", count)
    for i := 0; i < count; i++ {
        fmt.Println(configVM.Putty.Sessions[i])
    }

    count = len(configVM.VirtualBoxVMS)
    fmt.Printf("There is %d VMs:\n", count)
    for i := 0; i < count; i++ {
        fmt.Println("    Name:", configVM.VirtualBoxVMS[i].Name)
        fmt.Println("    UUID:", configVM.VirtualBoxVMS[i].Uuid)
        fmt.Println("    User:", configVM.VirtualBoxVMS[i].Username)
        fmt.Println("    Pass:", configVM.VirtualBoxVMS[i].Password)
        fmt.Println("----------")
    }
}

// In 'vboxmanage.go' init() is always called, regardless if there's main or not,
// so if you import a package that has an init function, it will be executed
func main() {
    if vmbox.VBoxManagePath != "" {
        fmt.Println("VboxManage is in the following directory:", vmbox.VBoxManagePath)
        fmt.Println("----------")
        listVMS, _ := vmbox.ListVMS()
        fmt.Printf("List of %d Virtualbox's VMs:\n", len(listVMS))

        for _, listVM := range listVMS {
            fmt.Printf("%-40s - UUID: %s\n", listVM.Name, listVM.UUID)
        }

        fmt.Println("----------")
    } else {
        fmt.Println("VirtualBox is not installed or cannot find VBoxManage.exe")
    }

    fmt.Println("List of Putty's sessions:")

    sessions, count := vmputty.GetPuttySession()
    if count > 0 {
        fmt.Println("Putty have", count, "sessions:")
        for _, session := range sessions {
            fmt.Println("-", session)
        }
    } else {
        fmt.Println("There is no Putty's sessions available.")
    }

    fmt.Println("----------")
    fmt.Printf("Data in %s file:\n", JSONFILENAME)

    extractJSON(JSONFILENAME)
}
