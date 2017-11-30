package vmputty

import (
    "log"
    "strings"

    "golang.org/x/sys/windows/registry"
)

/**
 * Get registry's keys
 */
func GetPuttySession() ([]string, int) {
    //key, err := registry.OpenKey(registry.CURRENT_USER, "Software\\SimonTatham\\PuTTY\\Sessions", registry.QUERY_VALUE)
    //key, err := registry.OpenKey(registry.CURRENT_USER, `Software\SimonTatham\PuTTY\Sessions`, registry.QUERY_VALUE)

    /********************************
     USE registry.ALL_ACCESS to avoid

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

    /*
    // Use '_' if we don't need the value after
    // but function stat() returns 2 values,
    // so it's mandatory to have 2 variables to get the returned values
    if _, err = key.Stat(); err != nil {
        log.Fatal(err)
    }
    */

    // convert to int to remove error "cannot use count (type uint32) as type int in argument to key.ReadSubKeyNames"
    // count has not be declared, so we use :=
    // We can also do:
    // 		var count
    // 		count = int(keyInfo.SubKeyCount)
    count := int(keyInfo.SubKeyCount)

    if count > 0 {
        //Sessions := make([]vmjson.Session, 0)
        //fmt.Println("Putty have", count, "sessions:")

        //sessions, err := key.ReadSubKeyNames(count)
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
    } else {
        //fmt.Println("There is no Putty's sessions available.")
        return nil, 0
    }
}
