package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/syucream/protodump"
)

func toJson(msg protodump.Message) (string, error) {
	jsonStr, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}

func main() {
	force := flag.Bool("force", false, "force to print on error")
	flag.Parse()

	args := flag.Args()

	var in []byte
	var err error
	if len(args) >= 1 {
		// filename by arg1
		in, err = ioutil.ReadFile(args[0])
	} else {
		// stdin
		in, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatal(err)
	}

	msg := protodump.Message{}
	err = protodump.Unmarshal(in, msg)

	if err == nil || *force {
		jsonStr, err := toJson(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(jsonStr)
	}
	if err != nil {
		log.Fatal(err)
	}
}
