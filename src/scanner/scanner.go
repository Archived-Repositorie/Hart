package scanner

import (
	"fmt"
	"hart/src/token"
	"log"
)

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
	fmt.Println(len(tokens))
}

func scanToken() {
	nowChar := advance()
	switch nowChar {
	case '(':
		addCleanToken(token.OPEN_PAREN)
	case ')':
		addCleanToken(token.CLOSE_PAREN)
	case '?':
		if match(':') {
			addCleanToken(token.ELSE_IF)
		} else {
			addCleanToken(token.IF)
		}
	case ':':
		addCleanToken(token.ELSE)
	case '~':
		if match('*') {
			addCleanToken(token.RETURN)
		} else {
			addCleanToken(token.APPROX)
		}
	case '!':
		if match('!') {
			addCleanToken(token.FACTORY)
		} else {
			addCleanToken(token.NEGATION)
		}
	case '$':
		if match('@') {
			addCleanToken(token.CONST)
		} else {
			addCleanToken(token.VAR)
		}
	case '@':
		for peek() != '\n' && !isEnd() {
			advance()
		}
	case '#':
		if match('#') {
			addCleanToken(token.EXPORT)
		} else {
			addCleanToken(token.IMPORT)
		}
	case '%':
		if match('%') {
			addCleanToken(token.PERCENT)
		} else {
			addCleanToken(token.MODULE)
		}
	case '^':
		if match('^') {
			addCleanToken(token.LOG)
		} else {
			addCleanToken(token.BIT_XOR)
		}
	case '&':
		if match('&') {
			addCleanToken(token.BOOL_AND)
		} else {
			addCleanToken(token.BIT_AND)
		}
	case '*':
		if match('*') {
			addCleanToken(token.POWER)
		} else {
			addCleanToken(token.MULTIPLY)
		}
	case '-':
		addCleanToken(token.MINUS)
	case '+':
		addCleanToken(token.PLUS)
	case '=':
		if match('=') {
			addCleanToken(token.EQUALS)
		} else if match('>') {
			addCleanToken(token.FUNCTION)
		} else {
			addCleanToken(token.CREATE_VAR)
		}
	case '<':
		if match('=') {
			addCleanToken(token.EQUALS_LESS)
		} else if match('<') {
			addCleanToken(token.PRINT)
		} else {
			addCleanToken(token.LESS_THAN)
		}
	case '>':
		if match('=') {
			addCleanToken(token.EQUALS_MORE)
		} else if match('>') {
			addCleanToken(token.INPUT)
		} else {
			addCleanToken(token.MORE_THAN)
		}
	case '|':
		if match('|') {
			addCleanToken(token.BOOL_OR)
		} else {
			addCleanToken(token.BIT_OR)
		}
	case '/':
		if match('/') {
			addCleanToken(token.SQUARE_ROOT)
		} else {
			addCleanToken(token.DIVIDE)
		}
	case '"':
		convertString()
	default:
		if nowChar == '[' && match(']') {
			addCleanToken(token.LIST)
		} else if isDigit(nowChar) {
			number()
		} else if isAlpha(peekNext()) {
			if match('@') {
				variable(true)
			} else {
				variable(false)
			}
		} else if nowChar == ' ' || nowChar == '\r' || nowChar == '\t' {
			// Ignore whitespace.
		} else if nowChar == '\n' {
			line++
		} else {
			log.Fatalf("Unexpected character %c", nowChar)
		}
	}
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

func convertString() {
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
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' || c == '.'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func variable(constant bool) {
	for isAlphaNumeric(peek()) {
		advance()
	}
	text := string(file[start:char])
	tokenType := token.VAR
	if constant {
		tokenType = token.CONST
	}
	addToken(tokenType, text)
}
