%{
package okdoc

import (
    "fmt"
)

%}


%union {
    token string
    node *Node
    content *Content

    // CR counter
    totallines uint
}

%type <content>
Content

%type <node>
Ident Literal Comment DefStatement WhiteSpace /*OperAndPunct*/

%token <token>
whiteSpace
commentInline commentMultiline
identToken boolLiteral integerLiteral floatLiteral stringLiteral
defKeyword
'(' ')' ',' '=' '{' '}'


/*
    p := yylex.(*Parser)
*/

%start doc
%%

doc : Content;

Content: /*empty*/ {
    yylex.(*Parser).Content = &Content{}
    $$ = yylex.(*Parser).Content
}
|
Content Literal {
    $1.Nodes = append($1.Nodes, $2)
}
|
Content Ident {
    $1.Nodes = append($1.Nodes, $2)
}
|
Content Comment {
    $1.Nodes = append($1.Nodes, $2)
}
|
Content DefStatement {
    $1.Nodes = append($1.Nodes, $2)
}
|
Content WhiteSpace {
    $1.Nodes = append($1.Nodes, $2)
}
// |
// Content OperAndPunct {
//     $1.Nodes = append($1.Nodes, $2)
// }
;

DefStatement:
defKeyword WhiteSpace Ident {
    fmt.Println("def")
    $$ = &Node{Type: "defStatement", Raw: $3.Raw, Line: yyVAL.totallines}
}
;


Ident: identToken {
    $$ = &Node{Type: "ident", Raw: $1, Line: yyVAL.totallines}
};

Literal:
stringLiteral {
    $$ = &Node{Type: "string", Raw: $1, Line: yyVAL.totallines}
  }
| integerLiteral {
    $$ = &Node{Type: "integer", Raw: $1, Line: yyVAL.totallines}
 }
| floatLiteral {
    $$ = &Node{Type: "float", Raw: $1, Line: yyVAL.totallines}
 }
| boolLiteral {
    $$ = &Node{Type: "bool", Raw: $1, Line: yyVAL.totallines}
 }
;

Comment: commentInline {
    $$ = &Node{Type: "commentInline", Raw: $1, Line: yyVAL.totallines}
} | commentMultiline {
    $$ = &Node{Type: "commentMultiline", Raw: $1, Line: yyVAL.totallines}
};

WhiteSpace: whiteSpace {
    $$ = &Node{Type: "whiteSpace", Raw: $1, Line: yyVAL.totallines}
};

// OperAndPunct:
// '(' {
//     $$ = &Node{Type: "OperAndPunct"}
// }
// |
// ')' {
//     $$ = &Node{Type: "OperAndPunct"}
// }
// |
// '=' {
//     $$ = &Node{Type: "OperAndPunct"}
// }
// |
// '{' {
//     $$ = &Node{Type: "OperAndPunct"}
// }
// |
// '}' {
//     $$ = &Node{Type: "OperAndPunct"}
// }
// |
// ',' {
//     $$ = &Node{Type: "OperAndPunct"}
// }
// ;
