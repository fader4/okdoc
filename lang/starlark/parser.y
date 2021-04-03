%{
package starlark

import (
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
    Assign *Assign
    Assigns []*Assign
    CallFunc *CallFunc
    Operand *Operand
    arrint []int

    val syntax.Value
    arr syntax.Array
    map_ syntax.Map
}

%token <token>
ident '(' ')' '=' ',' '{' '}' '[' ']' '.' '"' '\'' '*' ':' '+' '-' '/' '>' '<'
nullLiteral stringLiteral boolLiteral integerLiteral floatLiteral

not or and
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
%type <CallFunc> CallFunc
%type <Def> Def
%type <Operand> Operand
%type <Assign> Assign
%type <arr> LoadFields Array ArrayFields StructFields StructField CallFuncArgs ModuleField ModuleFields
%type <val> Literal LoadField CallFuncArg
%type <map_> Struct
%type <arrint> Ops

// %type <val> Literal Ident ArrayField
// %type <arr> StructField StructFields ArrayFields Array
// %type <map_> Struct


%start doc
%%

doc : Content;

Content:
    /*empty*/ {
    } | Content Comment {
        starlarklex.(*starlarkLex).Comment = $2
    } | Content Return {
        starlarklex.(*starlarkLex).Return = $2
    } | Content Load {
        starlarklex.(*starlarkLex).Load = $2
    } | Content Module {
        starlarklex.(*starlarkLex).Module = $2
    } | Content Def {
        starlarklex.(*starlarkLex).Def = $2
    };

//////////////////
// Def
//////////////////

Def:
    def ident '(' ')' ':' {
        $$ = &Def{TokenWithData: $1, Name: $2.Ident()}
    } | def ident '(' DefFields ')' ':' {
        $$ = &Def{TokenWithData: $1, Name: $2.Ident(), Fields: $4}
    };

DefFields:
    DefField {
        $$ = []*DefField{$1}
    } | DefFields ',' DefField {
        $$ = append($1, $3)
    } | DefFields ',' {
        $$ = $1
    };

DefField:
    Operand {
        $$ = &DefField{
            Key: $1,
        }
    } |
    '*' {
        $$ = &DefField{
            Varargs: true,
        }
    } |
    '*' '*' {
        $$ = &DefField{
            Kwargs: true,
        }
    } |
    '*' Operand {
        $$ = &DefField{
            Key: $2,
            Varargs: true,
        }
    } |
    '*' '*' Operand {
        $$ = &DefField{
            Key: $3,
            Kwargs: true,
        }
    } |
    Operand '=' Operand {
        $$ = &DefField{
            Key: $1,
            Value: $3,
        }
    } |
    Operand '=' Assign {
        $$ = &DefField{
            Key: $1,
            ValueExpr: $3,
        }
    }
;

//////////////////
// Comment
//////////////////

Comment:
    commentInline {
        $$ = &Comment{$1, false}
    } | commentMultiline {
        $$ = &Comment{$1, true}
    };

//////////////////
// Return
//////////////////

Return:
    returnKeyword {
        $$ = &Return{$1}
    };

//////////////////
// Load
//////////////////

Load:
    load '(' LoadFields ')' {
        $$ = &Load{Fields: $3}
    };

LoadFields:
    LoadField {
        $$ = syntax.Array{$1}
    } | LoadFields ',' LoadField {
        $$ = append($1, $3)
    } | LoadFields ',' {
        $$ = $1
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

Module:
    ident '=' module '(' stringLiteral ',' ModuleFields ')' {
        $$ = &Module{
            Export: $1.Ident(),
            Name: $5.String(),
            Fields: $7,
        }
    } | ident '=' module '(' stringLiteral ')' {
        $$ = &Module{
            Export: $1.Ident(),
            Name: $5.String(),
        }
    };

ModuleFields:
    ModuleField {
        $$ = syntax.Array{$1}
    } |
    ModuleFields ',' ModuleField {
        $$ = append($1, $3)
    } |
    ModuleFields ',' {
        $$ = $1
    };

ModuleField:
    Operand '=' Operand {
        $$ = syntax.Array{$1, $3}
    } |
    Operand '=' Assign {
        $$ = syntax.Array{$1, $3}
    };

CallFunc:
    ident '(' ')'  {
        $$ = &CallFunc{
            Name: $1.Ident(),
        }
    } |
    ident '(' CallFuncArgs ')'  {
        $$ = &CallFunc{
            Name: $1.Ident(),
            Fields: $3,
        }
    } | ident '(' '*' '*' CallFunc ')'  {
        if $5.FuncName() != "dict" {
            panic("should be only dictionary for kwargs, got " + $5.FuncName())
        }
        $$ = &CallFunc{
            Name: $1.Ident(),
            Fields: $5.Fields,
            Kwarrgs: true,
        }
    } | ident '(' '*'  '[' ArrayFields ']' ')'  {
        $$ = &CallFunc{
            Name: $1.Ident(),
            Fields: $5,
            Varrarrgs: true,
        }
    };

CallFuncArgs:
    CallFuncArgs ',' CallFuncArg {
        $$ = append($1, $3)
    } | CallFuncArg {
        $$ = syntax.Array{$1}
    } | CallFuncArgs  ',' {
        $$ = $1
    };

CallFuncArg:
    Operand {
        $$ = syntax.Array{$1, syntax.Null_{}}
    } |
    Assign {
        $$ = syntax.Array{syntax.Null_{}, $1}
    } |
    Operand '=' Operand {
        $$ = syntax.Array{$1, $3}
    } |
    Operand '=' Assign {
        $$ = syntax.Array{$1, $3}
    };

//////////////////
// Basic types
//////////////////


Literal:
    stringLiteral {
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
    } | '[' ']' {
        $$ = syntax.Array{}
    };

ArrayFields:
    Assign {
        $$ = syntax.Array{$1}
    } | Operand {
        $$ = syntax.Array{$1}
    } | ArrayFields ',' Operand {
        $1.Add($3)
        $$ = $1
    } | ArrayFields ',' Assign {
        $1.Add($3)
        $$ = $1
    } | ArrayFields ',' {
        $$ = $1
    };

Struct:
    '{' '}' {
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
                        $$.Keys = append($$.Keys, strings.Join(in, "."))
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

StructFields:
    StructField {
        $$ = syntax.Array{$1}
    } | StructFields ',' StructField {
        $1.Add($3)
        $$ = $1
    } | StructFields ',' {
        $$ = $1
    };

StructField:
    stringLiteral ':' Operand {
        $$ = syntax.Array{$1.String(), $3}
    } | stringLiteral ':' Assign {
        $$ = syntax.Array{$1.String(), $3}
    };

Ops:
    or {
        $$ = []int{$1.Symbol}
    } | and {
        $$ = []int{$1.Symbol}
    } | '+' {
        $$ = []int{$1.Symbol}
    } | '-' {
        $$ = []int{$1.Symbol}
    } | '=' '=' {
        $$ = []int{$1.Symbol, $2.Symbol}
    } | '>' {
        $$ = []int{$1.Symbol}
    } | '<' {
        $$ = []int{$1.Symbol}
    } | '<' '=' {
        $$ = []int{$1.Symbol, $2.Symbol}
    } | '>' '=' {
        $$ = []int{$1.Symbol, $2.Symbol}
    } | '*' {
        $$ = []int{$1.Symbol}
    } | '/' {
        $$ = []int{$1.Symbol}
    };

Assign:
    Operand Ops Operand {
        $$ = &Assign{Left: $1, Op: $2, Right: $3}
    } | Assign Ops Operand {
        $$ = &Assign{Parent: $1, Op: $2, Right: $3}
    };

Operand:
    ident {
        $$ = &Operand{
            Value: $1.Ident(),
        }
    } |
    Literal {
        $$ = &Operand{
            Value: $1,
        }
    } |
    Array {
        $$ = &Operand{
            Value: $1,
        }
    } |
    Struct {
        $$ = &Operand{
            Value: $1,
        }
    } |
    CallFunc {
        $$ = &Operand{
            Value: $1,
        }
    } |
    // Call Suffix
    Operand '.' CallFunc {
        $$ = &Operand{
            Parent: $1,
            Value: $3,
        }
    } |
    // Dot Suffix
    Operand '.' ident {
        $$ = &Operand{
            Parent: $1,
            Value: $3.Ident(),
        }
    };
