package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//Main y gestión ===

func main() {
	fmt.Println("=== Proyecto 2 - Coloreador Léxico ===")
	fmt.Println("Analizador léxico con colores ANSI (lenguaje Go)\n")

	if len(os.Args) < 2 {
		fmt.Println("Uso: go run . -- <ruta_archivo>")
		fmt.Println("Ejemplo: go run . -- test/valido1.go")
		return
	}

	// Obtener ruta del archivo
	path := ""
	for _, arg := range os.Args[1:] {
		if arg != "--" {
			path = arg
			break
		}
	}

	path = filepath.FromSlash(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("❌ El archivo no existe:", path)
		return
	}

	fmt.Printf("📂 Analizando archivo: %s\n\n", path)

	if err := AnalyzeFile(path); err != nil {
		fmt.Println("Error durante el análisis:", err)
		return
	}

	fmt.Println("\n--- Fin del análisis ---")
	fmt.Println("Repositorio: https://github.com/DakAfk/Coloreador_lexico")
}
