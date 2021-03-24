package annotation

import (
    "github.com/fader4/okdoc/syntax"
    _ "fmt"
)

%%{

    machine annotation_preprocessing;

    variable p lex.P;
    variable pe lex.Pe;
    variable cs lex.Cs;
    variable top lex.Top;
    variable stack lex.Stack;
    variable ts lex.Ts;
    variable te lex.Te;
    variable act lex.Act;
    variable eof lex.EOF;

    action incLine {
        lex.ReleaseSymbol("CR")
    }
    action incLineWhiteSpaces {
        lex.ReleaseSymbol("SP")
    }

    WhiteSpace = [ \t] @incLineWhiteSpaces;
    NewLine    = [\n\r] @incLine;

    AT = "@";
    Bool     = "True"|"False";
    Ident    = ([a-zA-Z_] [a-zA-Z0-9_]*) - Bool;

    Annotation1 = AT Ident NewLine;
    Annotation2 = AT Ident "(" (any|[\n\r] @{lex.ReleaseNewLine("CR", "CR_nested")})* :>> ")";

    singleQuoteString := |*
        ['] => {
            if lex.IsEndPairedChar('\'') {
                fret;
            };
        };

        ( [^'\\] | /\\./ )+ => {
            lex.ReleaseToken(stringLiteral, "literal", "string")
        };
    *|;

    doubleQuoteString := |*
        ["] => {
            if lex.IsEndPairedChar('"') {
                fret;
            };
        };

        ( [^"\\] | /\\./ )+ => {
            lex.ReleaseToken(stringLiteral, "literal", "string")
        };
    *|;


    main := |*

        WhiteSpace;
        NewLine;

        AT => {
            lex.ReleaseSymbol("at")
        };


        Ident    => {
            lex.ReleaseToken(ident, "ident")
        };

        Annotation1 => {
            lex.ReleaseToken(1, "annotation")
        };

        Annotation2 => {
            lex.ReleaseToken(2, "annotation")
        };


        "(" => {
            lex.ReleaseToken('(', "bracket", "open_bracket")
            lex.BeginPairedChar(')');
            fcall main;
        };
        [)] => {
            if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
                lex.ReleaseSymbol("bracket", "close_bracket")
                fret;
            };
        };

        any;
    *|;

    # The prepush statement allows the user to supply stack management code that is written out during the generation of fcall, immediately before the current state is pushed to the stack. This statement can be used to test the number of available spaces and dynamically grow the stack if necessary.
    prepush { lex.StackGrowth(); }
    # postpop { lex.StackShrinking() }
}%%

%% write data;

func newPreprocessing(data []byte) (*syntax.Lexer, error) {
    lex := syntax.NewLexer(data)

    %% write init;
    %% write exec;

    if err := lex.Valid(); err != nil {
        return lex, err
    }

    return lex, nil
}
