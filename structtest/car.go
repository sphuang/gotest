package structtest

type Car struct {
	Price int
	Brand string
}

func (c *Car) GetPrice() int {
	return c.Price
}

func (c *Car) GetBrand() string {
	return c.Brand
}

func (c *Car) SetPrice(price int) {
	c.Price = price
}

func (c *Car) SetBrand(brand string) {
	c.Brand = brand
}