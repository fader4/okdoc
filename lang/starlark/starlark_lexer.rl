package starlark

import _ "fmt"

%%{

    machine starlark_lang;

    access lex.;
    variable p lex.p;
    variable pe lex.pe;

    action incLine {
        lex.releaseNEL()
    }
    action incLineWhiteSpaces {
        lex.releaseWhiteSpace()
    }

    WhiteSpace = [ \t]+ @incLineWhiteSpaces;
    NewLine    = [\n\r] @incLine;

    CommentInline = "#" [^\r\n]*;
    CommentMultiline1 = "'''" (any|[\n\r] @{lex.numNotExplicitNEL++})* :>> "'''";
    CommentMultiline2 = '"""' (any|[\n\r] @{lex.numNotExplicitNEL++})* :>> '"""';

    keywords = ("load"|"module"|"def");

    Int      = [0-9]+;
    Float    = (([1-9] [0-9]* [.] [0-9]*) | (0? [.] [0-9]+)) ([Ee] [+\-]? [0-9]+)?;
    Bool     = "True"|"False";
    Ident    = ([a-zA-Z_] [a-zA-Z0-9_]*) - Bool - keywords;

    singleQuoteString := |*
        ['] => {
            if lex.isEndPairedChar('\'') {
                fret;
            };
        };

        ( [^'\\] | /\\./ )+ => {
            lex.releaseToken(stringLiteral, "string")
        };
    *|;

    doubleQuoteString := |*
        ["] => {
            if lex.isEndPairedChar('"') {
                fret;
            };
        };

        ( [^"\\] | /\\./ )+ => {
            lex.releaseToken(stringLiteral, "string")
        };
    *|;


    main := |*

        WhiteSpace;
        NewLine;

        CommentInline => {
            lex.releaseToken(commentInline, "comment")
        };
        CommentMultiline1|CommentMultiline2 => {
            lex.releaseToken(commentMultiline, "comment")
        };

        # keywords
        "load" => {
            lex.releaseToken(loadKeyword, "keyword")
        };
        "module" => {
            lex.releaseToken(moduleKeyword, "keyword")
        };
        "def" => {
          lex.releaseToken(defKeyword, "keyword")
        };
        "return" => {
           lex.releaseToken(returnKeyword, "keyword")
        };
        [=] => {
            lex.releaseToken(int(lex.data[lex.ts]), "op_and_punct")
        };


        Ident    => {
            lex.releaseToken(ident, "ident")
        };
        #Bool     => {
        #    lex.releaseToken(boolLiteral)
        #};
        #Int      => {
        #    lex.releaseToken(integerLiteral)
        #};
        #Float    => {
        #    lex.releaseToken(floatLiteral)
        #};

        ['] => {
            lex.beginPairedChar('\'')
            fcall singleQuoteString;
        };

        ["] => {
            lex.beginPairedChar('"')
            fcall doubleQuoteString;
        };

        "(" => {
            lex.releaseToken('(', "bracket", "open_bracket")
            lex.beginPairedChar(')');
            fcall main;
        };
        #"[" => {
        #    lex.releaseToken('[')
        #    lex.beginPairedChar(']');
        #    fcall main;
        #};
        #"{" => {
        #    lex.releaseToken('{')
        #    lex.beginPairedChar('}');
        #    fcall main;
        #};
        # [}\])] => {
        [)] => {
            if lex.isEndPairedChar(int(lex.data[lex.ts])) {
                lex.releaseToken(int(lex.data[lex.ts]), "bracket", "close_bracket")
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

func starlarkLex(data []byte) (*lexer, error) {
    lex, eof := newLexer(data)

    %% write init;
    %% write exec;

    return lex.validAndReturn()
}
