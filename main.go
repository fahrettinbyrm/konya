package main

import (
	"fmt"
	"github.com/fahrettinbyrm/konya/database"
	"github.com/fahrettinbyrm/konya/urun"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		fmt.Println("Database Connection Hatası:", err)
		return
	}
	urunRepo, err := urun.NewUrunRepository(db)
	if err != nil {
		fmt.Println("Urun Repo Hatası: ", err)
		return
	}

	laptop := urun.NewUrun("MacBook Pro", 44000, "gri")
	telefon := urun.NewUrun("İphone 13 Pro Max", 37000, "siyah")

	fmt.Println("Ürün ekleniyor... HATA: ", urunRepo.Create(laptop))
	fmt.Println("Ürün ekleniyor... HATA: ", urunRepo.Create(telefon))
	//
}
