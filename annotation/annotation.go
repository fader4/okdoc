package annotation

import (
	"fmt"

	"github.com/fader4/okdoc/syntax"
	"github.com/pkg/errors"
)

type Annotation struct {
	Start, End *syntax.Token
	Name       Ident_
	Fields     AnnotationFields
}

type AnnotationFields []*AnnotationField

type AnnotationField struct {
	Key   Value
	Value Value
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

// Extract returns found annotations.
func Parse(dat []byte) ([]*Annotation, error) {
	lex, err := newPreprocessing([]byte(dat))
	if err != nil {
		return nil, err
	}
	lexForParser := &annotationLex{
		LexIter: syntax.NewLexIter(lex, "annotation"),
		// debug:   true,
	}

	var annotations = []*Annotation{}
	var startToken *syntax.Token
	container := &annotationSymType{}
	for {
		symbol := lexForParser.Lex(container)
		if symbol == 0 {
			break
		}

		switch symbol {
		case beginAnnotation:
			startToken = container.token.Token
		case endAnnotation:
			annotations = append(annotations, &Annotation{
				Start: startToken,
				End:   container.token.Token,
			})
		}
	}

	return annotations, nil
}

func parseAnnotation(annot *Annotation, dat []byte) error {
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
		annot: annot,
	}
	out := annotationParse(lexer)
	if out != 0 {
		return fmt.Errorf("parser error: %v", lexer.Errors)
	}

	return nil
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
	out.token = &Token{token, l.LexIter.Lex()}

	if l.debug {
		tokenBytesm, err := token.HumanString(l.LexIter.Lex().Data)
		if err != nil {
			panic("failed render token")
		}
		fmt.Println(string(tokenBytesm))
	}
	return token.Symbol
}
