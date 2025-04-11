package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scan(&op)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&b)

	switch op {
	case "+":
		fmt.Printf("Hasil: %.2f\n", a+b)
	case "-":
		fmt.Printf("Hasil: %.2f\n", a-b)
	case "*":
		fmt.Printf("Hasil: %.2f\n", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("Hasil: %.2f\n", a/b)
		} else {
			fmt.Println("Error: Pembagian dengan nol bro.")
		}
	default:
		fmt.Println("Operator gak valid.")
	}
}
