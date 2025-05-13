grammar BabyDuck;
// Lexer Rules

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
STRING : '"' ~('"')* '"';



WS : [ \t\r\n]+ -> skip ;

// Parser Rules

// Main Program
program: 'program' ID SEMICOLON (vars)? (funcs)* 'main' body 'end'  ;




// Variable Declaration
vars : 'var' var_decl+ ;
var_decl : id_list COLON type SEMICOLON ;
id_list : ID (',' ID)* ;
type : INTTYPE | FLOATTYPE ;

// Body
body : '{' statement* '}' ;

// Statement
statement : assign
          | cycle
          | f_call
          | print_stmt
          | condition ;

//  Assign
assign : ID '=' expression SEMICOLON ;

// While
cycle : 'while' '('expression ')' 'do' body SEMICOLON ;

// Conditional
condition : 'if' '(' expression ')' body else_part SEMICOLON ;
else_part : ('else' body)? ;

// Print
print_stmt : 'print' '(' printexpr (COMMA printexpr)* ')' SEMICOLON ;
printexpr : exp
          | STRING ;

// Constant
constant : INT
        | FLOAT ;


//Expressions
expression : exp (relational)? ;
relational : relop exp;
relop : '>' | '<'  | '!=' ;

// Arithmetic Expressions
exp : term (addop term)* ;
addop : '+' | '-' ;
term : factor (mulop factor)* ;
mulop : '*' | '/' ;

// Factor
factor : parexpr | factorsign;
parexpr : LPAREN expression RPAREN ;
factorsign : (addop)? value;
value: ID | constant;




// Functions
funcs : func;
func : 'void' ID LPAREN param_list? RPAREN funcbody ;
param_list : param (COMMA param)* ;
param : ID COLON type  ;
funcbody : LBRACKET vars? body RBRACKET SEMICOLON ;


f_call : ID LPAREN arg_list? RPAREN SEMICOLON ;
arg_list : expression (COMMA expression)* ;



