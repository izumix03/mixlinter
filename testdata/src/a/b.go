package a

func main() {
	a := Hoge{ // want "uninitialised field found: Name"
		Test: "",
	}
	println(a)
	println(Hoge{ // want "uninitialised field found: Name" "uninitialised field found: Test"
	})

	c := Fuga{
		"",
		aaa(),
		nil,
	}
	println(c)
	c = Fuga{ // want "uninitialised field found: Stringer" "uninitialised field found: Test"
		Name: aaa(),
	}
	println(c)
}

func aaa() int {
	return 1
}

type Hoge struct {
	Test string
	Name int
}
