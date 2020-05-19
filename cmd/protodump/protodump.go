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
	in, err := ioutil.ReadAll(os.Stdin)
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
