package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	parsingXml "github.com/parsing-files-go/parsing-xml"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    string `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	// Open JSON file
	jsonFile, err := os.Open("users.json")

	// errors...
	if err != nil {
		fmt.Println(err)
	}
	// if not error...
	fmt.Println("Succesfully Opened users.json")

	// defer the closing of our jsonFile
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	for i := 0; i < len(result); i++ {
		fmt.Println(result["Users"])
	}

	parsingXml.ReadXML()
}
