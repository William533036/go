package main

import "fmt"

func main() {
	m1 := LocalFood{"pigMeat", 38}
	c1 := ImportFood{"France eggplant",100,50}
	c2 := ImportFood{"America fish",200,99}
	e := []Foods{m1, c1, c2}
	totalCal(e)
}

type ImportFood struct {
	 name  string
	 price int
	 tax   int
}


type LocalFood struct {
	 name  string
	 price int

}

type Foods interface {
	calculate() (sum int)
}

func (m LocalFood) calculate() int {
	return m.price
}

func (c ImportFood) calculate() int {
	return c.price + c.tax
}

func totalCal(e []Foods) int{
	expense := 0
	for _, v := range e{
		expense = expense + v.calculate()
	}
	fmt.Printf("total expense is %d", expense)
	return expense
}
