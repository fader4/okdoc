package syntax

import (
	"fmt"
)

func NewLexer(data []byte) (lex *Lexer) {
	return &Lexer{
		Data:   data,
		Pe:     len(data),
		EOF:    len(data),
		Tokens: Tokens{},
	}
}

// Lexer helper struct for tokenizer.
//
// ragel snippet - init variables
// variable p lex.P;
// variable pe lex.Pe;
// variable cs lex.Cs;
// variable top lex.Top;
// variable stack lex.Stack;
// variable ts lex.Ts;
// variable te lex.Te;
// variable act lex.Act;
// variable eof lex.EOF;
type Lexer struct {
	// It must be an array containting the data to process.
	Data DataBytes

	// Data end pointer. This should be initialized to p plus the data length on every run of the machine. In Java and Ruby code this should be initialized to the data length.
	Pe int

	// Data pointer. In C/D code this variable is expected to be a pointer to the character data to process. It should be initialized to the beginning of the data block on every run of the machine. In Java and Ruby it is used as an offset to data and must be an integer. In this case it should be initialized to zero on every run of the machine.
	P int

	// This must be a pointer to character data. In Java and Ruby code this must be an integer. See Section 6.3 for more information.
	Ts int

	// Also a pointer to character data.
	Te int

	// This must be an integer value. It is a variable sometimes used by scanner code to keep track of the most recent successful pattern match.
	Act int

	// Current state. This must be an integer and it should persist across invocations of the machine when the data is broken into blocks that are processed independently. This variable may be modified from outside the execution loop, but not from within.
	Cs int

	// This must be an integer value and will be used as an offset to stack, giving the next available spot on the top of the stack.
	Top int

	// This must be an array of integers. It is used to store integer values representing states. If the stack must resize dynamically the Pre-push and Post-Pop statements can be used to do this (Sections 5.6 and 5.7).
	Stack []int

	EOF int

	////////////////////////////////
	// helper fields
	////////////////////////////////

	// stack (LCFS) for storing paired symbols
	// e.g. to register closing quotes and brackets
	stackPairedCharacters IntStack

	// all tokens received from the tokenizer
	Tokens Tokens

	// numCR current number of newlines
	numCR int
	// pointer to token of last newline
	lastCR *Token
	// numSPCurCR current number of spaces in the current line
	numSPCurCR int
}

// ReleaseSymbol current symbol int(l.Data[l.Ts])
func (l *Lexer) ReleaseSymbol(labels ...string) int {
	return l.ReleaseToken(int(l.Data[l.Ts]), labels...)
}

func (l *Lexer) ReleaseNewLine(labels ...string) int {
	return l.ReleaseToken(int('\n'), labels...)
}

// ReleaseToken release current slice of token with alias
func (l *Lexer) ReleaseToken(symbol int, labels ...string) int {
	newToken := &Token{
		Symbol: symbol,
		Start:  l.Ts,
		End:    l.Te,
		Pos:    l.CurrentPos(),
	}
	newToken.AddLabels(labels...)

	switch symbol {
	case '\n', '\r':
		l.numCR++
		l.numSPCurCR = 0
		l.lastCR = newToken
	case '\t', ' ':
		l.numSPCurCR++
	}

	return l.Tokens.Add(newToken)
}

// CurrentPos returns current position
func (l *Lexer) CurrentPos() Pos {
	// total number of characters on the current line
	clNumChars := 0
	if l.lastCR != nil {
		clNumChars = l.Ts - l.lastCR.End - 1
	}
	return Pos{
		l.numCR,
		clNumChars,
		l.numSPCurCR,
	}
}

// BeginPairedChar added to stack for waiting to char.
func (l *Lexer) BeginPairedChar(char int) {
	l.stackPairedCharacters.Push(char)
}

// TopPairedChar returns the top char in stack.
func (l *Lexer) TopPairedChar() int {
	return l.stackPairedCharacters.Top()
}

// IsEndPairedChar returns true if top char in stack is matched.
func (l *Lexer) IsEndPairedChar(char int) bool {
	if l.Top == 0 || len(l.stackPairedCharacters) == 0 {
		return false
	}
	topChar := l.TopPairedChar()
	if topChar < 0 {
		return false
	}
	if topChar != char {
		return false
	}
	l.stackPairedCharacters.Pop()
	return true
}

// TODO: added functions for growing stack
// func (l *Lexer) Nostack() bool {
// 	if l.Top != len(l.Stack) {
// 		return false
// 	}
// 	fmt.Println("ERR: exceeds recursion limit")
// 	return true
// }

func (l *Lexer) StackGrowth() {
	if len(l.Stack) == cap(l.Stack) {
		newStack := make([]int, len(l.Stack), cap(l.Stack)*2)
		copy(newStack, l.Stack)
		l.Stack = newStack
	}
	// Append new element
	l.Stack = append(l.Stack, 0)
}

func (l *Lexer) StackShrinking() {
	l.Stack = l.Stack[0 : len(l.Stack)-1]
}

func (l *Lexer) Valid() error {
	if l.P != len(l.Data) {
		return fmt.Errorf("syntax error: p != eof (add human message)")
	}

	if l.TopPairedChar() > 0 {
		return fmt.Errorf("syntax error: not all paired symbols were closed")
	}
	return nil
}
