package pegawai

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Pegawai struct {
	ID          uint `gorm:"primaryKey"`
	Nama        string
	Posisi      string
	GajiBulanan float64
}

func (p Pegawai) GajiTahunan() float64 {
	return p.GajiBulanan * 12
}

func Migrate(DB *gorm.DB) {
	DB.Exec("TRUNCATE TABLE pegawais")

	fmt.Println("[v] Migrasi database")
	DB.AutoMigrate(&Pegawai{})
}

func AllPegawai(DB *gorm.DB) {
	fmt.Println("\n=== List Pegawai ===")
	var listPegawai []Pegawai
	DB.Find(&listPegawai)

	for _, p := range listPegawai {
		fmt.Printf(
			"ID: %d | Nama: %s | Posisi: %s | GajiBulanan: %.0f | GajiTahunan: %.0f\n",
			p.ID, p.Nama, p.Posisi, p.GajiBulanan, p.GajiTahunan(),
		)
	}
}

func Store(DB *gorm.DB) {
	pegawai := []Pegawai{
		{Nama: "Agung", Posisi: "Satpam", GajiBulanan: 3000000},
		{Nama: "Budi", Posisi: "Admin", GajiBulanan: 4000000},
		{Nama: "Citra", Posisi: "Kasir", GajiBulanan: 3500000},
		{Nama: "Dedi", Posisi: "OB", GajiBulanan: 2500000},
		{Nama: "Eka", Posisi: "Manager", GajiBulanan: 7000000},
	}

	if err := DB.Create(&pegawai).Error; err != nil {
		log.Fatal("Gagal insert data pegawai:", err)
	}
	fmt.Println("[v] Data pegawai berhasil ditambahkan")

}

func Update(DB *gorm.DB) {
	fmt.Println("\n=== Update Gaji Pegawai dengan ID 2 ===")
	var peg Pegawai
	if err := DB.First(&peg, 2).Error; err == nil {
		peg.GajiBulanan = 5000000
		DB.Save(&peg)
		fmt.Printf("Updated: ID: %d | Nama: %s | Posisi: %s | GajiBulanan: %.0f | GajiTahunan: %.0f\n",
			peg.ID, peg.Nama, peg.Posisi, peg.GajiBulanan, peg.GajiTahunan())
	} else {
		fmt.Println("Pegawai dengan ID 2 tidak ditemukan")
	}
}

func Delete(DB *gorm.DB) {
	fmt.Println("\n=== Hapus Pegawai dengan ID 3 ===")
	DB.Delete(&Pegawai{}, 3)
}
