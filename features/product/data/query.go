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
func (pd *productData) DeleteItemData(productID int) bool {
	res := pd.db.Where("ID = ?", productID).Delete(&Product{})
	if res.Error != nil {
		log.Println("Cannot delete data", res.Error.Error())
		return false
	}

	if res.RowsAffected < 1 {
		log.Println("No data deleted", res.Error.Error())
		return false
	}

	return true
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
func (pd *productData) PostItemData(newProduct domain.Product) domain.Product {
	var product = FromModel(newProduct)
	err := pd.db.Create(&product).Error

	if product.ID == 0 {
		log.Println("Invalid ID")
		return domain.Product{}
	}

	if err != nil {
		log.Println("Cant create product object", err.Error())
		return domain.Product{}
	}

	return product.ToModel()
}

// UpdateItemData implements domain.ProductData
func (pd *productData) UpdateItemData(newProduct domain.Product, productID int) domain.Product {
	var product = FromModel(newProduct)
	err := pd.db.Model(&Product{}).Where("ID = ?", productID).Updates(product)

	if err.Error != nil {
		log.Println("Cant update product object", err.Error.Error())
		return domain.Product{}
	}

	if err.RowsAffected == 0 {
		log.Println("Data Not Found")
		return domain.Product{}
	}

	return product.ToModel()
}
