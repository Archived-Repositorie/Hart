package token

const (
	// Character tokens
	OPEN_PAREN  TokenType = 0x0
	CLOSE_PAREN TokenType = 0x1
	IF          TokenType = 0x2
	ELSE        TokenType = 0x3
	ELSE_IF     TokenType = 0x4
	APPROX      TokenType = 0x5
	RETURN      TokenType = 0x6
	NEGATION    TokenType = 0x7
	FACTORY     TokenType = 0x8
	CONST       TokenType = 0x9
	IMPORT      TokenType = 0xA
	EXPORT      TokenType = 0xB
	MODULE      TokenType = 0xC
	PERCENT     TokenType = 0xD
	BIT_XOR     TokenType = 0xE
	LOG         TokenType = 0xF
	BIT_AND     TokenType = 0x10
	BOOL_AND    TokenType = 0x11
	MULTIPLY    TokenType = 0x12
	POWER       TokenType = 0x13
	MINUS       TokenType = 0x14
	PLUS        TokenType = 0x15
	CREATE_VAR  TokenType = 0x16
	EQUALS      TokenType = 0x17
	FUNCTION    TokenType = 0x18
	LESS_THAN   TokenType = 0x19
	PRINT       TokenType = 0x1A
	EQUALS_LESS TokenType = 0x1B
	MORE_THAN   TokenType = 0x1C
	INPUT       TokenType = 0x1D
	EQUALS_MORE TokenType = 0x1E
	LIST        TokenType = 0x1F
	BIT_OR      TokenType = 0x20
	BOOL_OR     TokenType = 0x21
	DIVIDE      TokenType = 0x22
	SQUARE_ROOT TokenType = 0x23
	//Literals
	STRING TokenType = 0x27
	NUMBER TokenType = 0x28
	VAR    TokenType = 0x29
)
