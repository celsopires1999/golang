package main

import (
	"errors"
	"fmt"

	"github.com/celsopires1999/golang/entity"
)

func main ()  {
	fmt.Println("Funcionando")

	resultado, err := soma(-1, 20)
	if (err != nil) {
		fmt.Println(err.Error())
	}
	
	fmt.Println(resultado)

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

func soma (a int, b int) (int, error) {
	if (a < 0) {
		return 0, errors.New("a < 0")
	}
	return a + b, nil
}