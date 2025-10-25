package main

import "fmt"

// === Trabajo Persona 2: Coloreador ===

// Códigos de color ANSI para consola (modo oscuro)
const (
	Reset  = "\033[0m"
	Blue   = "\033[34m"         // Palabras reservadas
	Orange = "\033[38;5;208m"   // Números / Constantes
	White  = "\033[97m"         // Agrupadores
	Yellow = "\033[33m"         // Operadores / Lógicos
	Green  = "\033[92m"         // Cadenas
	Pink   = "\033[95m"         // Identificadores
	Gray   = "\033[90m"         // Comentarios
	RedBg  = "\033[41m\033[97m" // Errores (fondo rojo, texto blanco)
)

// ColorizeToken devuelve el lexema coloreado según el tipo de token
func ColorizeToken(t Token) string {
	switch t.Type {
	case "KEYWORD":
		return fmt.Sprintf("%s%s%s", Blue, t.Lexeme, Reset)
	case "NUMBER":
		return fmt.Sprintf("%s%s%s", Orange, t.Lexeme, Reset)
	case "STRING":
		return fmt.Sprintf("%s\"%s\"%s", Green, t.Lexeme, Reset)
	case "GROUP":
		return fmt.Sprintf("%s%s%s", White, t.Lexeme, Reset)
	case "OPERATOR":
		return fmt.Sprintf("%s%s%s", Yellow, t.Lexeme, Reset)
	case "IDENTIFIER":
		return fmt.Sprintf("%s%s%s", Pink, t.Lexeme, Reset)
	case "COMMENT":
		return fmt.Sprintf("%s%s%s", Gray, t.Lexeme, Reset)
	case "ERROR_UNCLOSED_STRING", "ERROR_UNKNOWN":
		return fmt.Sprintf("%s%s%s", RedBg, t.Lexeme, Reset)
	default:
		return t.Lexeme
	}
}
