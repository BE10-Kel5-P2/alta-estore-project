package data

import (
	"altaproject2/domain"
	"log"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductData {
	return &productData{
		db: db,
	}
}

// DeleteItemData implements domain.ProductData
func (*productData) DeleteItemData() {
	panic("unimplemented")
}

// GetAllItemData implements domain.ProductData
func (pd *productData) GetAllItemData() []domain.Product {
	var products []Product

	err := pd.db.Find(&products)

	if err.Error != nil {
		log.Println("cant get all data from db", err.Error)
		return nil
	}

	return ParseToArr(products)
}

// GetItemData implements domain.ProductData
func (pd *productData) GetItemData(productID int) (domain.Product, error) {
	var tmp Product
	err := pd.db.Where("ID = ?", productID).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return domain.Product{}, err
	}

	return tmp.ToModel(), nil
}

// PostItemData implements domain.ProductData
func (*productData) PostItemData() {
	panic("unimplemented")
}

// UpdateItemData implements domain.ProductData
func (*productData) UpdateItemData() {
	panic("unimplemented")
}
