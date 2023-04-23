grammar lexDotGraph;

GRAPH_TYPE:
  'digraph'|'graph';
NAME:
  [a-zA-Z0-9]+;
OPEN_BRACE
  : '{';
CLOSE_BRACE
  : '}';
OPEN_SQUARE_BRACKET
  : '[';
CLOSE_SQUARE_BRACKET
  :  ']';
DIRECT_EDGE_OP
  :  '->';
UNDIRECT_EDGE_OP
  :  '==';
EQUALS_OP
  :  '=';
SEMICOLON
  :  ';';
WS:
   [ \t\n\r]+ -> skip;



