package main

import "fmt"

func main() {
	var number1, number2 int
	var operator string

	fmt.Print("Masukkan angka pertama : ")
	fmt.Scanln(&number1)
	fmt.Print("Masukkan operator (+,-,*,/,%) : ")
	fmt.Scanln(&operator)
	fmt.Print("Masukkan angka kedua : ")
	fmt.Scanln(&number2)
	hasil := 0
	if number1 <= 1000000 && number2 <= 1000000 {
			switch operator {
			case "+":
				hasil = number1 + number2
				fmt.Print("\n")
			case "-":
				hasil = number1 - number2
				fmt.Print("\n")
			case "*":
				hasil = number1 * number2
				fmt.Print("\n")
			case "/":
				hasil = number1 / number2
				fmt.Print("\n")
			case "%":
				hasil = number1 % number2
				fmt.Print("\n")
			}
	} else {
		fmt.Print("Angka tidak valid, lebih dari 1000000")
	}
	fmt.Printf("%d \n", hasil)
	fmt.Printf("%d %s %d = %d", number1,operator,number2,hasil)
}