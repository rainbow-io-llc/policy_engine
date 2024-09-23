grammar privacy;

policy     : rule+ ;
rule       : 'ALLOW' | 'DENY' action 'TO' target ('WHEN' condition)? ';' ;
action     : 'share' | 'collect' | 'use' ;
target     : 'third parties' | 'partners' | 'internal services' ;
condition  : 'time is within' TIME_RANGE | 'location is' LOCATION ;
TIME_RANGE : '"' [0-9]{2} ':' [0-9]{2} '-' [0-9]{2} ':' [0-9]{2} '"' ;
LOCATION   : '"' [a-zA-Z0-9, ]+ '"' ;

WS : [ \t\r\n]+ -> skip ;
