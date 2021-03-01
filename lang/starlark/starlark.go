package starlark

import (
	"fmt"

	"github.com/pkg/errors"
)

func Parse(dat []byte) (*ReportFile, error) {
	lex, err := starlarkLex(dat)
	if err != nil {
		return nil, errors.Wrap(err, "tokenization error")
	}
	lexForParser := &lexerTokenIter{lex: lex, onlyTokens: []string{
		"bracket",
		"string",
		"comment",
		"keyword",
		"op_and_punct",
		"ident",
	}}
	out := starlarkParse(lexForParser)
	if out != 0 {
		return nil, fmt.Errorf("parser error: %v", lexForParser.errors)
	}

	// hydrate comments to each function (if exists)
	for _, fn := range lexForParser.res.Functions {
		for _, comment := range lexForParser.res.Comments {
			if comment.Pos.Line() == fn.Pos.Line()-1 {
				fn.Comment = comment
			}
		}
	}

	return lexForParser.res, nil
}

var _ reporter = (*lexerTokenIter)(nil)

type lexerTokenIter struct {
	// associated tokenizer
	lex *lexer
	// iterator of tokens
	iter int
	// filter tokens interesting to the parser by labels
	onlyTokens []string
	errors     []string

	res *ReportFile
}

func (l *lexerTokenIter) SetReport(r *ReportFile) {
	l.res = r
}

type reporter interface {
	SetReport(*ReportFile)
}

func (l *lexerTokenIter) Lex(out *starlarkSymType) (symbol int) {

skipToken:
	if !l.hasNext() {
		// EOF
		return 0
	}

	token := l.next()

	// if sets filter and not matched - skip token for current parser
	if len(l.onlyTokens) != 0 && !token.matchAtLeastOneLabels(l.onlyTokens...) {
		goto skipToken
	}

	out.token = &tokenWithLexer{token: token, lex: l.lex}

	if starlarkDebug == 1 {
		firstChar := token.pos[1] == token.pos[2]

		lineInfo := fmt.Sprintf("\t>%d:\t", token.pos[1])
		if firstChar {
			lineInfo = fmt.Sprintf("L%dT%d:", token.pos[0], token.pos[2])
		}
		charInfo := fmt.Sprintf("%q", string(token.Read(l.lex)))
		fmt.Println(
			"[PARSER]",
			lineInfo,
			charInfo,
		)
	}
	return token.symbol
}

func (l *lexerTokenIter) next() *token {
	if !l.hasNext() {
		return nil
	}
	token := l.lex.tokens[l.iter]
	l.iter++
	return token
}

func (l *lexerTokenIter) hasNext() bool {
	if l.iter == len(l.lex.tokens) {
		return false
	}
	return true
}

func (l *lexerTokenIter) current() *token {
	return l.lex.tokens[l.iter]
}

func (l *lexerTokenIter) Error(msg string) {
	token := l.current()
	msg = fmt.Sprintf(
		"parser error #%d: err=%q before %q",
		l.iter,
		msg,
		string(token.Read(l.lex)),
	)
	l.errors = append(l.errors, msg)
}
