// Programa factorial simple
package main

func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	var num = 5
	println("El factorial es:", factorial(num))
}
