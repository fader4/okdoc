%{
package okdoc

import (
    "fmt"
)

%}


%union {
    token *tokenWithData
    node *node
}

%type <node>
Content Comment
LoadStmt


%token <token>
commentInline commentMultiline
ident boolLiteral integerLiteral floatLiteral stringLiteral
defKeyword loadKeyword returnKeyword moduleKeyword
'.' ',' ':' '(' ')' '[' ']' '{' '}' '=' '*'

/*
    p := yylex.(*Parser)
*/

%start doc
%%

doc : Content;

Content: /*empty*/ {
}
|
Content Comment {
    fmt.Println("Comment")
}
|
Content DefStmt {
    // fmt.Println("DefStmt")
}
|
Content ReturnStmt {
    fmt.Println("ReturnStmt")
}
|
Content LoadStmt {
    fmt.Println("LoadStmt")
}
|
Content ModuleStmt {
    // fmt.Println("ModuleStmt")
}
|
Content stringLiteral {
    fmt.Println("stringLiteral")
}
|
Content FreeTokens;


Comment: commentInline {
} | commentMultiline {
};


LoadStmt:
    loadKeyword '(' Arguments ')' {
    };

ModuleStmt:
    moduleKeyword '(' stringLiteral Arguments ')' {
        fmt.Println("Named module:", $3)
    };

DefStmt:
    defKeyword ident {
        fmt.Println("Function:", $2)
    };

ReturnStmt:
    returnKeyword {
    };

Argument:
    stringLiteral |
    ident |
    Argument '=';

Arguments:
    Argument |
    Arguments Argument;

FreeTokens:
    ident | '(' | ')' | '=';
