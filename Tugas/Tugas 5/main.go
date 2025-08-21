package main

import "aplikasi_karyawan/pegawai"

func main() {
	karyawan := pegawai.Pegawai{
		Nama:        "Agung Irawan",
		Posisi:      "Staff",
		GajiBulanan: 1000000,
	}

	var info pegawai.InformasiPegawai = karyawan
	info.TampilkanInformasi()

}
