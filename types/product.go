package types

type Product struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float32 `json:"price,omitempty"`
}

type InsertProductRequest struct {
	Name  string  `json:"name,omitempty" validate:"required"`
	Price float32 `json:"price,omitempty" validate:"required,gte=0"`
}

type ProductsRequest struct {
	Offset int `json:"offset,omitempty" validate:"gte=0"`
	Limit  int `json:"limit,omitempty" validate:"gte=1"`
}

func NewProduct(name string, price float32) Product {
	return Product{
		Name:  name,
		Price: price,
	}
}
