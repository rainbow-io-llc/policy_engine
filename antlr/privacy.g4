grammar privacy;

policy     : rule+ ;
rule       : ('ALLOW' | 'DENY') action 'TO' target ('WHEN' condition)? ';' ;
action     : 'share' | 'collect' | 'use' ;
target     : 'third parties' | 'partners' | 'internal services' ;
condition  : condition_ (('AND' | 'OR') condition_)* ;
condition_ : 'time_range' '=' TIME_RANGE | 'location' '=' LOCATION ;
TIME_RANGE : '"' [0-9] [0-9] ':' [0-9] [0-9] '-' [0-9] [0-9] ':' [0-9] [0-9] '"';
LOCATION   : '"' [a-zA-Z0-9, ]+ '"' ;

WS : [ \t\r\n]+ -> skip ;
