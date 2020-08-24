// go build  -gcflags "-m"
//   to see the escape analysis process

package memorytest

type People struct {
	Name string
	Age  int
}

var peoples = map[string]*People{"SP": {Name: "SP", Age: 31}, "Lily": {Name: "Lily", Age: 29}}

func AddPeople(name string, age int) error {
	// escape analysis will output:
	//   &p escapes to heap
	//   moved to heap: p
	p := People{Name: name, Age: age}
	peoples[name] = &p
	return nil
}

func AddPeopleWithNew(name string, age int) error {
	p := new(People)
	p.Name = name
	p.Age = age
	peoples[name] = p
	return nil
}

func ListPeople() (ps []*People) {
	for _, p := range peoples {
		ps = append(ps, p)
	}
	return ps
}
