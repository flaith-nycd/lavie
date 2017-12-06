package vmjson

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
	// Open our jsonFile
	jsonFile, err := os.Open(filename)

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened File as a byte array.
	dataJSON, _ := ioutil.ReadAll(jsonFile)

	return dataJSON
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
