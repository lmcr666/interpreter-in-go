package token

type TokenType string

type Token struct { 
  Type    TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  // Identifiers + literals
  IDENT = "IDENT" // add, foobar, x, y, ...
  INT   = "INT"   // 1343456

  // Operators
  ASSIGN   = "be"
  PLUS     = "+"

  // Delimitersi
  COMMA    = ","
  SEMICOLON = ";"
  LPAREN   = "(" 
  RPAREN   = ")"
  LBRACE   = "{"
  RBRACE   = "}"

  // Keywords
  FUNCTION = "fun"
  LET      = "make"
)
