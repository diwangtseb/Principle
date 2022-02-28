package main

import "fmt"

type Animals interface {
	run()
}

type Rabbit struct{}

func (r *Rabbit) run() {
	fmt.Println("rabbit run")
}

type Dog struct{}

func (r *Dog) run() {
	fmt.Println("dog run")
}

type Person struct {
	ani Animals
}

func (p *Person) WalkAnimal() {
	fmt.Println("person start walk animal")
	p.ani.run()
}

func main() {
	p := Person{ani: &Rabbit{}}
	p.WalkAnimal()
}
