package c

import (
	"fmt"
	"github.com/izumix03/mixlinter/cmd/b"
)

func main() {
	a := Hoge{ // want "uninitialised field found: Name"
		Test: "",
	}
	fmt.Println(a)
	println(Hoge{ // want "uninitialised field found: Name" "uninitialised field found: Test"
	})

	c := b.Fuga{
		"",
		aaa(),
		nil,
	}
	c = b.Fuga{ // want "uninitialised field found: Stringer" "uninitialised field found: Test"
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
