package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("===================================")
		fmt.Println("🔥 KONVERSI SUHU CELSIUS 🔥")
		fmt.Println("===================================")

		fmt.Print("Masukkan suhu (°C), lalu tekan ENTER: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		celsius, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("⚠️ Input tidak valid. Harap masukkan angka.")
			fmt.Println("")
			continue
		}

		fahrenheit := (celsius * 9 / 5) + 32
		reamur := celsius * 4 / 5

		fmt.Println("\n💡 Hasil Konversi:")
		fmt.Println("-----------------------------------")
		fmt.Printf("🔥 Fahrenheit : %.2f °F\n", fahrenheit)
		fmt.Printf("💧 Reamur     : %.2f °R\n", reamur)
		fmt.Println("-----------------------------------")

		fmt.Println()
		break
	}
}
