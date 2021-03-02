package meta

import (
    "github.com/fader4/okdoc/syntax"
    _ "fmt"
)

%%{

    machine meta;

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
    CommentMultiline1 = "'''" (any|[\n\r] @{lex.ReleaseNewLine("CR", "CR_nested")})* :>> "'''";
    CommentMultiline2 = '"""' (any|[\n\r] @{lex.ReleaseNewLine("CR", "CR_nested")})* :>> '"""';

    # Int      = [0-9]+;
    # Float    = (([1-9] [0-9]* [.] [0-9]*) | (0? [.] [0-9]+)) ([Ee] [+\-]? [0-9]+)?;
    Bool     = "True"|"False";
    Ident    = ([a-zA-Z_] [a-zA-Z0-9_]*) - Bool;

    singleQuoteString := |*
        ['] => {
            if lex.IsEndPairedChar('\'') {
                fret;
            };
        };

        ( [^'\\] | /\\./ )+ => {
            lex.ReleaseToken(stringLiteral, "string")
        };
    *|;

    doubleQuoteString := |*
        ["] => {
            if lex.IsEndPairedChar('"') {
                fret;
            };
        };

        ( [^"\\] | /\\./ )+ => {
            lex.ReleaseToken(stringLiteral, "string")
        };
    *|;


    main := |*

        WhiteSpace;
        NewLine;

        # CommentInline => {
        #    lex.ReleaseToken(commentInline, "comment")
        #};
        #CommentMultiline1|CommentMultiline2 => {
        #    lex.ReleaseToken(commentMultiline, "comment")
        #};

        [=] => {
            lex.ReleaseSymbol("op_and_punct")
        };


        Ident    => {
            lex.ReleaseToken(ident, "ident")
        };
        #Bool     => {
        #    lex.ReleaseToken(boolLiteral)
        #};
        #Int      => {
        #    lex.ReleaseToken(integerLiteral)
        #};
        #Float    => {
        #    lex.ReleaseToken(floatLiteral)
        #};

        ['] => {
            lex.BeginPairedChar('\'')
            fcall singleQuoteString;
        };

        ["] => {
            lex.BeginPairedChar('"')
            fcall doubleQuoteString;
        };

        "(" => {
            lex.ReleaseToken('(', "bracket", "open_bracket")
            lex.BeginPairedChar(')');
            fcall main;
        };
        #"[" => {
        #    lex.ReleaseToken('[')
        #    lex.BeginPairedChar(']');
        #    fcall main;
        #};
        #"{" => {
        #    lex.ReleaseToken('{')
        #    lex.BeginPairedChar('}');
        #    fcall main;
        #};
        # [}\])] => {
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

func newTokenizer(data []byte) (*syntax.Lexer, error) {
    lex := syntax.NewLexer(data)

    %% write init;
    %% write exec;

    if err := lex.Valid(); err != nil {
        return lex, err
    }

    return lex, nil
}
