grammar eval;

program
    : statement EOF;

variable
    : BOOL                               # Bool
    | STRING                             # String
    | FLOAT                              # Float
    | INTEGER                            # Integer
    | VAR                                # Var
    ;

statement
   : variable # V
   | function # F
   | op=(NOT|SUB) statement # Unary
   | statement op=(MUL|DIV|MOD) statement # MulDiv
   | statement op=(ADD|SUB) statement # AddSUb
   | statement op=(LE|LT|GE|GT) statement # Comparison1
   | statement op=(NE|EQ) statement # Comparison2
   | statement op=AND statement # And
   | statement op=OR statement # Or
//   | statement op=(MUL|DIV|ADD|SUB|LE|LT|GE|GT|NE|EQ|AND|OR) statement # Binary
   | '(' statement ')' # Parens
   ;

function
   : f=VAR '(' statement (',' statement)* ')'
   | f=VAR '()'
   ;


MUL: '*';
DIV: '/';
MOD: '%';
ADD: '+';
SUB: '-';
AND: '&&';
NOT: '!';
OR: '||';
LE: '<=';
LT: '<';
GE: '>=';
GT: '>';
NE: '~='|'!=';
EQ: '==';
BOOL
    : TRUE
    | FALSE
    ;
TRUE: 'true';
FALSE: 'false';

STRING: '"'.*?'"'|'\''.*?'\'';
INTEGER: DIGIT+;
FLOAT: DIGIT+ '.'DIGIT+;
VAR: [a-zA-Z_][a-zA-Z0-9_.]*;

WHITESPACE: [ \r\n\t]+ -> skip;

fragment DIGIT  : [0-9];