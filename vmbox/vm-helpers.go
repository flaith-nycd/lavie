package vmbox

import (
    "bufio"
    "regexp"
    "strings"
)

// To be called from outside, the function must start with a uppercase letter
func RunningMachine() (error) {
    runningVMS := vbm("list", "runningvms")
    if runningVMS != nil {
        return runningVMS
    }
    return nil
}

//func ListVM() (string) {
func ListVM() ([]string) {
    value, _ := vbmOut("list", "vms")

    // split and return a string slice
    listVMS := strings.Split(value, "\n")
    return listVMS
}

// FROM the help of https://github.com/riobard/go-virtualbox/blob/master/machine.go

// Machine information.
type Machine struct {
    Name string
    UUID string
}

func ListVMS() ([]*Machine, error) {
    // Create a slice of type Machine
    var allMachines []*Machine

    // Generate our regex
    regexVM := regexp.MustCompile(`"(.+)" {([0-9a-f-]+)}`)

    out, err := vbmOut("list", "vms")
    if err != nil {
        return nil, err
    }

    // Get all lines returned by the command output in 'vbmOut'
    scanLine := bufio.NewScanner(strings.NewReader(out))
    // For each line
    for scanLine.Scan() {
        // Use the regex for the catched line
        match := regexVM.FindStringSubmatch(scanLine.Text())

        // No match, we continue (break at this point and go to the next line)
        if match == nil {
            continue
        }

        // Generate one machine information
        // match[0] is the complete line
        oneMachine := new(Machine)
        oneMachine.Name = match[1]
        oneMachine.UUID = match[2]

        // Add this new machine to our list of machines
        allMachines = append(allMachines, oneMachine)
    }
    if err := scanLine.Err(); err != nil {
        return nil, err
    }
    return allMachines, nil
}
