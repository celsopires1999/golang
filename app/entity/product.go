package entity

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewProduct () *Product {
	return &Product {
		ID: uuid.NewV4().String(),
	}
}

type Products struct {
	Product []Product 
}

func (p *Products) Add (product Product)  {
	p.Product = append(p.Product, product)
}