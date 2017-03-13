package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "The input file")
	flag.Parse()
	var b []byte
	var err error
	if filename != "" {
		var err error
		if b, err = ioutil.ReadFile(filename); err != nil {
			fmt.Println("Could not open file", filename, "error ", err)
		}
	} else {
		if b, err = ioutil.ReadAll(os.Stdin); err != nil {
			fmt.Println("Could not read input", err)
		}
	}
	var input map[string]interface{}
	if err := json.Unmarshal(b, &input); err != nil {
		fmt.Println("Not valid json", err)
	}
	result, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		fmt.Println("Could not export", err)
	}
	fmt.Println(string(result))
}
