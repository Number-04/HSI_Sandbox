package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//deklarasi reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==============================")
	fmt.Println("#       KONVERTER SUHU       #")
	fmt.Println("==============================")

	//Perulangan untuk mengulangi input tidak angka
	for {
		fmt.Println("* Pakai titik '.' untuk angka desimal")
		fmt.Print("Masukkan suhu Celcius (°C), lalu tekan ENTER: ")

		input, _ := reader.ReadString('\n')

		/*
			//Mengantipasi jika stdin error dalam input
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("XXX Terjadi kesalahan saat membaca input:", err)
				return
				//dan program berhenti
			}

		*/

		//menghilangkan spasi
		input = strings.TrimSpace(input)
		//input = strings.ReplaceAll(input," ","")

		celsius, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("!!! Input tidak valid. Harap masukkan angka.")
			fmt.Println()
			//menggulangi inputan angka
			continue

		}

		fahrenheit := (celsius * 9 / 5) + 32
		reamur := celsius * 4 / 5

		fmt.Println("")
		fmt.Println("-----------------------------------")
		fmt.Println("           Hasil Konversi          ")
		fmt.Println("-----------------------------------")
		fmt.Printf("=> Fahrenheit : %.2f \n", fahrenheit) // disini saya pakai format 2 digit desimal untuk outputnya
		fmt.Printf("=> Reamur     : %.2f \n", reamur)
		fmt.Println("-----------------------------------")

		fmt.Println()

		//Mengakhiri program
		break
	}
}
