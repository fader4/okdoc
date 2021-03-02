%{
package meta

import (
    "fmt"
    "github.com/fader4/okdoc/syntax"
)

%}


%union {
    token *syntax.Token
}

%token <token>
ident stringLiteral '(' ')' '='

%type <token> Content FreeTokens


%start doc
%%

doc : Content {
    fmt.Println("Init content")
};

Content: /*empty*/ {
    fmt.Println("empty")
}
|
Content FreeTokens {
    fmt.Println("FreeToken:", $2.Pos)
};

FreeTokens:
    ident {$$ = $1} | '(' {$$ = $1} | ')' {$$ = $1} | '=' {$$ = $1} | stringLiteral {$$ = $1};
