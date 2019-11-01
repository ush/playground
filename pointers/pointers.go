package pointers

import "fmt"

func PrintPointee(p *int) {
	fmt.Println(*p)
}
