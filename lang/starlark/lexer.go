package starlark

import "fmt"

func newLexer(data []byte) (lex *lexer, eof int) {
	return &lexer{
		data:   data,
		pe:     len(data),
		tokens: tokens{},
	}, len(data)
}

type lexer struct {
	// It must be an array containting the data to process.
	data []byte

	// Data end pointer. This should be initialized to p plus the data length on every run of the machine. In Java and Ruby code this should be initialized to the data length.
	pe int

	// Data pointer. In C/D code this variable is expected to be a pointer to the character data to process. It should be initialized to the beginning of the data block on every run of the machine. In Java and Ruby it is used as an offset to data and must be an integer. In this case it should be initialized to zero on every run of the machine.
	p int

	// This must be a pointer to character data. In Java and Ruby code this must be an integer. See Section 6.3 for more information.
	ts int

	// Also a pointer to character data.
	te int

	// This must be an integer value. It is a variable sometimes used by scanner code to keep track of the most recent successful pattern match.
	act int

	// Current state. This must be an integer and it should persist across invocations of the machine when the data is broken into blocks that are processed independently. This variable may be modified from outside the execution loop, but not from within.
	cs int

	// This must be an integer value and will be used as an offset to stack, giving the next available spot on the top of the stack.
	top int

	// This must be an array of integers. It is used to store integer values representing states. If the stack must resize dynamically the Pre-push and Post-Pop statements can be used to do this (Sections 5.6 and 5.7).
	stack []int

	////////////////////////////////
	// helper fields
	////////////////////////////////

	// Stack of waiting to closed paired characters.
	// It is helper structure.
	stackPairedCharacters stackChars

	// all tokens mined from the tokenizer
	tokens tokens
	// indexes that should be skipped to the parser
	skipTokenIndx []int
	// indexes referring to newline tokens
	NELIndx []int
	// number of spaces on current line (reset every new line)
	numSpacesOnCurrentLine int
	// number of new line in tokens (part of the token).
	numNotExplicitNEL int
}

func (l *lexer) releaseToken(symbol int, labels ...string) int {
	if symbol == 0 {
		// if the symbol has no alias in the parser
		symbol = int(l.data[l.ts])
	}
	newToken := &token{
		symbol: symbol,
		start:  l.ts,
		end:    l.te,
		pos:    l.currentPos(),
	}
	newToken.addLabel(labels...)
	return l.tokens.release(newToken)
}

func (l *lexer) releaseNEL(labels ...string) {
	indx := l.releaseToken(int(l.data[l.ts]), labels...)

	l.NELIndx = append(l.NELIndx, indx)

	// reset counter
	l.numSpacesOnCurrentLine = 0
}

func (l *lexer) releaseWhiteSpace(labels ...string) {
	l.numSpacesOnCurrentLine += l.te - l.ts + 1
}

func (l *lexer) currentPos() pos {
	return pos{
		l.currentLineNumber(),
		l.charNumberOnCurrentLine(),
		l.numSpacesOnCurrentLine}
}

// returns the current line number
func (l *lexer) currentLineNumber() int {
	return len(l.NELIndx) + l.numNotExplicitNEL
}

func (l *lexer) charNumberOnCurrentLine() int {
	if len(l.NELIndx) == 0 {
		return 0
	}
	lastNewLine := l.tokens[l.NELIndx[len(l.NELIndx)-1]]
	return l.ts - lastNewLine.end - 1
}

// added to stack for waiting to closed
func (l *lexer) beginPairedChar(in int) {
	l.stackPairedCharacters.push(in)
}

func (l *lexer) topPairedChar() int {
	return l.stackPairedCharacters.top()
}

func (l *lexer) isEndPairedChar(in int) bool {
	if l.top == 0 || len(l.stackPairedCharacters) == 0 {
		return false
	}
	top := l.topPairedChar()
	if top < 0 {
		return false
	}
	if top != in {
		return false
	}
	l.popWaitPairedChar()
	return true
}

func (l *lexer) popWaitPairedChar() int {
	return l.stackPairedCharacters.pop()
}

// TODO: added functions for growing stack
func (l *lexer) nostack() bool {
	if l.top != len(l.stack) {
		return false
	}
	fmt.Println("ERR: exceeds recursion limit")
	return true
}

//  Handle stack growth
func (l *lexer) stackGrowth() {
	if len(l.stack) == cap(l.stack) {
		newStack := make([]int, len(l.stack), cap(l.stack)*2)
		copy(newStack, l.stack)
		l.stack = newStack
	}
	// Append new element
	l.stack = append(l.stack, 0)
}

func (l *lexer) stackShrinking() {
	l.stack = l.stack[0 : len(l.stack)-1]
}

func (l *lexer) validAndReturn() (*lexer, error) {
	if l.p != len(l.data) {
		return l, fmt.Errorf("syntax error: p != eof (add human message)")
	}

	if l.topPairedChar() > 0 {
		return l, fmt.Errorf("syntax error: not all paired symbols were closed")
	}
	return l, nil
}
