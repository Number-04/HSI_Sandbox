package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ReqData struct {
	Lokasi     string
	Celcius    float64
	Fahrenheit float64
	Reamur     float64
}

func KlasifikasiSuhu(suhu float64) string {
	switch {
	case suhu < 18:
		return "Dingin"
	case suhu >= 18 && suhu <= 25:
		return "Hangat"
	default:
		return "Panas"
	}
}

func KonversiSuhu(req *ReqData) {
	req.Fahrenheit = (req.Celcius * 9 / 5) + 32
	req.Reamur = req.Celcius * 4 / 5
}

func IsChar(input string) bool {
	char := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	return char.MatchString(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==============================")
	fmt.Println("#       KONVERTER SUHU       #")
	fmt.Println("==============================")

	var lokasi string
	for {
		fmt.Print("Masukkan lokasi pengukuran suhu : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if !IsChar(input) {
			fmt.Println("!!! Input tidak valid. Hanya masukkan huruf.")
			fmt.Println()
			continue
		}

		lokasi = input
		break
	}

	var suhu float64
	for {
		fmt.Println()
		fmt.Println("* Pakai titik '.' untuk angka desimal")
		fmt.Print("Masukkan suhu Celcius (°C), lalu tekan ENTER: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		suhuFlt, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("!!! Input tidak valid. Harap masukkan angka.")
			fmt.Println()
			continue
		}

		suhu = suhuFlt
		break
	}

	req := ReqData{
		Lokasi:  lokasi,
		Celcius: suhu,
	}

	klasifikasi := KlasifikasiSuhu(req.Celcius)

	KonversiSuhu(&req)

	fmt.Println("")
	fmt.Println("-----------------------------------")
	fmt.Println("           Hasil Konversi          ")
	fmt.Println("-----------------------------------")
	fmt.Printf("=> Lokasi      : %s \n", req.Lokasi)
	fmt.Printf("=> Celcius     : %.2f \n", req.Celcius)
	fmt.Printf("=> Fahrenheit  : %.2f \n", req.Fahrenheit)
	fmt.Printf("=> Reamur      : %.2f \n", req.Reamur)
	fmt.Printf("=> Klasifikasi : %s \n", klasifikasi)
	fmt.Println("-----------------------------------")

	fmt.Println()
}
