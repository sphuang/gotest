package structtest

type CarRO struct {
	Price int
	Brand string
}

func (c CarRO) GetPrice() int {
	return c.Price
}

func (c CarRO) GetBrand() string {
	return c.Brand
}

func (c CarRO) SetPrice(price int) {
	c.Price = price
}

func (c CarRO) SetBrand(brand string) {
	c.Brand = brand
}