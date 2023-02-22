package main


type Product struct {
	ID           int  `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	OriginalPrice float64 `json:"orginal_price"`
	CategoryID   int     `json:"category_id"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	ProductID int    `json:"product_id"`
}
