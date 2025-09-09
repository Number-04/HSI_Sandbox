package main

import (
	"go_gorm_tes/database"
	"go_gorm_tes/pegawai"
	"log"
)

func main() {

	DB, err := database.ConnectDB()
	if err != nil {
		log.Fatalln("[x] Database error!")
	}

	//MIGRATE
	pegawai.Migrate(DB)

	//STORE
	pegawai.Store(DB)
	pegawai.AllPegawai(DB)

	//UPDATE
	pegawai.Update(DB)
	pegawai.AllPegawai(DB)

	//DELETE
	pegawai.Delete(DB)
	pegawai.AllPegawai(DB)
}
