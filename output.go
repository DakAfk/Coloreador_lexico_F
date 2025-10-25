package main

import (
	"bufio"
	"fmt"
	"os"
)

// === Trabajo Persona 3: Integración y salida ===

// AnalyzeFile analiza un archivo línea por línea usando el lexer y colorizador
func AnalyzeFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		line := scanner.Text()

		tokens, errTok := tokenizeLine(line, lineNo)

		// Muestra tokens coloreados en consola
		for _, t := range tokens {
			fmt.Print(ColorizeToken(t) + " ")
		}
		fmt.Println()

		// Si hay error, mostrarlo y detener análisis
		if errTok != nil {
			fmt.Printf("\n❌ Error en línea %d, columna %d: %s (%s)\n",
				errTok.Line, errTok.Col, errTok.Lexeme, errTok.Type)
			return nil
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error leyendo archivo: %v", err)
	}

	fmt.Println("\n✅ Archivo válido. No se detectaron errores léxicos.")
	return nil
}

// === Fin Trabajo Persona 3 ===
