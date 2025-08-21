package main

import (
	"aplikasi_karyawan/pegawai"
	"time"

	"github.com/brianvoe/gofakeit"
)

func main() {

	gofakeit.Seed(time.Now().UnixNano())

	var ListPegawai []pegawai.Pegawai

	for i := 0; i < 10; i++ {
		ListPegawai = append(ListPegawai, pegawai.Pegawai{
			Nama:        gofakeit.Name(),
			Posisi:      gofakeit.JobTitle(),
			GajiBulanan: gofakeit.Price(1000000, 5000000),
		})
	}

	data := pegawai.ListPegawai{List: ListPegawai}

	var info2 pegawai.InformasiPegawaiV2 = data
	info2.TampilkanInformasiV2()

}
