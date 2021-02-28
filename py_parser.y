%{
package okdoc

import (
    "fmt"
)

%}


%union {
    token int
    node *node
}

%type <node>
Content Comment
Ident Literal
LoadStmt


%token <token>
commentInline commentMultiline
identToken boolLiteral integerLiteral floatLiteral stringLiteral
defKeyword loadKeyword returnKeyword
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
Content Literal {
    fmt.Println("Literal")
}
|
Content Ident {
    fmt.Println("Ident")
}
|
Content LoadStmt {
    fmt.Println("LoadStmt")
}
|
Content DefStmt {
    fmt.Println("DefStmt")
}
|
Content ReturnStmt {
    fmt.Println("ReturnStmt")
}
;



Comment: commentInline {
    fmt.Println("commentInline")
} | commentMultiline {
    fmt.Println("commentMultiline")
};


Ident: identToken {
};

Literal:
stringLiteral {
    fmt.Println("stringLiteral")
  }
| integerLiteral {
    fmt.Println("integerLiteral")
 }
| floatLiteral {
    fmt.Println("floatLiteral")
 }
| boolLiteral {
    fmt.Println("boolLiteral")
 }
;

LoadStmt:
    loadKeyword '(' Arguments ')' {
    };

Argument:
    stringLiteral |
    Ident '=' stringLiteral;

Arguments:
    Argument |
    Argument ',' |
    Arguments Argument;

DefStmt:
    defKeyword Ident '(' Parameters ')' ':' |
    defKeyword Ident '(' ')' ':'
;

Parameters:
    Parameter |
    Parameter ',' |
    Parameters Parameter
;

Parameter:
    Ident |
    Ident '=' Literal |
    '*'
;

ReturnStmt: returnKeyword;
