grammar authz;

policy     : rule+ ;
rule: ('ALLOW' | 'DENY') subject 'FROM' action 'TO' object ('WHEN' condition)? ';' ;
subject    : 'user' | 'service' ;
action     : 'share' | 'collect' | 'use' ;
object     : 'data' | 'service' | 'feature' ;
condition  : 'time_range' '=' TIME_RANGE | 'location' '=' LOCATION ;
TIME_RANGE : '"' [0-9] [0-9] ':' [0-9] [0-9] '-' [0-9] [0-9] ':' [0-9] [0-9] '"';
LOCATION   : '"' [a-zA-Z0-9, ]+ '"' ;

WS : [ \t\r\n]+ -> skip ;
