%{
package annotation

import (
    _ "fmt"
    "strings"
    "log"
)

%}


%union {
    token *Token
    annot *Annotation
    annotField *AnnotationField
    annotFields []*AnnotationField

    val Value
    arr Array
    map_ Map
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
    $$ = $1.(Ident_).Append($3.Ident())
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
    $$ = Array{}
};

ArrayFields: ArrayField {
    $$ = Array{$1}
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
    $$ = Map{}
} | '{' StructFields '}' {
    $$ = Map{}
    for _, item := range $2 {
        key := item.(Array)[0]
        switch in := key.(type) {
            case StringLiteral:
                $$.Keys = append($$.Keys, string(in))
            case Ident_:
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
        $$.Values.Add(item.(Array)[1])
    }
};

StructFields: StructField {
    $$ = Array{$1}
} | StructFields ',' StructField {
    $1.Add($3)
    $$ = $1
};

StructField: Ident '=' Literal {
    $$ = Array{$1, $3}
} | Ident '=' Array {
    $$ = Array{$1, $3}
} | Ident '=' Ident {
    $$ = Array{$1, $3}
} | Ident '=' Struct {
    $$ = Array{$1, $3}
};
