package okdoc

import "fmt"

%%{

    machine okdoc;

    write data;

    # The prepush statement allows the user to supply stack management code that is written out during the generation of fcall, immediately before the current state is pushed to the stack. This statement can be used to test the number of available spaces and dynamically grow the stack if necessary.
    prepush { if lex.nostack() { return foundToken }; }

    access lex.;
    variable p lex.p;
    variable pe lex.pe;

    action lineCounter { out.totallines++}

    WhiteSpace = [ \t]+;
    NewLine    = [\n\r] @lineCounter;
    CommentInline = "#" [^\r\n]*;
    CommentMultiline1 = "'''" (any - NewLine | NewLine)* :>> "'''";
    CommentMultiline2 = '"""' (any - NewLine | NewLine)* :>> '"""';

    keywords = ("def"|"if"|"while"|"return"|"else");

    Int      = [0-9]+;
    Float    = (([1-9] [0-9]* [.] [0-9]*) | (0? [.] [0-9]+)) ([Ee] [+\-]? [0-9]+)?;
    Bool     = "True"|"False";
    Ident    = ([a-zA-Z] [a-zA-Z0-9_]*) - Bool - keywords;

    string := |*
        ['] => {
            if !lex.isEndBoundary('\'') {
                fhold;
            };
            fret;
        };

        ( [^'\\] | /\\./ )+ => {
            releaseToken(stringLiteral, "stringLiteral for single quotes")
            fbreak;
        };
    *|;

    rstring := |*
        ["] => {
            if !lex.isEndBoundary('"') {
                fhold;
            };
            fret;
        };

        ( [^"\\] | /\\./ )+ => {
            releaseToken(stringLiteral, "stringLiteral for double quotes")
            fbreak;
        };
    *|;

    # DefStmt = 'def' identifier '(' [Parameters [',']] ')' ':' Suite .
    # Parameters = Parameter {',' Parameter}.
    # Parameter = identifier | identifier '=' Test | '*' | '*' identifier | '**' identifier .
    # DefStmt = WhiteSpace* "def" WhiteSpace+ Ident WhiteSpace* "(" [^\r\n]*  ")" WhiteSpace* :>> ":";

}%%



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

    // stack of waiting boundaries to close
    stackBoundary []int

    // This must be an array of integers. It is used to store integer values representing states. If the stack must resize dynamically the Pre-push and Post-Pop statements can be used to do this (Sections 5.6 and 5.7).
    stack [64]int
}

func newLexer(data []byte) *lexer {
    lex := &lexer{
        data: data,
        pe: len(data),
    }
    %% write init;
    return lex
}


func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe

    // found token for parser
    foundToken  := 0

    releaseToken := func(tok int, humanName string) {
        if tok == -1 {
            tok = int(lex.data[lex.ts])
        }
        foundToken = tok
        out.token = string(lex.data[lex.ts:lex.te])

        if yyDebug == 1 {
            fmt.Printf("lexer: Release %q: %q \n", humanName, string(lex.data[lex.ts:lex.te]))
        }
    }

    %%{
        expr := |*

            WhiteSpace => {
                releaseToken(whiteSpace, "whiteSpace")
                fbreak;
            };
            NewLine;

            "def" => {
                releaseToken(defKeyword, "defKeyword")
                fbreak;
            };

            CommentInline => {
                releaseToken(commentInline, "commentInline")
                fbreak;
            };
            CommentMultiline1|CommentMultiline2 => {
                releaseToken(commentMultiline, "commentMultiline")
                fbreak;
            };

            Ident    => {
                releaseToken(identToken, "identToken")
                fbreak;
            };
            Bool     => {
                releaseToken(boolLiteral, "boolLiteral")
                fbreak;
            };
            Int      => {
                releaseToken(integerLiteral, "integerLiteral")
                fbreak;
            };
            Float    => {
                releaseToken(floatLiteral, "floatLiteral")
                fbreak;
            };

            ['] => { lex.beginBoundary('\''); fcall string; };
            ["] => { lex.beginBoundary('"'); fcall rstring; };
            # [()=,{}] => {
            #    releaseToken(-1, "operators and punctuation")
            #    fbreak;
            # };

            any;

        *|;

         write exec;
    }%%

    return foundToken;
}

func (lex *lexer) beginBoundary(fin int) {
    // Push
    lex.stackBoundary = append(lex.stackBoundary, fin)
}

func (lex *lexer) endBoundary() int{
    // Pop
    top := len(lex.stackBoundary) - 1
    last := lex.stackBoundary[top]
    lex.stackBoundary = lex.stackBoundary[:top]

    return last
}

func (lex *lexer) isEndBoundary(sym int) bool {
    if lex.top == 0 || len(lex.stackBoundary) == 0 {
        fmt.Println("ERR: does not close anything")
        return false
    }
    prevsym := lex.endBoundary()
    if prevsym != sym {
        fmt.Println("ERR: does not close")
        return false
    }
    return true
}

func (lex *lexer) nostack() bool {
    if lex.top != len(lex.stack) {
        return false
    }
    fmt.Println("ERR: exceeds recursion limit")
    return true
}



func (lex *lexer) Error(e string) {
    fmt.Println("lexer: Error:", e)
}
