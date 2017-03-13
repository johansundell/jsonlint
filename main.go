package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
			log.Fatal("Could not open file", filename, "error ", err)
		}
	} else {
		info, _ := os.Stdin.Stat()
		if (info.Mode() & os.ModeCharDevice) == 0 {
			fmt.Println("Pipe data", info.Size())
			if b, err = ioutil.ReadAll(os.Stdin); err != nil {
				log.Fatal("Could not read input", err)
			}
		} else {
			log.Fatal("No input")
		}
	}
	var input map[string]interface{}
	if err := json.Unmarshal(b, &input); err != nil {
		log.Fatal("Not valid json", err)
	}
	result, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		log.Fatal("Could not export", err)
	}
	fmt.Println(string(result))
}
