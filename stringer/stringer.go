package main

import "fmt"

type Name struct {
	Name string
}

func (n Name) String() string {
	return fmt.Sprintf("This name sucks: %v", n.Name)
}

func main() {
	p := Name{"Elena"}
	d := Name{"Pudge"}
	fmt.Println(p)
	fmt.Println(d)
}
