%{
package starlark

import (
    "fmt"
    "log"
    "github.com/fader4/okdoc/syntax"
    "strings"
)

%}


%union {
    token *syntax.TokenWithData


    Comment *Comment
    Return *Return
    Module *Module
    Load *Load
    Def *Def
    DefField *DefField
    DefFields []*DefField
    DictField *DictField
    DictFields []*DictField
    CallFunc *CallFunc

    val syntax.Value
    arr syntax.Array
    map_ syntax.Map
}

%token <token>
ident '(' ')' '=' ',' '{' '}' '[' ']' '.' '"' '\'' '*' ':'
nullLiteral stringLiteral boolLiteral integerLiteral floatLiteral

dict
returnKeyword
commentInline
def endDef
load endLoad
module endModule
commentMultiline endCommentMultiline


%type <Comment> Comment
%type <Return> Return
%type <Load> Load
%type <Module> Module
%type <DefField> DefField
%type <DefFields> DefFields
%type <DictField> DictField
%type <DictFields> Dict DictFields ArbitraryNamedArgs
%type <CallFunc> CallFunc
%type <Def> Def
%type <arr> LoadFields Array ArrayFields StructFields StructField ModuleFields CallFuncArgs ArbitraryArrArgs
%type <val> Literal LoadField Ident ArrayField CallFuncArg
%type <map_> Struct

// %type <val> Literal Ident ArrayField
// %type <arr> StructField StructFields ArrayFields Array
// %type <map_> Struct


%start doc
%%

doc : Content;

Content: /*empty*/ {
} | Content Comment {
    fmt.Printf("Comment %T: %q\n", $2, string($2.MustBytes()))
} | Content Return {
    fmt.Printf("Return %T: %q\n", $2, string($2.MustBytes()))
} | Content Load {
    fmt.Printf("Load %T: %d\n", $2, len($2.Fields))
} | Content Module {
    fmt.Printf("Module %T: %d\n", $2, len($2.Fields))
} | Content Def {
    fmt.Printf("Def %T: %d\n", $2, len($2.Fields))
};

//////////////////
// Def
//////////////////

Def: def ident '(' ')' ':' {
    $$ = &Def{TokenWithData: $1, Name: $2.Ident()}
} | def ident '(' DefFields ')' ':' {
    $$ = &Def{TokenWithData: $1, Name: $2.Ident(), Fields: $4}
};

DefFields: DefField {
    $$ = []*DefField{$1}
} | DefFields ',' DefField {
    $$ = append($1, $3)
};

DefField:
    Literal {
        $$ = &DefField{
            Value: $1,
        }
    } |
    Struct {
        $$ = &DefField{
            Value: $1,
        }
    } |
    Array {
        $$ = &DefField{
            Value: $1,
        }
    } |
    Ident {
        $$ = &DefField{
            Value: $1,
        }
    } |
    '*' Ident {
        $$ = &DefField{
            Value: $1,
            Varargs: true,
        }
    } |
    '*' '*' Ident {
        $$ = &DefField{
            Value: $1,
            Kwargs: true,
        }
    } |
    Ident '=' Ident {
        $$ = &DefField{
            Key: $1,
            Value: $3,
        }
    } |
    Ident '=' Literal {
        $$ = &DefField{
            Key: $1,
            Value: $3,
        }
    } |
    Ident '=' Array {
        $$ = &DefField{
            Key: $1,
            Value: $3,
        }
    } |
    Ident '=' Struct {
        $$ = &DefField{
            Key: $1,
            Value: $3,
        }
    } |
    Ident '=' Dict {
        $$ = &DefField{
            Key: $1,
            Value: $3,
        }
    }
;

//////////////////
// Comment
//////////////////

Comment: commentInline {
    $$ = &Comment{$1, false}
} | commentMultiline {
    $$ = &Comment{$1, true}
};

//////////////////
// Return
//////////////////

Return: returnKeyword {
    $$ = &Return{$1}
};

//////////////////
// Load
//////////////////

Load: load '(' LoadFields ')' {
    $$ = &Load{Fields: $3}
};

LoadFields: LoadField {
    $$ = syntax.Array{$1}
} | LoadFields ',' LoadField {
    $$ = append($1, $3)
};

LoadField:
    stringLiteral {
        $$ = syntax.Array{syntax.Null_{}, $1.String()}
    } |
    ident '=' stringLiteral {
        $$ = syntax.Array{$1.Ident(), $3.String()}
    };

//////////////////
// Module
//////////////////

Module: ident '=' module '(' stringLiteral ',' ModuleFields ')' {
    $$ = &Module{
        Ident: $1.Ident(),
        Name: $5.String(),
        Fields: $7,
    }
} | ident '=' module '(' stringLiteral ')' {
    $$ = &Module{
        Ident: $1.Ident(),
        Name: $5.String(),
    }
};

ModuleFields: ModuleFields ',' DictField {
    $$ = append($$, $3)
} | DictField {
    $$ = syntax.Array{$1}
};

CallFunc: ident '(' CallFuncArgs ')'  {
    $$ = &CallFunc{
        Name: $1.Ident(),
        Fields: $3,
    }
} | ident '(' ArbitraryArrArgs ')' {
    $$ = &CallFunc{
        Name: $1.Ident(),
        ArbitraryArrArgs: $3,
    }
} | ident '(' ArbitraryNamedArgs ')' {
    $$ = &CallFunc{
        Name: $1.Ident(),
        ArbitraryNamedArgs: $3,
    }
};

CallFuncArgs:
    CallFuncArgs ',' CallFuncArg {
        $$ = append($1, $3)
    } | CallFuncArg {
        $$ = syntax.Array{$1}
    };

CallFuncArg: DictField {
    $$ = $1
} | ArrayField {
    $$ = $1
};

ArbitraryArrArgs: '*' '*' '[' ArrayFields ']' {
    $$ = $4
};
ArbitraryNamedArgs: '*' '*' Dict {
    $$ = $3
};

//////////////////
// Basic types
//////////////////

Ident: ident {
    $$ = $1.Ident()
} | Ident '.' ident {
    $$ = $1.(syntax.Ident_).Append($3.Ident())
};

Literal: stringLiteral {
    $$ = $1.String()
} | boolLiteral {
    $$ = $1.Bool()
} | integerLiteral {
    $$ = $1.Int()
} | floatLiteral {
    $$ = $1.Float()
} | nullLiteral {
    $$ = $1.Null()
};

Array:
'(' ArrayFields ')' {
    // TODO: it is tuple
    $$ = $2
} | '[' ArrayFields ']' {
    // TODO: it is list
    $$ = $2
} | '(' ')' {
    $$ = syntax.Array{}
}   | '[' ']' {
    $$ = syntax.Array{}
};

ArrayFields: ArrayField {
    $$ = syntax.Array{$1}
} | ArrayFields ',' ArrayField {
    $1.Add($3)
    $$ = $1
};

ArrayField: Literal {
    $$ = $1
} | Array {
    $$ = $1
} | Struct {
    $$ = $1
} | Ident {
    $$ = $1
} | CallFunc {
    $$ = $1
} | Dict {
    $$ = $1
};

Struct: '{' '}' {
    $$ = syntax.Map{}
} | '{' StructFields '}' {
    $$ = syntax.Map{}
    for _, item := range $2 {
        key := item.(syntax.Array)[0]
        switch in := key.(type) {
            case syntax.StringLiteral:
                $$.Keys = append($$.Keys, string(in))
            case syntax.Ident_:
                if len(in) > 0 {
                    $$.Keys = append($$.Keys, "@"+strings.Join(in, "."))
                } else {
                    $$.Keys = append($$.Keys, "")
                }
            case nil:
                $$.Keys = append($$.Keys, "")
            default:
                log.Printf("starlark#Fields: not supported key type %T\n", key)
        }
        $$.Values.Add(item.(syntax.Array)[1])
    }
};

StructFields: StructField {
    $$ = syntax.Array{$1}
} | StructFields ',' StructField {
    $1.Add($3)
    $$ = $1
};

StructField: stringLiteral ':' Literal {
    $$ = syntax.Array{$1.String(), $3}
} | Ident ':' Array {
    $$ = syntax.Array{$1, $3}
} | Ident ':' Ident {
    $$ = syntax.Array{$1, $3}
} | Ident ':' Struct {
    $$ = syntax.Array{$1, $3}
} | Ident ':' CallFunc {
    $$ = syntax.Array{$1, $3}
}  Ident ':' Dict {
    $$ = syntax.Array{$1, $3}
};

Dict: dict '(' DictFields ')' {
    $$ = $3
};

DictFields: DictFields ',' DictField {
    $$ = append($1, $3)
} | DictField {
    $$ = []*DictField{$1}
};

DictField:
    Ident '=' Ident {
        $$ = &DictField{Key: $1, Value: $3}
    } |
    Ident '=' Literal {
        $$ = &DictField{Key: $1, Value: $3}
    } |
    Ident '=' Array {
        $$ = &DictField{Key: $1, Value: $3}
    } |
    Ident '=' Struct {
        $$ = &DictField{Key: $1, Value: $3}
    } |
    Ident '=' CallFunc {
        $$ = &DictField{Key: $1, ValueCallFn: $3}
    } |
    Ident '=' Dict {
        $$ = &DictField{Key: $1, ValueDict: $3}
    };
