package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/fader4/okdoc/annotation"
)

var pathFileF = flag.String("file", "", "Path to file")

func main() {
	flag.Parse()

	list, err := annotation.ParseFile(*pathFileF)
	if err != nil {
		log.Fatalln("Failed parse file", err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent(" ", " ")
	enc.Encode(list)
}
