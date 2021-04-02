package annotation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fader4/okdoc/syntax"
	"github.com/pkg/errors"
)

type Annotation struct {
	Start, End *syntax.Token

	name   syntax.Ident_
	fields AnnotationFields

	// ref to source data
	lex *syntax.Lexer
}

func (a *Annotation) MarshalJSON() ([]byte, error) {
	res := map[string]interface{}{
		"name":      a.Name(),
		"fields":    a.Fields(),
		"num_chars": a.Len(),
		"pos": map[string]interface{}{
			"start_line":        a.Start.Pos.Line(),
			"end_line":          a.End.Pos.Line(),
			"start_left_chars":  a.Start.Pos.Spaces(),
			"start_left_spaces": a.Start.Pos.Spaces(),
		},
	}
	return json.MarshalIndent(res, " ", " ")
}

func (a *Annotation) Parse() error {
	if a.lex == nil {
		return fmt.Errorf("annotation#Parse: nil lexer")
	}
	dat, err := a.Token().Bytes(a.lex.Data)
	if err != nil {
		return err
	}
	lex, err := newTokenizer(dat)
	if err != nil {
		return errors.Wrap(err, "tokenization error")
	}
	interestedTokens := []string{
		"bracket",
		"at",
		"op_and_punct",
		"ident",
		"literal",
	}
	lexer := &annotationLex{
		LexIter: syntax.NewLexIter(lex, interestedTokens...),
		// debug:   true,
		annot: a,
	}
	out := annotationParse(lexer)
	if out != 0 {
		return fmt.Errorf("parser error: %v", lexer.Errors)
	}

	return nil
}

func (a Annotation) Name() string {
	if len(a.name) == 0 {
		return ""
	}
	return a.name.String()
}

func (a Annotation) Fields() syntax.Array {
	res := syntax.Array{}
	for _, field := range a.fields {
		switch in := field.Key.(type) {
		case syntax.Null_:
			res.Add(syntax.Array{syntax.Null_{}, field.Value})
		case syntax.StringLiteral:
			res.Add(syntax.Array{in, field.Value})
		case syntax.Ident_:
			if len(in) > 0 {
				res.Add(syntax.Array{in, field.Value})
			} else {
				res.Add(syntax.Array{syntax.Null_{}, field.Value})
			}
		case nil:
			res.Add(syntax.Array{syntax.Null_{}, field.Value})
		default:
			log.Fatalf("Annotation#Fields: not supported key type %T\n", field.Key)
		}
	}
	return res
}

type AnnotationFields []*AnnotationField

type AnnotationField struct {
	Key   syntax.Value
	Value syntax.Value
}

func (a *Annotation) Token() *syntax.Token {
	return &syntax.Token{
		Symbol: annotation,
		Start:  a.Start.Start,
		End:    a.End.End,
		Pos:    a.Start.Pos,
	}
}

func (a *Annotation) Len() int {
	return a.End.End - a.Start.Start
}

func ParseFile(file string) ([]*Annotation, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return Parse(dat)
}

// Extract returns found annotations.
func Parse(dat []byte) ([]*Annotation, error) {
	lex, err := newPreprocessing([]byte(dat))
	if err != nil {
		return nil, err
	}
	iter := syntax.NewLexIter(lex, "annotation")

	var annotations = []*Annotation{}
	var startToken *syntax.Token
	for {
		token := iter.Next()
		if token == nil {
			break
		}

		switch token.Symbol {
		case beginAnnotation:
			startToken = token
		case endAnnotation:
			annot := &Annotation{
				Start: startToken,
				End:   token,
				lex:   lex,
			}
			if err := annot.Parse(); err != nil {
				log.Println("Failed parse annot", err)
			}
			annotations = append(annotations, annot)
		}
	}

	return annotations, nil
}

type annotationLex struct {
	*syntax.LexIter
	debug bool
	annot *Annotation
}

func (l *annotationLex) Lex(out *annotationSymType) (symbol int) {
	token := l.Next()
	if token == nil {
		return 0
	}
	out.token = &syntax.TokenWithData{token, l.LexIter.Lex()}

	if l.debug {
		tokenBytesm, err := token.HumanString(l.LexIter.Lex().Data)
		if err != nil {
			panic("failed render token")
		}
		fmt.Println(string(tokenBytesm))
	}
	return token.Symbol
}
