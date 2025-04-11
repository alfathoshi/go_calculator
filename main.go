package main

import "fmt"

func main() {
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
