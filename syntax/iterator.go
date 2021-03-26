package syntax

import "fmt"

func NewLexIter(lex *Lexer, interestedTokens ...string) *LexIter {
	return &LexIter{
		lex:        lex,
		onlyTokens: interestedTokens,
	}
}

type LexIter struct {
	// associated tokenizer
	lex *Lexer
	// iterator of tokens
	iter int
	// filter tokens interesting to the parser by labels
	onlyTokens []string
	Errors     []string
}

func (l *LexIter) Lex() *Lexer {
	return l.lex
}

func (l *LexIter) Next() *Token {
skipToken:
	if !l.hasNext() {
		// EOF
		return nil
	}

	token := l.next()

	// if sets filter and not matched - skip token for current parser
	if len(l.onlyTokens) != 0 && !token.MatchAtLeastOneLabels(l.onlyTokens...) {
		goto skipToken
	}

	return token
}

func (l *LexIter) next() *Token {
	if !l.hasNext() {
		return nil
	}
	token := l.lex.Tokens[l.iter]
	l.iter++
	return token
}

func (l *LexIter) hasNext() bool {
	if l.iter == len(l.lex.Tokens) {
		return false
	}
	return true
}

func (l *LexIter) current() *Token {
	return l.lex.Tokens[l.iter]
}

func (l *LexIter) Error(msg string) {
	token := l.lex.Tokens[l.iter-1]
	tokenBytes, _ := l.lex.Data.Render(token)
	msg = fmt.Sprintf(
		"parser error ##%d: err=%q before %q",
		l.iter,
		msg,
		string(tokenBytes),
	)
	l.Errors = append(l.Errors, msg)
}
