package scanner

import "hart/src/token"

var (
	file   []byte
	char   = 0
	start  int
	line   int
	tokens []token.Token
)

func Scanner(fileIn []byte) {
	file = fileIn
	scanTokens()
}

func addToken(tokenType token.TokenType, literal interface{}) {
	text := string(file[start:char])
	tokens = append(tokens, token.Token{TokenType: tokenType, Lexeme: text, Line: line, Literal: literal})
}

func addCleanToken(tokenType token.TokenType) {
	addToken(tokenType, nil)
}

func scanTokens() {
	if !isEnd() {
		start = char
		scanToken()
	}
}

func scanToken() {
	//nowChar := advance()
}

func advance() byte {
	char++
	return file[char-1]
}

func peek() byte {
	if isEnd() {
		return 0
	}
	return file[char]
}

func peekNext() byte {
	if char+1 >= len(file) {
		return 0
	}
	return file[char+1]
}

func match(expected byte) bool {
	if isEnd() {
		return false
	}
	if file[char] != expected {
		return false
	}
	char++
	return true
}

func convertstring() {
	for peek() != '"' && !isEnd() {
		if peek() == '\n' {
			line++
		}
		advance()
	}
	if isEnd() {
		return
	}
	advance()
	value := string(file[start+1 : char-1])
	addToken(token.STRING, value)
}

func isEnd() bool {
	return char >= len(file)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func number() {
	for isDigit(peek()) {
		advance()
	}
	if peek() == '.' && isDigit(peekNext()) {
		advance()
		for isDigit(peek()) {
			advance()
		}
	}
	addToken(token.NUMBER, string(file[start:char]))
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func variable() {
	for isAlphaNumeric(peek()) {
		advance()
	}
	text := string(file[start:char])
	tokenType := token.VAR
	addToken(tokenType, text)
}
