package db

import "fmt"

type DB interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func doSomething() DB {
	i := T{S: "hello"}
	i.M()
	return i
}
