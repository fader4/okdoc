%{
package starlark

import (
    "fmt"
)

%}


%union {
    token *tokenWithLexer
    node string

    argument ArgumentNode
    arguments []ArgumentNode
    reportFile *ReportFile

    loadNode *ReportFile_Import
    exportNode *ReportFile_Export
    fnNode *ReportFile_Function
    commentNode *ReportFile_Comment
    returnNode *ReportFile_Return
}

%type <commentNode> Comment
%type <loadNode> LoadStmt
%type <reportFile> Content
%type <argument>
LoadStmtArgument
%type <arguments>
LoadStmtArguments
%type <exportNode> ModuleStmt
%type <fnNode> DefStmt
%type <returnNode> ReturnStmt

%token <token>
commentInline commentMultiline
ident boolLiteral integerLiteral floatLiteral stringLiteral
defKeyword loadKeyword returnKeyword moduleKeyword
'.' ',' ':' '(' ')' '[' ']' '{' '}' '=' '*'

%start doc
%%

doc : Content {
    starlarklex.(reporter).SetReport($1)
};

Content: /*empty*/ {
    $$ = &ReportFile{}
}
|
Content Comment {
    fmt.Println("Comment:", $2)
    $$.Comments = append($$.Comments, $2)
}
|
Content DefStmt {
    fmt.Println("DefStmt:", $2)
    $$.Functions = append($$.Functions, $2)
}
|
Content ReturnStmt {
    // fmt.Println("ReturnStmt")
    $$.Returns = append($$.Returns, $2)
}
|
Content LoadStmt {
    fmt.Println("LoadStmt:", $2)
    $$.Imports = append($$.Imports, $2)
}
|
Content ModuleStmt {
    fmt.Println("ModuleStmt:", $2)
    $$.Exports = append($$.Exports, $2)
}
|
Content FreeTokens;


Comment: commentInline {
    $$ = &ReportFile_Comment{
        Value: $1.String(),
        Pos: $1.Pos(),
    }
} | commentMultiline {
    $$ = &ReportFile_Comment{
        Value: $1.String(),
        MultiLine: true,
        Pos: $1.Pos(),
    }
};


LoadStmt:
    loadKeyword '(' stringLiteral  LoadStmtArguments ')' {
        $$ = &ReportFile_Import{Name: $3.String(), Args: $4, Pos: $1.Pos()}
    };

LoadStmtArgument:
    ident '=' stringLiteral {
        $$ = ArgumentNode{
            ArgumentNode_Value{Ident: $1.String()},
            ArgumentNode_Value{StringLiteral: $3.String()},
        }
    } |
    stringLiteral {
        $$ = ArgumentNode{
            ArgumentNode_Value{StringLiteral: $1.String()},
        }
    };
LoadStmtArguments: LoadStmtArgument {
    $$ = []ArgumentNode{$1}
} | LoadStmtArguments LoadStmtArgument {
    $$ = append($1, $2)
};

ModuleStmt:
    moduleKeyword '(' stringLiteral {
        $$ = &ReportFile_Export{Name: $3.String(), Pos: $1.Pos()}
    };

DefStmt:
    defKeyword ident {
        $$ = &ReportFile_Function{Name: $2.String(), Pos: $1.Pos()}
    };

ReturnStmt:
    returnKeyword {
        $$ = &ReportFile_Return{Pos: $1.Pos()}
    };

FreeTokens:
    ident | '(' | ')' | '=' | stringLiteral;
