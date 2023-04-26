grammar lexDotGraph;

graph: GRAPH_TYPE NAME OPEN_BRACE
(vertex_declaration | edge_declaration)*
CLOSE_BRACE EOF;

vertex_declaration: vname=NAME attribute_declaration? SEMICOLON;

edge_declaration: (unvar=undirect_edge_declaration | dirvar=direct_edge_declaration) attribute_declaration?  SEMICOLON;

undirect_edge_declaration: vname1=NAME UNDIRECT_EDGE_OP vname2=NAME;

direct_edge_declaration: vname1=NAME DIRECT_EDGE_OP vname2=NAME;

attribute_declaration: OPEN_SQUARE_BRACKET NAME EQUALS_OP NAME CLOSE_SQUARE_BRACKET;

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



