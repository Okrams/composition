package main

import (
	"fmt"

	"github.com/Okrams/composition/pkg/customer"
	"github.com/Okrams/composition/pkg/invoice"
	"github.com/Okrams/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"México",
		"Puebla",
		customer.New("Eduardo", "Calle 123 #45", "123456789"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Curso de Go", 12.34),
			invoiceitem.New(1, "Curso de POO con GO", 54.23),
			invoiceitem.New(1, "Curso de TEST con GO", 90),
		},
	)
	i.SetTotal()
	fmt.Printf("%+v\n", i)
}
