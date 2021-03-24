%{
package annotation

import (
    "fmt"
    "github.com/fader4/okdoc/syntax"
)

%}


%union {
    token *syntax.Token
}

%token <token>
ident '(' ')' '@' '=' ',' '{' '}'
stringLiteral boolLiteral integerLiteral floatLiteral

%type <token> Content Annotation Field Fields


%start doc
%%

doc : Content {
    fmt.Println("Init content")
};

Content: /*empty*/ {
    fmt.Println("empty")
}
|
Content '@' Annotation {
    fmt.Println("Annotation:", $2.Pos)
}
;

Annotation:
    ident {} |
    ident '(' ')' {}|
    ident '(' Fields ')' {};

Fields: Fields ',' Field | Field;

Field:
    Literal {} |
    Ident {} |
    ident '=' Ident {} |
    ident '=' Literal {} |
    ident '=' Array {}
;

Ident: ident | Ident '.' ident;

Literal: stringLiteral {} | boolLiteral {} | integerLiteral {} | floatLiteral {};

Array:
    '{' ArrayFields '}' |
    '{' '}';

ArrayFields:
    ArrayField |
    ArrayFields ',' ArrayField;

ArrayField:
    Literal {} |
    Array {} ;
