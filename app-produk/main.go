package main

import (
	"app-produk/config"
	"app-produk/controller"
	"app-produk/entity"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func input() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return strings.TrimSuffix(input, "\n")
}

func main() {
	
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error: failed to load file env file")
	}

	conn := config.InitDB()
	config.MigrateDB(conn)
	aksesProduct := controller.AksesProduct{DB: conn}
	fmt.Println("\t    SELAMAT DATANG")
	var id uint
	var pilih int = 0
	for pilih != 6 {
		fmt.Println("\t Silahkan Pilih Menu")
		fmt.Println("1. Daftar Semua Produk")
		fmt.Println("2. Detail Produk")
		fmt.Println("3. Tambah Produk")
		fmt.Println("4. Edit Produk")
		fmt.Println("5. Hapus Produk")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih Menu : ")
		fmt.Scanln(&pilih)
		fmt.Print("\n")

		switch pilih {
		case 1:
			fmt.Println("\t Daftar Semua Produk")
			fmt.Println("====================================")
			var allProduct = aksesProduct.GetAllProduct()
			for _, p := range allProduct {
				fmt.Printf(" Id     : %d \n", p.Id)
				fmt.Printf(" Nama   : %s \n", p.Name)
				fmt.Printf(" Harga  : %d \n", p.Price)
				fmt.Printf(" Rating : %.1f \n", p.Rating)
				fmt.Printf(" Likes  : %d \n", p.Likes)
				fmt.Println(" ===================================")
			}
			fmt.Print("\n")
		case 2:
			fmt.Println("\t Detail Produk")
			var detailProduct entity.Product
			fmt.Print("Masukkan Nama Produk : ")
			detailProduct.Name = input()
			detail := aksesProduct.DetailProductbyName(detailProduct.Name)
			if detail.Name == "" {
				fmt.Println("Produk tidak tersedia *~*")
				break
			} else {
				fmt.Println("Produk ditemukan *v*")
			}
			fmt.Print("\n")
		case 3:
			fmt.Println("\t Tambah Produk")
			var newProduct entity.Product
			fmt.Print("Masukkan Nama  :")
			newProduct.Name = input()
			fmt.Print("Masukkan Harga :")
			fmt.Scanln(&newProduct.Price)
			fmt.Print("Masukkan Rating:")
			fmt.Scanln(&newProduct.Rating)
			fmt.Print("Masukkan Likes :")
			fmt.Scanln(&newProduct.Likes)
			result := aksesProduct.AddNewProduct(newProduct)
			if result.Name == "" {
				fmt.Println("Gagal tambah produk, ada error *~*")
				break
			} else {
				fmt.Println("Berhasil tambah produk *v*")
				id = result.Id
				fmt.Println("Id Produk:", result.Id)
			}
			fmt.Print("\n")
		case 4:
			fmt.Println("\t Update Produk")
			var updateProduct entity.Product 
			fmt.Print("Id Product:")
			fmt.Scanln(&id)
			fmt.Print("Ubah Nama:")
		    updateProduct.Name = input()
			fmt.Print("Ubah Harga:")
		    fmt.Scanln(&updateProduct.Price)
			fmt.Print("Ubah Rating:")
		    fmt.Scanln(&updateProduct.Rating)
			fmt.Print("Ubah Likes:")
		    fmt.Scanln(&updateProduct.Likes)
			res := aksesProduct.UpdateProduct(id, updateProduct)
			if !res {
			    fmt.Println("Gagal update produk")
			    break
		    } else {
				  	fmt.Println("Produk telah diupdate")
				}
			fmt.Print("\n")
			case 5:
			var ProductID uint
			fmt.Print("Masukkan id produk yang akan dihapus :")
			fmt.Scanln(&ProductID)
			fmt.Println(aksesProduct.DeleteProduct(ProductID))
			fmt.Print("\n")
		}
	}
}