package types

type Product struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float32 `json:"price,omitempty"`
}

func NewProduct(name string, price float32) Product {
	return Product{
		Name:  name,
		Price: price,
	}
}
