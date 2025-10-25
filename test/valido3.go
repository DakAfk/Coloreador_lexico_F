package main

const limite = 10

func main() {
	var contador = 0
	for contador < limite {
		contador = contador + 1
		if contador == limite {
			println("Se alcanzó el límite:", true)
		}
	}
}
