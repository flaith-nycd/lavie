Fpackage main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "regexp"

    "./vmjson"
)

func main() {
    //var buffer []byte
    buffer := make([]byte, 0)

    var config vmjson.ConfigVM
    regexComment := regexp.MustCompile(`(//.*)`)

    // Open the file.
    fjsonFile, err := os.Open("JSON_WITH_COMMENTS.json")

    // if we os.Open returns an error then handle it
    if err != nil {
        log.Fatal(err)
    }

    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(fjsonFile)

    // Loop over all lines in the file and print them.
    for scanner.Scan() {
        line := scanner.Text()
        // Use the regex for the catched line
        match := regexComment.FindStringSubmatch(line)

        // Matched, so we jump to the next line, because we don't want to keep the comments for json
        if match != nil {
            continue
        }

        // Now add the current line to a []byte
        // Need to use "..." as suffix in order to append a slice to another slice
        // https://golang.org/pkg/builtin/#append
        buffer = append(buffer, []byte(line)...)
    }

    json.Unmarshal(buffer, &config)
    fmt.Println("access.default.program:", config.Access)
}
