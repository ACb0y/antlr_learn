grammar LabeledExpr;
import CommonLexerRules;

prog : expr_list += stat+ ;

stat : expr NEWLINE         # printExpr
     | ID '=' expr NEWLINE  # assign
     | NEWLINE              # blank
     ;

expr : expr op=('*'|'/') expr   # MulDiv
     | expr op=('+'|'-') expr   # AddSub
     | INT                      # int
     | ID                       # id
     | '(' expr ')'             # parens
     ;

MUL : '*';
DIV : '/';
ADD : '+';
SUB : '-';
