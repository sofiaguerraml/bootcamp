package main

import "fmt"

type Producto interface {
	precio() float64
}

type pequeno struct {
	costo float64
}

type mediano struct {
	costo float64
}

type grande struct {
	costo float64
}

const (
	peqType = "PEQUENO"
	medType = "MEDIANO"
	graType = "GRANDE"
)

func main() {
	peq := factory(peqType, 10)
	fmt.Println(peq.precio())

	med := factory(medType, 20)
	fmt.Println(med.precio())

	gra := factory(graType, 30)
	fmt.Println(gra.precio())
}

func (p pequeno) precio() float64 {
	return p.costo
}

func (m mediano) precio() float64 {
	return m.costo * 1.03
}

func (g grande) precio() float64 {
	return g.costo*1.06 + 2500
}

func factory(productType string, precio float64) Producto {
	switch productType {
	case peqType:
		return pequeno{costo: precio}
	case medType:
		return mediano{costo: precio}
	case graType:
		return grande{costo: precio}
	}
	return nil
}
