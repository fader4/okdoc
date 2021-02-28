package okdoc

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

	// helper

	// Stack of waiting to closed paired characters.
	// It is helper structure.
	stackPairedCharacters stackChars

	tokens                  tokens
	idxNewlines             []int
	countLineWhiteSpaces    int
	countNewLinesInComments int
}

func (l *lexer) releaseToken(symbol int) {
	if symbol == 0 {
		// if the symbol has no alias in the parser
		symbol = int(l.data[l.ts])
	}
	l.tokens.release(&token{
		symbol: symbol,
		start:  l.ts,
		end:    l.te,
		pos:    l.currentPos(),
	})
}

func (l *lexer) releaseNEL() {
	idx := l.tokens.release(&token{
		symbol:     int(l.data[l.ts]),
		start:      l.ts,
		end:        l.te,
		pos:        l.currentPos(),
		skipParser: true,
	})
	l.idxNewlines = append(l.idxNewlines, idx)
	l.countLineWhiteSpaces = 0
}

func (l *lexer) currentPos() [3]int {
	return [3]int{
		l.lineNumber() + l.countNewLinesInComments,
		l.charNumberOnLine(),
		l.countLineWhiteSpaces}
}

func (l *lexer) lineNumber() int {
	return len(l.idxNewlines)
}

func (l *lexer) releaseWhiteSpace() {
	l.countLineWhiteSpaces += l.te - l.ts + 1
}

func (l *lexer) charNumberOnLine() int {
	if len(l.idxNewlines) == 0 {
		return 0
	}
	lastNewLine := l.tokens[l.idxNewlines[len(l.idxNewlines)-1]]
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

type stackChars []int

func (s *stackChars) push(in int) {
	*s = append(*s, in)
}

func (s stackChars) top() int {
	idx := len(s) - 1
	if idx < 0 {
		return -1
	}
	return s[idx]
}

func (s *stackChars) pop() int {
	idx := len(*s) - 1
	if idx < 0 {
		return -1
	}

	top := (*s)[idx]
	*s = (*s)[:idx]

	return top
}
