package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
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
		lineFile := lineFile(file)

		tokens, err := starlark.ParseFile(file)
		outTokens := []interface{}{}
		for _, token := range tokens {
			if err := token.Parse(); err != nil {
				log.Printf("Failed parse token %q: %v\n", file, err)
				continue
			}
			report := newRepot(token)
			report.AddAroundCode(lineFile)
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
	Comment         *starlark.Comment        `json:"comment,omitempty"`
	Return          *starlark.Return         `json:"return,omitempty"`
	Load            *starlark.Load           `json:"load,omitempty"`
	Module          *starlark.Module         `json:"module,omitempty"`
	Def             *starlark.Def            `json:"def,omitempty"`
	Raw             string                   `json:"raw,omitempty"`
	CodeAroundLines []int                    `json:"code_around_lines,omitempty"`
	CodeAround      []string                 `json:"code_around,omitempty"`
	NumChars        int                      `json:"num_chars"`
	Annots          []*annotation.Annotation `json:"annots,omitempty"`
	Pos             struct {
		StartLine       int `json:"start_line"`
		EndLine         int `json:"end_line"`
		StartLeftChars  int `json:"start_left_chars"`
		StartLeftSpaces int `json:"start_left_spaces"`
	} `json:"pos"`
}

func (r *Report) AddAroundCode(lineFile []string) {
	r.CodeAroundLines = []int{}

	if r.Pos.StartLine-2 > 0 {
		r.CodeAroundLines = append(r.CodeAroundLines, r.Pos.StartLine-2)
		r.CodeAround = append(r.CodeAround, lineFile[r.Pos.StartLine-2])
	}

	if r.Pos.StartLine-1 > 0 {
		r.CodeAroundLines = append(r.CodeAroundLines, r.Pos.StartLine-1)
		r.CodeAround = append(r.CodeAround, lineFile[r.Pos.StartLine-1])
	}

	r.CodeAroundLines = append(r.CodeAroundLines, r.Pos.StartLine)
	r.CodeAround = append(r.CodeAround, lineFile[r.Pos.StartLine])

	if r.Pos.EndLine != r.Pos.StartLine {
		for i := r.Pos.StartLine + 1; i <= r.Pos.EndLine && i <= len(lineFile); i++ {
			r.CodeAroundLines = append(r.CodeAroundLines, i)
			r.CodeAround = append(r.CodeAround, lineFile[i])
		}
	}

	if r.Pos.EndLine+1 < len(lineFile) {
		r.CodeAroundLines = append(r.CodeAroundLines, r.Pos.EndLine+1)
		r.CodeAround = append(r.CodeAround, lineFile[r.Pos.EndLine+1])
	}

	if r.Pos.EndLine+2 < len(lineFile) {
		r.CodeAroundLines = append(r.CodeAroundLines, r.Pos.EndLine+2)
		r.CodeAround = append(r.CodeAround, lineFile[r.Pos.EndLine+2])
	}
}

func lineFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}
