package vmjson

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"regexp"
)

// ConfigVM : Main structure
type ConfigVM struct {
	Access        string          `json:"access.default.program"`
	SSH           SSH             `json:"access.ssh"`
	Putty         Putty           `json:"access.putty"`
	VBoxManage    VBoxManage      `json:"vbox.manager"`
	VirtualBoxVMS []VirtualBoxVMS `json:"vbox.vms"`
}

// SSH : Name of the program to launch for a ssh access
type SSH struct {
	Program string `json:"program"`
}

// Putty structure
type Putty struct {
	Program  string   `json:"program"`
	Default  string   `json:"default.session"`
	Sessions []string `json:"sessions"`
}

// VBoxManage structure
type VBoxManage struct {
	Program   string `json:"program"`
	DefaultVM string `json:"default.vm"`
	KeyIP     string `json:"key.ip"`
}

// VirtualBoxVMS : Structure of each VM
type VirtualBoxVMS struct {
	Name     string `json:"name"`
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func readJSON(filename string) []byte {
	//var buffer []byte
	buffer := make([]byte, 0)

	// Use regex to avoid the comments
	regexComment := regexp.MustCompile(`(//.*)`)

	// Open our jsonFile
	jsonFile, err := os.Open(filename)

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(jsonFile)

	// Loop over all lines in the file and check them.
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

	return buffer
}

// GetJSON : Exported function
func GetJSON(filename string) ConfigVM {
	dataJSON := readJSON(filename)

	// we initialize our configVM as type ConfigVM
	var config ConfigVM

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'configVM' which we defined above
	if err := json.Unmarshal(dataJSON, &config); err != nil {
		log.Fatalf("Unmarshal: %v\n", err)
	}

	return config
}
