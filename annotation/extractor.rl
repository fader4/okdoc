package annotation

import (
    "github.com/fader4/okdoc/syntax"
    _ "fmt"
)

%%{

    machine annotation_extractor;

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
    Ident    = ([A-Z] [a-zA-Z0-9_]*) - Bool;


    singleQuoteString := |*
        ['] => {
            if lex.IsEndPairedChar('\'') {
                fret;
            };
        };

        ( [^'\\] | /\\./ | NewLine )+;
    *|;

    doubleQuoteString := |*
        ["] => {
            if lex.IsEndPairedChar('"') {
                fret;
            };
        };

        ( [^"\\] | /\\./ | NewLine )+;
    *|;

    annotationBody := |*
        NewLine;

        ['] => {
            lex.BeginPairedChar('\'')
            fcall singleQuoteString;
        };

        ["] => {
            lex.BeginPairedChar('"')
            fcall doubleQuoteString;
        };


        "(" => {
           lex.BeginPairedChar(')');
        };
        ")" => {
            if lex.IsEndPairedChar(endAnnotation) {
                lex.ReleaseToken(endAnnotation, "annotation")
                fret;
            }
           if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
           };
        };

        any;
    *|;


    main := |*

        WhiteSpace;
        NewLine;


        AT Ident [^(] => {
            lex.ReleaseToken(beginAnnotation, "annotation")
            lex.ReleaseToken(endAnnotation, "annotation")
        };
        AT Ident "(" => {
            lex.ReleaseToken(beginAnnotation, "annotation")
            lex.BeginPairedChar(endAnnotation)
            fcall annotationBody;
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
