package urun

type Urun struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Adi   string  `json:"adi"`
	Fiyat float64 `json:"fiyat"`
	Renk  string  `json:"renk"`
}

func NewUrun(adi string, fiyat float64, renk string) *Urun {
	return &Urun{Adi: adi, Fiyat: fiyat, Renk: renk}
}
