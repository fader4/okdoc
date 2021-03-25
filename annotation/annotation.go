package annotation

import (
	"fmt"

	"github.com/fader4/okdoc/syntax"
	"github.com/pkg/errors"
)

type Annotation struct {
	RawToken *syntax.Token
	RawData  []byte
}

// Extract returns found annotations.
func Extract(dat []byte) ([]*Annotation, error) {
	lex, err := newPreprocessing([]byte(dat))
	if err != nil {
		return nil, err
	}
	lexForParser := &lexerTokenIter{lex: lex, onlyTokens: []string{
		"annotation",
	}}
	var annotations = []*Annotation{}
	var nextAnnotation *syntax.Token
	for {
		container := &annotationSymType{}
		symbol := lexForParser.Lex(container)
		if symbol == 0 {
			break
		}

		switch symbol {
		case beginAnnotation:
			nextAnnotation = &syntax.Token{
				Symbol: annotation,
				Start:  container.token.Start,
				Pos: syntax.Pos{
					container.token.Pos[0],
					// NOTE: Pos[1] computed when end is found
					0,
					// NOTE: Pos[2] save number of spaces before the first token
					container.token.Pos[2],
				},
			}
		case endAnnotation:
			nextAnnotation.End = container.token.End
			nextAnnotation.Pos[1] = nextAnnotation.End - nextAnnotation.Start
			annotations = append(annotations, &Annotation{
				RawToken: nextAnnotation,
				RawData:  dat[nextAnnotation.Start:nextAnnotation.End],
			})
			nextAnnotation = nil
		}
	}

	return annotations, nil
}

func Parse(dat []byte) error {
	lex, err := newTokenizer(dat)
	if err != nil {
		return errors.Wrap(err, "tokenization error")
	}
	lexForParser := &lexerTokenIter{lex: lex, onlyTokens: []string{
		"bracket",
		"at",
		"op_and_punct",
		"ident",
		"literal",
	}}
	out := annotationParse(lexForParser)
	if out != 0 {
		return fmt.Errorf("parser error: %v", lexForParser.errors)
	}

	return nil
}

type lexerTokenIter struct {
	// associated tokenizer
	lex *syntax.Lexer
	// iterator of tokens
	iter int
	// filter tokens interesting to the parser by labels
	onlyTokens []string
	errors     []string
}

func (l *lexerTokenIter) Lex(out *annotationSymType) (symbol int) {
skipToken:
	if !l.hasNext() {
		// EOF
		return 0
	}

	token := l.next()

	// if sets filter and not matched - skip token for current parser
	if len(l.onlyTokens) != 0 && !token.MatchAtLeastOneLabels(l.onlyTokens...) {
		goto skipToken
	}

	out.token = token

	if annotationDebug == 1 {
		l.PrintToken(token)
	}
	return token.Symbol
}

func (l *lexerTokenIter) PrintToken(token *syntax.Token) {
	lineInfo := fmt.Sprintf("L%dSP%dLEN%d:", token.Pos[0], token.Pos[2], token.Pos[1])

	tokenBytes, _ := l.lex.Data.Render(token)
	charInfo := fmt.Sprintf("%q", string(tokenBytes))
	fmt.Println(
		lineInfo,
		charInfo,
	)
}

func (l *lexerTokenIter) next() *syntax.Token {
	if !l.hasNext() {
		return nil
	}
	token := l.lex.Tokens[l.iter]
	l.iter++
	return token
}

func (l *lexerTokenIter) hasNext() bool {
	if l.iter == len(l.lex.Tokens) {
		return false
	}
	return true
}

func (l *lexerTokenIter) current() *syntax.Token {
	return l.lex.Tokens[l.iter]
}

func (l *lexerTokenIter) Error(msg string) {
	token := l.lex.Tokens[l.iter-1]
	tokenBytes, _ := l.lex.Data.Render(token)
	msg = fmt.Sprintf(
		"parser error #%d: err=%q before %q",
		l.iter,
		msg,
		string(tokenBytes),
	)
	l.errors = append(l.errors, msg)
}
