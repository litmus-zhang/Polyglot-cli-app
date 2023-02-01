package main

import "fmt"

func returnError(value int) error {
	return fmt.Errorf("This is an error with value %v", value)
}

type Foo struct{}

func CreateFoo(fail bool) (Foo, error) {
	if fail {
		return Foo{}, fmt.Errorf("This is an error with value")
	}

}
func (f *Foo) thisIsOnFoo(value int) error {
	return fmt.Errorf("This is an error with value %v", value)
}
func main() {
	err := returnError(5)
	foo := Foo{}
	foo.thisIsOnFoo()

}
