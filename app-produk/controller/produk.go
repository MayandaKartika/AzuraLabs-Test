package controller

import (
	"app-produk/entity"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type AksesProduct struct {
	DB *gorm.DB
}

func (ap *AksesProduct) AddNewProduct(newProduct entity.Product) entity.Product  {
	err := ap.DB.Create(&newProduct).Error
	if err != nil {
		log.Println(err)
		return entity.Product{}
	}

	return newProduct
}

func (ap *AksesProduct) GetAllProduct() []entity.Product  {
	var daftarProduct = []entity.Product{}
	err := ap.DB.Find(&daftarProduct)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}
	 
	return daftarProduct
}

func (ap *AksesProduct) DetailProductbyName(name string) entity.Product  {
	var nameProduct = entity.Product{}
	err := ap.DB.Select([]string{"id", "name", "price", "rating", "likes"}).Where("name = ?", name).Find(&nameProduct)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return entity.Product{}
	}

	fmt.Println("ID     : ", nameProduct.Id)
	fmt.Println("Nama   : ", nameProduct.Name)
	fmt.Println("Harga  : ", nameProduct.Price)
	fmt.Println("Rating : ", nameProduct.Rating)
	fmt.Println("Likes  : ", nameProduct.Likes)
	
	return nameProduct
}

func (ap *AksesProduct) UpdateProduct(id uint, update entity.Product) bool {
	err := 	ap.DB.Model(entity.Product{}).Where("id = ?", id).Updates(&update)
	if err.Error != nil {
		log.Fatal(err.Error)
		return false
	}
	if aff := err.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang diupdate")
		return false
	}

	return true
}

func (ap *AksesProduct) DeleteProduct(idProduct uint) bool {
	delete := ap.DB.Where("id = ?", idProduct).Delete(&entity.Product{})
	if err := delete.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := delete.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}

	return true
}