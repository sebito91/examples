package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// native actions receive one argument, the JSON object as a string
	argv := os.Args[1]

	// unmarshal the string to a JSON object
	var obj map[string]interface{}

	if err := json.Unmarshal([]byte(argv), &obj); err != nil {
		log.Fatal(err)
	}

	name, ok := obj["name"].(string)
	if !ok {
		name = "Stranger"
	}

	msg := map[string]string{
		"msg": fmt.Sprintf("Hello, %s!", name),
	}

	res, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", res)
}
