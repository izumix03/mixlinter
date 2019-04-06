package b

import (
	"fmt"
	"github.com/izumix03/mixlinter/cmd/c/d"
)

type Fuga struct {
	Test string
	Name int
	fmt.Stringer
	Hoge *d.Hoge2
}
