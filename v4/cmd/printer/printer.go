package printer

import (
	"fmt"
)

type Printable interface {
	ToJSON() string
	ToTable() string
}

type Printer interface {
	Print(p Printable)
}

type JSONPrinter struct{}

func (jp JSONPrinter) Print(p Printable) {
	fmt.Println(p.ToJSON())
}

type TablePrinter struct{}

func (tp TablePrinter) Print(p Printable) {
	fmt.Println(p.ToTable())
}
