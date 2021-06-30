package main

import (
	"fmt"

	"github.com/celsopires1999/golang/entity"
)

func main ()  {
	fmt.Println("Funcionando")

	products := entity.Products{}	

	product := entity.Product{}
	product.ID = "abc"
	product.Name = "Fusca"

	product2 := entity.NewProduct()
	product2.Name = "BMW"

	products.Add(product)
	products.Add(*product2)

	fmt.Println(products)

}