package a

import (
	"fmt"
	"time"
)

type Fuga struct {
	Test string
	Name int
	fmt.Stringer
}

type Hoga struct {
	Test2 string
	Test3 string
	*Fuga
	*time.Time
}

func construct() {
	fmt.Printf("Hoga:%+v\n", Hoga{
		Test2: "test",
		Fuga: &Fuga{
			Test:     "",
			Name:     0,
			Stringer: nil,
		},
	})
}
