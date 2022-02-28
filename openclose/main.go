package main

import "fmt"

type ICar interface {
	// car's name
	GetName() string
	GetPrice() int
}

type Mercedes struct {
	Name  string
	price int
}

func (m *Mercedes) GetName() string {
	return m.Name
}

func (m *Mercedes) GetPrice() int {
	return m.price
}

//requirement change start ..
type FinanceMercedes struct {
	Mercedes
}

func (f *FinanceMercedes) GetPrice() int {
	if f.price > 100 {
		return f.price + 2
	} else {
		return f.price
	}
}

//requirement change end ..

func main() {
	cars := []ICar{}
	cars = append(cars, &FinanceMercedes{Mercedes{"amg", 130}})
	cars = append(cars, &FinanceMercedes{Mercedes{"maybach", 20}})
	for _, v := range cars {
		fmt.Println("car's name is", v.GetName(), "\t price is", v.GetPrice())
	}
}
