package main

import "unicode"

// === Trabajo Persona 1: Lexer ===

// Mapa de palabras reservadas (puedes ampliar si deseas)
var keywords = map[string]bool{
	"if": true, "else": true, "for": true, "while": true,
	"return": true, "func": true, "var": true, "const": true,
	"package": true, "import": true, "true": true, "false": true, "nil": true,
}

// Token representa un elemento reconocido
type Token struct {
	Type   string
	Lexeme string
	Line   int
	Col    int
}

// Funciones auxiliares
func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// tokenizeLine analiza una línea y devuelve tokens o el primer error encontrado
func tokenizeLine(line string, lineNo int) ([]Token, *Token) {
	var tokens []Token
	runes := []rune(line)
	n := len(runes)
	i := 0

	for i < n {
		ch := runes[i]

		// Ignorar espacios y tabulaciones
		if ch == ' ' || ch == '\t' {
			i++
			continue
		}

		col := i + 1

		// Identificadores o palabras reservadas
		if isLetter(ch) {
			start := i
			for i < n && (isLetter(runes[i]) || isDigit(runes[i])) {
				i++
			}
			lex := string(runes[start:i])
			if keywords[lex] {
				tokens = append(tokens, Token{"KEYWORD", lex, lineNo, col})
			} else {
				tokens = append(tokens, Token{"IDENTIFIER", lex, lineNo, col})
			}
			continue
		}

		// Números
		if isDigit(ch) {
			start := i
			for i < n && isDigit(runes[i]) {
				i++
			}
			tokens = append(tokens, Token{"NUMBER", string(runes[start:i]), lineNo, col})
			continue
		}

		// Cadenas entre comillas
		if ch == '"' {
			i++
			start := i
			for i < n && runes[i] != '"' {
				i++
			}
			if i < n {
				tokens = append(tokens, Token{"STRING", string(runes[start:i]), lineNo, col})
				i++ // saltar comilla final
			} else {
				errTok := Token{"ERROR_UNCLOSED_STRING", string(runes[start-1:]), lineNo, col}
				return tokens, &errTok
			}
			continue
		}

		// Comentarios "//"
		if ch == '/' && i+1 < n && runes[i+1] == '/' {
			tokens = append(tokens, Token{"COMMENT", string(runes[i:]), lineNo, col})
			break
		}

		// Agrupadores y operadores
		switch ch {
		case '{', '}', '(', ')', '[', ']', ',', ';', ':':
			tokens = append(tokens, Token{"GROUP", string(ch), lineNo, col})
		case '+', '-', '*', '/', '%', '=', '<', '>', '!':
			tokens = append(tokens, Token{"OPERATOR", string(ch), lineNo, col})
		default:
			errTok := Token{"ERROR_UNKNOWN", string(ch), lineNo, col}
			return tokens, &errTok
		}

		i++
	}

	return tokens, nil
}

// === Fin Trabajo Persona 1 ===
