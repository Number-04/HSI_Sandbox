package pegawai

import "fmt"

type Pegawai struct {
	Nama        string
	Posisi      string
	GajiBulanan float64
}

func (p Pegawai) HitungGajiTahunan() float64 {
	return p.GajiBulanan * 12
}

type InformasiPegawai interface {
	TampilkanInformasi()
}

func (p Pegawai) TampilkanInformasi() {
	fmt.Printf("Nama         : %s\n", p.Nama)
	fmt.Printf("Posisi       : %s\n", p.Posisi)
	fmt.Printf("Gaji Tahunan : Rp %f\n", p.HitungGajiTahunan())
}
