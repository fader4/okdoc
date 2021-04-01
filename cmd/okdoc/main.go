package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/fader4/okdoc/annotation"
	"github.com/fader4/okdoc/lang/starlark"
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
		tokens, err := starlark.ParseFile(file)
		outTokens := []interface{}{}
		for _, token := range tokens {
			if err := token.Parse(); err != nil {
				log.Printf("Failed parse token %q: %v\n", file, err)
				continue
			}
			report := newRepot(token)
			if token.Comment != nil {
				report.Annots, err = annotation.Parse(token.MustBytes())
				if err != nil {
					log.Printf("Failed parse annots %q: %v\n", file, err)
					continue
				}
			}

			outTokens = append(outTokens, report)
		}

		if err != nil {
			log.Printf("Failed parse file %q: %v\n", file, err)
			continue
		}
		res.Files[file] = outTokens
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

func newRepot(t *starlark.Token) *Report {
	return &Report{
		Comment:  t.Comment,
		Return:   t.Return,
		Load:     t.Load,
		Module:   t.Module,
		Def:      t.Def,
		NumChars: t.Len(),
		Raw:      string(t.MustBytes()),
		Pos: struct {
			StartLine       int `json:"start_line"`
			EndLine         int `json:"end_line"`
			StartLeftChars  int `json:"start_left_chars"`
			StartLeftSpaces int `json:"start_left_spaces"`
		}{
			StartLine:       t.Start.Pos.Line(),
			EndLine:         t.End.Pos.Line(),
			StartLeftChars:  t.Start.Pos.Spaces(),
			StartLeftSpaces: t.Start.Pos.Spaces(),
		},
	}
}

type Report struct {
	Comment  *starlark.Comment        `json:"comment,omitempty"`
	Return   *starlark.Return         `json:"return,omitempty"`
	Load     *starlark.Load           `json:"load,omitempty"`
	Module   *starlark.Module         `json:"module,omitempty"`
	Def      *starlark.Def            `json:"def,omitempty"`
	Raw      string                   `json:"raw,omitempty"`
	NumChars int                      `json:"num_chars"`
	Annots   []*annotation.Annotation `json:"annots,omitempty"`
	Pos      struct {
		StartLine       int `json:"start_line"`
		EndLine         int `json:"end_line"`
		StartLeftChars  int `json:"start_left_chars"`
		StartLeftSpaces int `json:"start_left_spaces"`
	} `json:"pos"`
}
