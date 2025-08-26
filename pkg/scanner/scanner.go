package scanner

import "github.com/kevhlee/glox/pkg/token"

// Contains the internal state and logic of the Lox scanner.
type scanner struct {
	source  []rune
	tokens  []*token.Token
	start   int
	current int
	line    int
}

func (s *scanner) isScanning() bool {
	return s.current < len(s.source)
}

func (s *scanner) advance() rune {
	s.current++
	return s.source[s.current-1]
}

func (s *scanner) peek() rune {
	if s.isScanning() {
		return s.source[s.current]
	}
	return 0
}

func (s *scanner) peekNext() rune {
	if s.current < len(s.source)-1 {
		return s.source[s.current+1]
	}
	return 0
}

func (s *scanner) match(expected rune) bool {
	if s.isScanning() && s.peek() == expected {
		s.advance()
		return true
	}
	return false
}

func (s *scanner) error(msg string) {
	s.tokens = append(s.tokens, &token.Token{
		Type:   token.ERROR,
		Lexeme: msg,
		Line:   s.line,
	})
}

func (s *scanner) addToken(t token.Type) {
	s.tokens = append(s.tokens, &token.Token{
		Type:   t,
		Lexeme: string(s.source[s.start:s.current]),
		Line:   s.line,
	})
}

func (s *scanner) scan() []*token.Token {
	if len(s.tokens) > 0 {
		// We can assume that the scanning was already done
		return s.tokens
	}

	for s.isScanning() {
		s.start = s.current

		switch ch := s.advance(); ch {
		case '(':
			s.addToken(token.LEFT_PAREN)
		case ')':
			s.addToken(token.RIGHT_PAREN)
		case '{':
			s.addToken(token.LEFT_BRACE)
		case '}':
			s.addToken(token.RIGHT_BRACE)
		case ',':
			s.addToken(token.COMMA)
		case '.':
			s.addToken(token.DOT)
		case '+':
			s.addToken(token.PLUS)
		case '-':
			s.addToken(token.MINUS)
		case '*':
			s.addToken(token.STAR)
		case ';':
			s.addToken(token.SEMICOLON)
		case '/':
			if s.match('/') {
				for s.isScanning() && s.peek() != '\n' {
					s.advance()
				}
			} else {
				s.addToken(token.SLASH)
			}
		case '!':
			if s.match('=') {
				s.addToken(token.BANG_EQUAL)
			} else {
				s.addToken(token.BANG)
			}
		case '=':
			if s.match('=') {
				s.addToken(token.EQUAL_EQUAL)
			} else {
				s.addToken(token.EQUAL)
			}
		case '>':
			if s.match('=') {
				s.addToken(token.GREATER_EQUAL)
			} else {
				s.addToken(token.GREATER)
			}
		case '<':
			if s.match('=') {
				s.addToken(token.LESS_EQUAL)
			} else {
				s.addToken(token.LESS)
			}
		case ' ', '\r', '\t':
			continue
		case '\n':
			s.line++
		case '"':
			s.scanString()
		default:
			if token.IsDigit(ch) {
				s.scanNumber()
			} else if token.IsIdentStart(ch) {
				s.scanIdentifier()
			} else {
				s.error("Unexpected character")
			}
		}
	}

	s.tokens = append(s.tokens, &token.Token{
		Type: token.EOF,
		Line: s.line,
	})

	return s.tokens
}

func (s *scanner) scanString() {
	for s.isScanning() && s.peek() != '"' {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if !s.isScanning() {
		s.error("Unterminated string")
		return
	}

	s.advance()
	s.addToken(token.STRING)
}

func (s *scanner) scanNumber() {
	for token.IsDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && token.IsDigit(s.peekNext()) {
		s.advance()

		for token.IsDigit(s.peek()) {
			s.advance()
		}
	}

	s.addToken(token.NUMBER)
}

func (s *scanner) scanIdentifier() {
	for token.IsIdentPart(s.peek()) {
		s.advance()
	}

	t, ok := token.ParseKeyword(string(s.source[s.start:s.current]))
	if !ok {
		t = token.IDENTIFIER
	}
	s.addToken(t)
}
