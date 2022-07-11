package main

import (
	"fmt"
	"sort"
)

func main() {
	product := []struct {
		nama   string
		harga  int
		rating float32
		likes  int
	}{
		{"Indomie", 3000, 5, 150},
		{"Laptop", 4000000, 4.5, 123},
		{"Aqua", 3000, 4, 250},
		{"Smart TV", 4000000, 4.5, 42},
		{"Headphone", 4000000, 3.5, 90},
		{"Very Smart TV", 4000000, 3.5, 87},
	}

	sort.Slice(product, func(a, b int) bool {
		if product[a].harga == product[b].harga {
			return int(product[a].rating) > int(product[b].rating)
		} else if int(product[a].rating) == int(product[b].rating){
			return int(product[a].likes) > int(product[b].likes)
		}
		return product[a].harga < product[b].harga
	})
	fmt.Println("Sebelum dirapikan :\n", product)
	fmt.Println("Setelah dirapikan :")
	fmt.Printf("%+v,\n", product[0:1])
	fmt.Printf("%+v,\n", product[1:2])
	fmt.Printf("%+v,\n", product[2:3])
	fmt.Printf("%+v,\n", product[3:4])
	fmt.Printf("%+v,\n", product[4:5])
	fmt.Printf("%+v,\n", product[5:6])
}