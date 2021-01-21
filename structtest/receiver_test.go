package structtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueReceiver(t *testing.T) {
	car := CarRO{ Price: 300, Brand: "BMW" }
	assert.Equal(t, 300, car.GetPrice())
	assert.Equal(t, "BMW", car.GetBrand())

	car.SetPrice(500)
	car.SetBrand("Porsche")
	// below test will be failed, because car's value doesn't change
	// assert.Equal(t, 500, car.GetPrice())
	// assert.Equal(t, "Porsche", car.GetBrand())
}

func TestPointerReceiver(t *testing.T) {
	car := Car{ Price: 300, Brand: "BMW" }
	assert.Equal(t, 300, car.GetPrice())
	assert.Equal(t, "BMW", car.GetBrand())

	(&car).SetPrice(500)
	(&car).SetBrand("Porsche")
	assert.Equal(t, 500, car.GetPrice())
	assert.Equal(t, "Porsche", car.GetBrand())

	car.SetPrice(50)
	car.SetBrand("Toyota")
	assert.Equal(t, 50, car.GetPrice())
	assert.Equal(t, "Toyota", car.GetBrand())
}