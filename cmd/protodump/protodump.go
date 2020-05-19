package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/syucream/protodump"
)

func main() {
	var in []byte
	var err error

	if len(os.Args) >= 2 {
		// filename by arg1
		in, err = ioutil.ReadFile(os.Args[1])
	} else {
		// stdin
		in, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatal(err)
	}

	msg := protodump.Message{}
	err = protodump.Unmarshal(in, msg)
	if err != nil {
		log.Fatal(err)
	}

	jsonStr, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonStr))
}
