%{
package annotation

import (
    _ "fmt"
)

%}


%union {
    token *Token
    annot *Annotation
    annotField *AnnotationField
    annotFields []*AnnotationField

    val Value
}

%token <token>
ident '(' ')' '@' '=' ',' '{' '}' '[' ']'
nullLiteral stringLiteral boolLiteral integerLiteral floatLiteral

// special tokents for pre processing
annotation beginAnnotation endAnnotation

%type <val> Literal Ident
%type <annot> Content Annotation
%type <annotField> Field
%type <annotFields> Fields


%start doc
%%

doc : Content;

Content: /*empty*/ {
    $$ = annotationlex.(*annotationLex).annot
}
|
Content '@' Annotation {
    annotationlex.(*annotationLex).annot.Name = $3.Name
    annotationlex.(*annotationLex).annot.Fields = $3.Fields
    $$ = annotationlex.(*annotationLex).annot
}
;

Annotation:
    ident {
        $$ = &Annotation{Name: $1.Ident()}
    } |
    ident '(' ')' {
        $$ = &Annotation{Name: $1.Ident()}
    } |
    ident '(' Fields ')' {
        $$ = &Annotation{Name: $1.Ident(), Fields: $3}
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
    Ident {
        $$ = &AnnotationField{
            Value: $1,
        }
    } |
    ident '=' Ident {
        $$ = &AnnotationField{
            Value: $1,
        }
    } |
    ident '=' Literal {
        $$ = &AnnotationField{
            Key: $1.Ident(),
            Value: $2,
        }
    } |
    ident '=' Array {
        // $$ = &AnnotationField{
        //     Key: PromiseIdent($1),
        //     Value: $1,
        // }
    } |
    ident '=' Struct {
        // $$ = &AnnotationField{
        //     Key: PromiseIdent($1),
        //     Value: $1,
        // }
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

Array:
    '[' ArrayFields ']' |
    '[' ']';

ArrayFields:
    ArrayField |
    ArrayFields ',' ArrayField;

ArrayField:
    Literal {} |
    Array {} |
    Struct {};

Struct:
    '{' '}' |
    '{' StructFields '}';

StructFields:
    StructField |
    StructFields ',' StructField;

StructField:
    ident '=' Literal |
    ident '=' Array |
    ident '=' Ident |
    ident '=' Struct
;
