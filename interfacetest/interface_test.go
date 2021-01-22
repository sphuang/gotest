package interfacetest

import (
	"log"
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEmptyInterface(t *testing.T) {
	type People struct {
		Name string
		Age  int
	}
	p := People{Name: "SP", Age: 31}

	// empty interface can hold values of any type
	var i interface{}
	i = p
	log.Printf("concrete type: %v, %v", reflect.TypeOf(i).PkgPath(), reflect.TypeOf(i).Name()) // github.com/sphuang/gotest/interfacetest, People
	log.Printf("concrete type: %T", i)                                                         // interfacetest.People

	i = 3
	log.Printf("concrete type: %v, %v", reflect.TypeOf(i).PkgPath(), reflect.TypeOf(i).Name()) // "", int
	log.Printf("concrete type: %T", i)                                                         // int
}

func TestImplement(t *testing.T) {
	var vehicle Vehicle

	var car Car
	vehicle = &car
	vehicle.Start()
	vehicle.Move()
	vehicle.Park()

	var bike Bike
	vehicle = &bike
	vehicle.Start()
	vehicle.Move()
	vehicle.Park()

	// compile error: "*House does not implement Vehicle (missing Move method)"
	// var house House
	// vehicle = &house
}

func TestNilReceiverVsNil(t *testing.T) {
	var car *Car
	var vehicle Vehicle = car
	vehicle.Start()

	// var vehicle2 Vehicle
	// panic: runtime error: invalid memory address or nil pointer dereference
	// vehicle2.Start()
}

func IsCar(i interface{}) bool {
	_, ok := i.(*Car)
	return ok
}

func GetTypeString(i interface{}) string {
	switch i.(type) {
	case *Car:
		return "Car"
	case *Bike:
		return "Bike"
	default:
		return "Unknown"
	}
}

func TestTypeAssertion(t *testing.T) {
	var vehicle Vehicle = &Car{}
	assert.True(t, IsCar(vehicle))

	vehicle = &Bike{}
	assert.False(t, IsCar(vehicle))
}

func TestTypeSwitch(t *testing.T) {
	assert.Equal(t, "Car", GetTypeString(&Car{}))
	assert.Equal(t, "Bike", GetTypeString(&Bike{}))
	assert.Equal(t, "Unknown", GetTypeString("just make"))
}