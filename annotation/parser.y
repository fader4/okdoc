%{
package annotation

import (
    _ "fmt"
    "strings"
    "log"
    "github.com/fader4/okdoc/syntax"
)

%}


%union {
    token *syntax.TokenWithData
    annot *Annotation
    annotField *AnnotationField
    annotFields []*AnnotationField

    val syntax.Value
    arr syntax.Array
    map_ syntax.Map
}

%token <token>
ident '(' ')' '@' '=' ',' '{' '}' '[' ']'
nullLiteral stringLiteral boolLiteral integerLiteral floatLiteral

// special tokents for pre processing
annotation beginAnnotation endAnnotation

%type <val> Literal Ident ArrayField
%type <annot> Content Annotation
%type <annotField> Field
%type <annotFields> Fields
%type <arr> StructField StructFields ArrayFields Array
%type <map_> Struct


%start doc
%%

doc : Content;

Content: /*empty*/ {
    $$ = annotationlex.(*annotationLex).annot
} | Content '@' Annotation {
    annotationlex.(*annotationLex).annot.name = $3.name
    annotationlex.(*annotationLex).annot.fields = $3.fields
    $$ = annotationlex.(*annotationLex).annot
};

Annotation: ident {
    $$ = &Annotation{name: $1.Ident()}
} | ident '(' ')' {
    $$ = &Annotation{name: $1.Ident()}
} | ident '(' Fields ')' {
    $$ = &Annotation{name: $1.Ident(), fields: $3}
};

Fields: Fields ',' Field {
    $$ = append($1, $3)
} | Field {
    $$ = []*AnnotationField{$1}
};

Field:
    Literal {
        $$ = &AnnotationField{
            Value: $1,
        }
    } |
    Struct {
        $$ = &AnnotationField{
            Value: $1,
        }
    } |
    Array {
        $$ = &AnnotationField{
            Value: $1,
        }
    } |
    Ident {
        $$ = &AnnotationField{
            Value: $1,
        }
    } |
    Ident '=' Ident {
        $$ = &AnnotationField{
            Value: $1,
        }
    } |
    Ident '=' Literal {
        $$ = &AnnotationField{
            Key: $1,
            Value: $3,
        }
    } |
    Ident '=' Array {
        $$ = &AnnotationField{
            Key: $1,
            Value: $3,
        }
    } |
    Ident '=' Struct {
        $$ = &AnnotationField{
            Key: $1,
            Value: $3,
        }
    }
;

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

Array: '[' ArrayFields ']' {
    $$ = $2
} | '[' ']' {
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
                log.Printf("Annotation#Fields: not supported key type %T\n", key)
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

StructField: Ident '=' Literal {
    $$ = syntax.Array{$1, $3}
} | Ident '=' Array {
    $$ = syntax.Array{$1, $3}
} | Ident '=' Ident {
    $$ = syntax.Array{$1, $3}
} | Ident '=' Struct {
    $$ = syntax.Array{$1, $3}
};
