package factorymethod

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	i := newItem("sword")
	fmt.Println(i.Name())

	i = newItem("wand")
	fmt.Println(i.Name())

	i = newItem("axe")
	fmt.Println(i.Name())

	i = newItem("potion")
	fmt.Println(i.Name())

	i = newItem("1234")
	fmt.Println(i.Name())

	// output:
	// sword
	// wand
	// axe
	// potion
	// none
}
