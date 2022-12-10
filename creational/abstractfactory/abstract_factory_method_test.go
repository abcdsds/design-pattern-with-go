package abstractfactory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	samsung := NewMaker("samsung")
	d := samsung.Drive()
	fmt.Println("samsung drive capa", d.Capacity())
	m := samsung.Memory()
	fmt.Println("samsung memory size", m.Size())
	hynix := NewMaker("hynix")
	d = hynix.Drive()
	fmt.Println("hynix drive capa", d.Capacity())
	m = hynix.Memory()
	fmt.Println("hynix memory size", m.Size())

	// output:
	// samsung drive capa 32TB
	// samsung memory size 1024GB
	// hynix drive capa 64TB
	// hynix memory size 1024GB
}
