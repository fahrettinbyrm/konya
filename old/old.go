package old

import "gorm.io/gorm"

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}
func (p *Product) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Product{})
}

func (p *Product) Save(db *gorm.DB) error {
	return db.Model(&Product{}).Save(p).Error
}

func (p *Product) FetchByID(db *gorm.DB, id uint) {
	tempProduct := Product{}
	db.Model(&Product{}).First(&tempProduct, id)
	*p = tempProduct
}
