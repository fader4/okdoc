package starlark

import (
    "github.com/fader4/okdoc/syntax"
)

%%{

    machine starlark_extractor;

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

    CommentInline = "#" [^\r\n]*;

    defKeyword = "def";
    moduleKeyword = "module";
    loadKeyword = "load";
    returnKeyword = "return";
    keywords = defKeyword | moduleKeyword | loadKeyword | returnKeyword;

    Ident    = ([a-zA-Z_] [a-zA-Z0-9_]*);

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

    defMethodBody := |*
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
        "):" => {
            if lex.IsEndPairedChar(endDef) {
                lex.ReleaseToken(endDef, "def")
                fret;
            }
        };
        ")" => {
            if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
            };
        };

        any;
    *|;

    loadBody := |*
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
            if lex.IsEndPairedChar(endLoad) {
                lex.ReleaseToken(endLoad, "load")
                fret;
            }
            if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
            };
        };

        any;
    *|;

    moduleBody := |*
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
            if lex.IsEndPairedChar(endModule) {
                lex.ReleaseToken(endModule, "module")
                fret;
            }
            if lex.IsEndPairedChar(int(lex.Data[lex.Ts])) {
            };
        };

        any;
    *|;

    commentMultilineBody := |*
        NewLine;
        "'''"|'"""' => {
            if lex.IsEndPairedChar(endCommentMultiline) {
                lex.ReleaseToken(endCommentMultiline, "comment")
                fret;
            }
        };
        any;
    *|;


    main := |*

        WhiteSpace;
        NewLine;

        CommentInline => {
            lex.ReleaseToken(commentInline, "comment")
        };
        defKeyword WhiteSpace* Ident "(" => {
            lex.ReleaseToken(def, "def")
            lex.BeginPairedChar(endDef)
            fcall defMethodBody;
        };
        loadKeyword WhiteSpace* "(" => {
            lex.ReleaseToken(load, "load")
            lex.BeginPairedChar(endLoad)
            fcall loadBody;
        };
        Ident WhiteSpace* '=' WhiteSpace* moduleKeyword WhiteSpace* "(" => {
            lex.ReleaseToken(module, "module")
            lex.BeginPairedChar(endModule)
            fcall moduleBody;
        };
        "'''"|'"""' => {
            lex.ReleaseToken(commentMultiline, "comment")
            lex.BeginPairedChar(endCommentMultiline)
            fcall commentMultilineBody;
        };
        returnKeyword => {
            lex.ReleaseToken(returnKeyword, "return")
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
