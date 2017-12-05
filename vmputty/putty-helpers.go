package vmputty

import (
    "log"
    "strings"

    "golang.org/x/sys/windows/registry"
)

// GetPuttySession : Get registry's keys
func GetPuttySession() ([]string, int) {
    //key, err := registry.OpenKey(registry.CURRENT_USER, "Software\\SimonTatham\\PuTTY\\Sessions", registry.QUERY_VALUE)
    //key, err := registry.OpenKey(registry.CURRENT_USER, `Software\SimonTatham\PuTTY\Sessions`, registry.QUERY_VALUE)

    /********************************
     USE registry.ALL_ACCESS
     NOT registry.QUERY_VALUE
     to avoid:

     Access is denied.

     in calling key.ReadSubKeyNames(-1)
     ********************************/

    key, err := registry.OpenKey(registry.CURRENT_USER, `Software\SimonTatham\PuTTY\Sessions`, registry.ALL_ACCESS)
    if err != nil {
        log.Fatal(err)
    }

    // Defer helps the followwing to be executed ONLY when the surrounding function returns
    // So key.close() will only append when registry.OpenKey() will finish
    defer key.Close()

    // Call stat() to get the amount of Putty's sessions
    keyInfo, err := key.Stat()
    if err != nil {
        log.Fatal(err)
    }

    // convert to int to remove error "cannot use count (type uint32) as type int in argument to key.ReadSubKeyNames"
    // count has not be declared, so we use :=
    // We can also do:
    // 		var count
    // 		count = int(keyInfo.SubKeyCount)
    count := int(keyInfo.SubKeyCount)

    if count > 0 {
        // If n <= 0, ReadSubKeyNames returns all the names from the key in a single slice
        sessions, err := key.ReadSubKeyNames(-1)
        if err != nil {
            log.Println("Error Registry.ReadSubKeyNames")
            log.Fatal(err)
        }

        // for index, value in range(value)
        // We can jump 'index' or 'value' by using '_'
        returnSessions := make([]string, 0, count)
        for _, session := range sessions {
            session = strings.Replace(session, "%20", " ", -1)
            //fmt.Println("-", session)
            returnSessions = append(returnSessions, session)
        }
        return returnSessions, count
    }
    
    return nil, 0
}
