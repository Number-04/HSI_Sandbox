package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := "root@tcp(localhost:3306)/pegawai_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[x] Gagal menghubungkan database")
		fmt.Println("")
		return nil, err
	}

	fmt.Println("[v] Database berhasil terhubung!")

	return db, nil
}
