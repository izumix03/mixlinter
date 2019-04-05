package a

import (
	"fmt"
)

func main() {
	a := Hoge{ // want "uninitialised field found: Name"
		Test: "",
	}
	fmt.Println(a)
	println(Hoge{ // want "uninitialised field found: Name" "uninitialised field found: Test"
	})

	c := Fuga{
		"",
		aaa(),
		nil,
	}
	c = Fuga{ // want "uninitialised field found: Stringer" "uninitialised field found: Test"
		Name: aaa(),
	}
	fmt.Println(c)
}

func aaa() int {
	return 1
}

type Hoge struct {
	Test string
	Name int
}
