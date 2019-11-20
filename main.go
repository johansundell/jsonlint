package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var Version = "development"

func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "The input file")
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "Display current version")
	flag.Parse()
	if showVersion {
		fmt.Println("Version: ", Version)
	}
	var b []byte
	var err error
	if filename != "" {
		var err error
		if b, err = ioutil.ReadFile(filename); err != nil {
			log.Fatal("Could not open file", filename, "error ", err)
		}
	} else {
		info, _ := os.Stdin.Stat()
		if (info.Mode() & os.ModeCharDevice) == 0 {
			if b, err = ioutil.ReadAll(os.Stdin); err != nil {
				log.Fatal("Could not read input", err)
			}
		} else {
			log.Fatal("No input")
		}
	}
	var input interface{}
	if err := json.Unmarshal(b, &input); err != nil {
		log.Fatal("Not valid json", err)
	}
	result, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		log.Fatal("Could not export", err)
	}
	fmt.Println(string(result))
}
