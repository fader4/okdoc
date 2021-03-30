%{
package starlark

import (
    "fmt"
)

%}


%union {
    token *Token


    // val Value
    // arr Array
    // map_ Map
}

%token <token>
ident '(' ')' '=' ',' '{' '}' '[' ']' '.'
nullLiteral stringLiteral boolLiteral integerLiteral floatLiteral

returnKeyword

commentInline
def endDef
load endLoad
module endModule
commentMultiline endCommentMultiline


%type <token> Comment
// %type <val> Literal Ident ArrayField
// %type <arr> StructField StructFields ArrayFields Array
// %type <map_> Struct


%start doc
%%

doc : Content;

Content: /*empty*/ {
} | Content Comment {
    fmt.Println("Comment", $2)
};

Comment: commentInline {
    $$ = $1
} | commentMultiline {
    $$ = $1
};

// Fields: Fields ',' Field {
//     $$ = append($1, $3)
// } | Field {
//     $$ = []*AnnotationField{$1}
// };

// Field:
//     Literal {
//         $$ = &AnnotationField{
//             Value: $1,
//         }
//     } |
//     Struct {
//         $$ = &AnnotationField{
//             Value: $1,
//         }
//     } |
//     Array {
//         $$ = &AnnotationField{
//             Value: $1,
//         }
//     } |
//     Ident {
//         $$ = &AnnotationField{
//             Value: $1,
//         }
//     } |
//     Ident '=' Ident {
//         $$ = &AnnotationField{
//             Value: $1,
//         }
//     } |
//     Ident '=' Literal {
//         $$ = &AnnotationField{
//             Key: $1,
//             Value: $3,
//         }
//     } |
//     Ident '=' Array {
//         $$ = &AnnotationField{
//             Key: $1,
//             Value: $3,
//         }
//     } |
//     Ident '=' Struct {
//         $$ = &AnnotationField{
//             Key: $1,
//             Value: $3,
//         }
//     }
// ;

// Ident: ident {
//     $$ = $1.Ident()
// } | Ident '.' ident {
//     $$ = $1.(Ident_).Append($3.Ident())
// };

// Literal: stringLiteral {
//     $$ = $1.String()
// } | boolLiteral {
//     $$ = $1.Bool()
// } | integerLiteral {
//     $$ = $1.Int()
// } | floatLiteral {
//     $$ = $1.Float()
// } | nullLiteral {
//     $$ = $1.Null()
// };

// Array: '[' ArrayFields ']' {
//     $$ = $2
// } | '[' ']' {
//     $$ = Array{}
// };

// ArrayFields: ArrayField {
//     $$ = Array{$1}
// } | ArrayFields ',' ArrayField {
//     $1.Add($3)
//     $$ = $1
// };

// ArrayField: Literal {
//     $$ = $1
// } | Array {
//     $$ = $1
// } | Struct {
//     $$ = $1
// } | Ident {
//     $$ = $1
// };

// Struct: '{' '}' {
//     $$ = Map{}
// } | '{' StructFields '}' {
//     $$ = Map{}
//     for _, item := range $2 {
//         key := item.(Array)[0]
//         switch in := key.(type) {
//             case StringLiteral:
//                 $$.Keys = append($$.Keys, string(in))
//             case Ident_:
//                 if len(in) > 0 {
//                     $$.Keys = append($$.Keys, "@"+strings.Join(in, "."))
//                 } else {
//                     $$.Keys = append($$.Keys, "")
//                 }
//             case nil:
//                 $$.Keys = append($$.Keys, "")
//             default:
//                 log.Printf("Annotation#Fields: not supported key type %T\n", key)
//         }
//         $$.Values.Add(item.(Array)[1])
//     }
// };

// StructFields: StructField {
//     $$ = Array{$1}
// } | StructFields ',' StructField {
//     $1.Add($3)
//     $$ = $1
// };

// StructField: Ident '=' Literal {
//     $$ = Array{$1, $3}
// } | Ident '=' Array {
//     $$ = Array{$1, $3}
// } | Ident '=' Ident {
//     $$ = Array{$1, $3}
// } | Ident '=' Struct {
//     $$ = Array{$1, $3}
// };
