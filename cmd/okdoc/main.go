package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/fader4/okdoc/annotation"
)

var pathF = flag.String("path", "", "Path to files")
var outF = flag.String("out", "okdoc_annots.json", "path to the file with the result")

func main() {
	flag.Parse()
	log.Println("Interested path with files", *pathF)
	found, err := filepath.Glob(*pathF)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Found interesting files", found)
	res := &Result{Files: map[string]interface{}{}}
	for _, file := range found {
		annots, err := annotation.ParseFile(file)
		if err != nil {
			log.Printf("Failed parse file %q: %v\n", file, err)
			continue
		}
		res.Files[file] = annots
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetIndent(" ", " ")
	if err := enc.Encode(res); err != nil {
		log.Fatalln("Failed to encode json", err)
	}
	if err := ioutil.WriteFile(*outF, buf.Bytes(), 0777); err != nil {
		log.Fatalln("Failed to write to file", err)
	}
}

type Result struct {
	Files map[string]interface{} `json:"files"`
}
