grammar BabyDuck;

// -------- Lexer Rules --------
VOID : 'void' ;
INTTYPE : 'int' ;
FLOATTYPE : 'float' ;

LPAREN : '(' ;
RPAREN : ')' ;
LBRACKET : '[' ;
RBRACKET : ']' ;
COLON : ':' ;
COMMA : ',' ;
SEMICOLON : ';' ;

ID : [a-zA-Z_][a-zA-Z_0-9]* ;
INT : [0-9]+ ;
FLOAT : [0-9]+ '.' [0-9]+ ;
STRING : '"' ~('"')* '"' ;

WS : [ \t\r\n]+ -> skip ;

// -------- Parser Rules --------

// Programa principal
program
    : 'program' ID SEMICOLON (vars)? (funcs)* 'main' body 'end'
    ;

// Declaración de variables
vars
    : 'var' var_decl+
    ;

var_decl
    : id_list COLON type SEMICOLON
    ;

id_list
    : ID (COMMA ID)*
    ;

type
    : INTTYPE
    | FLOATTYPE
    ;

// Bloque de código
body
    : '{' statement* '}'
    ;

// Sentencias
statement
    : assign
    | cycle
    | f_call
    | print_stmt
    | condition
    ;

// Asignación
assign
    : ID '=' expression SEMICOLON
    ;

// While
cycle
    : 'while' LPAREN expression RPAREN 'do' body SEMICOLON
    ;

// Condicional
condition
    : 'if' LPAREN expression RPAREN body else_part SEMICOLON
    ;

else_part
    : ('else' body)?
    ;

// Print
print_stmt
    : 'print' LPAREN printexpr (COMMA printexpr)* RPAREN SEMICOLON
    ;

printexpr
    : expression
    | STRING
    ;

// Constantes
constant
    : INT
    | FLOAT
    ;

// -------- Expresiones con precedencia --------

// Entrada principal
expression
    : rel_expr
    ;

// Relacionales
rel_expr
    : add_expr (relop add_expr)?
    ;

relop
    : '>'
    | '<'
    | '!='
    ;

// Suma / Resta
add_expr
    : term (addop term)*
    ;

addop
    : '+'
    | '-'
    ;

// Multiplicación / División
term
    : factor (mulop factor)*
    ;

mulop
    : '*'
    | '/'
    ;

// Factor (con soporte para paréntesis y signos)
factor
    : LPAREN expression RPAREN
    | (addop)? value
    ;

value
    : ID
    | constant
    ;

// -------- Funciones --------
funcs
    : func
    ;

func
    : VOID ID LPAREN param_list? RPAREN funcbody
    ;

param_list
    : param (COMMA param)*
    ;

param
    : ID COLON type
    ;

funcbody
    : LBRACKET vars? body RBRACKET SEMICOLON
    ;

f_call
    : ID LPAREN arg_list? RPAREN SEMICOLON
    ;

arg_list
    : expression (COMMA expression)*
    ;
