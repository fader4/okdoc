package starlark

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fader4/okdoc/syntax"
	"github.com/pkg/errors"
)

type Token struct {
	*syntax.CompositeToken

	Atoms []*syntax.Token

	Comment *Comment
	Return  *Return
	Load    *Load
	Module  *Module
	Def     *Def
}

func (a *Token) Len() int {
	return a.End.End - a.Start.Start
}

func (t *Token) MarshalJSON() ([]byte, error) {
	out := struct {
		Comment *Comment `json:"comment,omitempty"`
		Return  *Return  `json:"return,omitempty"`
		Load    *Load    `json:"load,omitempty"`
		Module  *Module  `json:"module,omitempty"`
		Def     *Def     `json:"def,omitempty"`
		Raw     string   `json:"raw,omitempty"`

		NumChars int `json:"num_chars"`
		Pos      struct {
			StartLine       int `json:"start_line"`
			EndLine         int `json:"end_line"`
			StartLeftChars  int `json:"start_left_chars"`
			StartLeftSpaces int `json:"start_left_spaces"`
		} `json:"pos"`
	}{
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

	return json.MarshalIndent(out, " ", " ")
}

func ParseFile(file string) ([]*Token, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return Parse(dat)
}

func Parse(dat []byte) ([]*Token, error) {
	lex, err := newPreprocessing([]byte(dat))
	if err != nil {
		return nil, err
	}

	iter := syntax.NewLexIter(lex, "def", "comment", "return", "load", "module")

	var tokens = []*Token{}
	var startToken *syntax.Token

	for {
		token := iter.Next()
		if token == nil {
			break
		}

		switch token.Symbol {
		case commentMultiline, def, load, module:
			startToken = token

		case returnKeyword:
			newToken := &syntax.CompositeToken{
				Start:  token,
				End:    token,
				Lex:    lex,
				Symbol: returnKeyword,
			}
			tokens = append(tokens, &Token{CompositeToken: newToken})
		case commentInline:
			newToken := &syntax.CompositeToken{
				Start:  token,
				End:    token,
				Lex:    lex,
				Symbol: commentInline,
			}
			tokens = append(tokens, &Token{CompositeToken: newToken})
		case endModule:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    token,
				Lex:    lex,
				Symbol: module,
			}
			tokens = append(tokens, &Token{CompositeToken: newToken})
		case endCommentMultiline:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    token,
				Lex:    lex,
				Symbol: commentMultiline,
			}
			tokens = append(tokens, &Token{CompositeToken: newToken})
		case endLoad:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    token,
				Lex:    lex,
				Symbol: load,
			}
			tokens = append(tokens, &Token{CompositeToken: newToken})
		case endDef:
			newToken := &syntax.CompositeToken{
				Start:  startToken,
				End:    token,
				Lex:    lex,
				Symbol: def,
			}
			tokens = append(tokens, &Token{CompositeToken: newToken})
		}
	}

	return tokens, nil
}

func (t *Token) Parse() error {
	lex, err := newTokenizer(t.MustBytes())
	if err != nil {
		return errors.Wrap(err, "tokenization error")
	}
	interestedTokens := []string{
		"bracket",
		"literal",
		"ident",
		"op_and_punct",
		"dict",

		"comment",
		"def",
		"load",
		"module",
		"return",
	}
	lexer := &starlarkLex{
		LexIter: syntax.NewLexIter(lex, interestedTokens...),
		// debug:   true,
		Token: t,
	}
	out := starlarkParse(lexer)
	if out != 0 {
		return fmt.Errorf("parser error: %v", lexer.Errors)
	}
	return nil
}

type starlarkLex struct {
	*syntax.LexIter
	debug bool
	*Token
}

func (l *starlarkLex) Lex(out *starlarkSymType) (symbol int) {
	token := l.Next()
	if token == nil {
		return 0
	}
	l.Atoms = append(l.Atoms, token)
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
