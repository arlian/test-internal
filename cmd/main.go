package main

import (
	"github.com/arlian/test-internal/internal"
	"github.com/arlian/test-internal/pkg/product"
)

func main() {
	println(internal.New())
	product.NewProduct()
}
