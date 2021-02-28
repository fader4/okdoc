package okdoc

import _ "fmt"

%%{

    machine py_lang;

    access lex.;
    variable p lex.p;
    variable pe lex.pe;

    WhiteSpace = [ \t]+;
    NewLine    = [\n\r];

    CommentInline = "#" [^\r\n]*;
    CommentMultiline1 = "'''" (any)* :>> "'''";
    CommentMultiline2 = '"""' (any)* :>> '"""';

    keywords = ("def"|"load"|"return");

    Int      = [0-9]+;
    Float    = (([1-9] [0-9]* [.] [0-9]*) | (0? [.] [0-9]+)) ([Ee] [+\-]? [0-9]+)?;
    Bool     = "True"|"False";
    Ident    = ([a-zA-Z] [a-zA-Z0-9_]*) - Bool - keywords;

    singleQuoteString := |*
        ['] => {
            if lex.isEndPairedChar('\'') {
                fret;
            };
        };

        ( [^'\\] | /\\./ )+ => {
            lex.releaseToken(stringLiteral)
        };
    *|;

    doubleQuoteString := |*
        ["] => {
            if lex.isEndPairedChar('"') {
                fret;
            };
        };

        ( [^"\\] | /\\./ )+ => {
            lex.releaseToken(stringLiteral)
        };
    *|;


    main := |*

        WhiteSpace;
        NewLine;

        CommentInline => {
            lex.releaseToken(commentInline)
        };
        CommentMultiline1|CommentMultiline2 => {
            lex.releaseToken(commentMultiline)
        };

        # keywords
        "load" => {
            lex.releaseToken(loadKeyword)
        };
        "def" => {
            lex.releaseToken(defKeyword)
        };
        "return" => {
            lex.releaseToken(returnKeyword)
        };
        [=:*] => {
            lex.releaseToken(int(lex.data[lex.ts]))
        };


        Ident    => {
            lex.releaseToken(identToken)
        };
        Bool     => {
            lex.releaseToken(boolLiteral)
        };
        Int      => {
            lex.releaseToken(integerLiteral)
        };
        Float    => {
            lex.releaseToken(floatLiteral)
        };

        ['] => {
            lex.beginPairedChar('\'')
            fcall singleQuoteString;
        };

        ["] => {
            lex.beginPairedChar('"')
            fcall doubleQuoteString;
        };

        "(" => {
            lex.releaseToken('(')
            lex.beginPairedChar(')');
            fcall main;
        };
        "[" => {
            lex.releaseToken('[')
            lex.beginPairedChar(']');
            fcall main;
        };
        "{" => {
            lex.releaseToken('{')
            lex.beginPairedChar('}');
            fcall main;
        };
        [}\])] => {
            if lex.isEndPairedChar(int(lex.data[lex.ts])) {
                lex.releaseToken(int(lex.data[lex.ts]))
                fret;
            };
        };

        any;
    *|;

    # The prepush statement allows the user to supply stack management code that is written out during the generation of fcall, immediately before the current state is pushed to the stack. This statement can be used to test the number of available spaces and dynamically grow the stack if necessary.
    prepush { lex.stackGrowth(); }
    # postpop { lex.stackShrinking() }
}%%

%% write data;

func pyLex(data []byte) (*lexer, error) {
    lex, eof := newLexer(data)

    %% write init;
    %% write exec;

    return lex.validAndReturn()
}
