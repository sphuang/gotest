package interfacetest

import (
	"log"
	"reflect"
	"testing"
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
