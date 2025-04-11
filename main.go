package main

import "fmt"

func main() {

	// Basic Calculator 

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

	// Temperature Calculator

	var temp float64
	var unit string
	
	fmt.Print("Masukkan suhu (contoh: 30): ")
	fmt.Scan(&temp)
	fmt.Print("Masukkan satuan asal (C, F, K): ")
	fmt.Scan(&unit)

	switch unit {
	case "C", "c":
		fmt.Printf("Fahrenheit: %.2f°F\n", (temp*9/5)+32)
		fmt.Printf("Kelvin: %.2fK\n", temp+273.15)
	case "F", "f":
		fmt.Printf("Celsius: %.2f°C\n", (temp-32)*5/9)
		fmt.Printf("Kelvin: %.2fK\n", ((temp-32)*5/9)+273.15)
	case "K", "k":
		fmt.Printf("Celsius: %.2f°C\n", temp-273.15)
		fmt.Printf("Fahrenheit: %.2f°F\n", ((temp-273.15)*9/5)+32)
	default:
		fmt.Println("Satuan gak dikenal.")
	}

	// Currency Calculator

	const (
		usdToIdr = 16000.0
		eurToIdr = 17500.0
	)

	var amount float64
	var from, to string

	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&amount)
	fmt.Print("Dari mata uang (IDR, USD, EUR): ")
	fmt.Scan(&from)
	fmt.Print("Ke mata uang (IDR, USD, EUR): ")
	fmt.Scan(&to)

	var result float64

	switch from {
	case "IDR":
		if to == "USD" {
			result = amount / usdToIdr
		} else if to == "EUR" {
			result = amount / eurToIdr
		} else if to == "IDR" {
			result = amount
		}
	case "USD":
		if to == "IDR" {
			result = amount * usdToIdr
		} else if to == "EUR" {
			result = (amount * usdToIdr) / eurToIdr
		} else if to == "USD" {
			result = amount
		}
	case "EUR":
		if to == "IDR" {
			result = amount * eurToIdr
		} else if to == "USD" {
			result = (amount * eurToIdr) / usdToIdr
		} else if to == "EUR" {
			result = amount
		}
	default:
		fmt.Println("Mata uang gak dikenal.")
		return
	}

	fmt.Printf("Hasil konversi: %.2f %s\n", result, to)
}
