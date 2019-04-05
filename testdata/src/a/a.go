package a

func main() {
	a := Hoge{ // want "uninitialised field found: Name"
		Test: "",
	}
	println(a)
	println(Hoge{ // want "uninitialised field found: Name" "uninitialised field found: Test"
	})
}

type Hoge struct {
	Test string
	Name int
}
