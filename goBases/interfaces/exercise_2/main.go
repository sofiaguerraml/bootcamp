package main

import "fmt"

type Product interface {
	price() float64
}

type small struct {
	cost float64
}

type medium struct {
	cost float64
}

type big struct {
	cost float64
}

const (
	smallType = "PEQUENO"
	medType   = "MEDIUM"
	bigType   = "GRANDE"
)

func main() {
	small := factory(smallType, 10)
	fmt.Println(small.price())

	med := factory(medType, 20)
	fmt.Println(med.price())

	big := factory(bigType, 30)
	fmt.Println(big.price())
}

func (p small) price() float64 {
	return p.cost
}

func (m medium) price() float64 {
	return m.cost * 1.03
}

func (g big) price() float64 {
	return g.cost*1.06 + 2500
}

func factory(productType string, price float64) Product {
	switch productType {
	case smallType:
		return small{cost: price}
	case medType:
		return medium{cost: price}
	case bigType:
		return big{cost: price}
	}
	return nil
}
