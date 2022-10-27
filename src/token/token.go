package token

type Token struct {
	TokenType TokenType
	Lexeme    string
	Line      int
	Literal   interface{}
}

type TokenType byte
