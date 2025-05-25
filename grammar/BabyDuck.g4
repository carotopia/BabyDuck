grammar BabyDuck;

// -------- Lexer Rules --------
// Palabras clave
PROGRAM : 'program' ;
MAIN : 'main' ;
END : 'end' ;
VAR : 'var' ;
VOID : 'void' ;
INTTYPE : 'int' ;
FLOATTYPE : 'float' ;
WHILE : 'while' ;
DO : 'do' ;
IF : 'if' ;
ELSE : 'else' ;
PRINT : 'print' ;

// Operadores de comparación
GT : '>' ;
LT : '<' ;
NE : '!=' ;

// Operadores aritméticos
PLUS : '+' ;
MINUS : '-' ;
MULT : '*' ;
DIV : '/' ;

// Operador de asignación
ASSIGN : '=' ;

// Símbolos de puntuación
LPAREN : '(' ;
RPAREN : ')' ;
LBRACKET : '[' ;
RBRACKET : ']' ;
LBRACE : '{' ;
RBRACE : '}' ;
COLON : ':' ;
COMMA : ',' ;
SEMICOLON : ';' ;

// Identificadores y literales
ID : [a-zA-Z_][a-zA-Z_0-9]* ;
INT : [0-9]+ ;
FLOAT : [0-9]+ '.' [0-9]+ ;
STRING : '"' ~('"')* '"' ;

// Espacios en blanco
WS : [ \t\r\n]+ -> skip ;

// -------- Parser Rules --------

// Programa principal
program
    : PROGRAM ID SEMICOLON (vars)? (funcs)* MAIN body END
    ;

// Declaración de variables
vars
    : VAR var_decl+
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
    : LBRACE statement* RBRACE
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
    : ID ASSIGN expression SEMICOLON
    ;

// While
cycle
    : WHILE LPAREN expression RPAREN DO body SEMICOLON
    ;

// Condicional
condition
    : IF LPAREN expression RPAREN body else_part SEMICOLON
    ;

else_part
    : (ELSE body)?
    ;

// Print
print_stmt
    : PRINT LPAREN printexpr (COMMA printexpr)* RPAREN SEMICOLON
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
    : GT
    | LT
    | NE
    ;

// Suma / Resta
add_expr
    : term (addop term)*
    ;

addop
    : PLUS
    | MINUS
    ;

// Multiplicación / División
term
    : factor (mulop factor)*
    ;

mulop
    : MULT
    | DIV
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