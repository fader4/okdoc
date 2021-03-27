package annotation

import (
    "github.com/fader4/okdoc/syntax"
    _ "fmt"
)

%%{

    machine annotation;

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
    # CommentInline = "#" [^\r\n]*;
    # CommentMultiline1 = "'''" (any|[\n\r] @{lex.ReleaseNewLine("CR", "CR_nested")})* :>> "'''";
    # CommentMultiline2 = '"""' (any|[\n\r] @{lex.ReleaseNewLine("CR", "CR_nested")})* :>> '"""';

    Int      = [0-9]+;
    Float    = (([1-9] [0-9]* [.] [0-9]*) | (0? [.] [0-9]+)) ([Ee] [+\-]? [0-9]+)?;
    Null     = "Null";
    Bool     = "True"|"False";
    Ident    = ([a-zA-Z_] [a-zA-Z0-9_]*) - Bool - AT - Null;

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
        [=,.] => {
            lex.ReleaseSymbol("op_and_punct")
        };

        Null => {
            lex.ReleaseToken(nullLiteral, "literal")
        };
        Ident    => {
            lex.ReleaseToken(ident, "ident")
        };
        Bool     => {
            lex.ReleaseToken(boolLiteral, "literal", "bool")
        };
        Int      => {
            lex.ReleaseToken(integerLiteral, "literal", "int")
        };
        Float    => {
            lex.ReleaseToken(floatLiteral, "literal", "float")
        };

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
        "{" => {
            lex.ReleaseToken('{', "bracket", "open_bracket")
            lex.BeginPairedChar('}');
            fcall main;
        };
        "[" => {
            lex.ReleaseToken('[', "bracket", "open_bracket")
            lex.BeginPairedChar(']');
            fcall main;
        };
        [)\]}] => {
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
