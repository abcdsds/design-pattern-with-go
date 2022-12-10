package abstractfactory

import "fmt"

type maker interface {
	Memory() memory
	Drive() drive
}

type samsung struct{}

func (m samsung) Memory() memory {
	return samsungMemory(1024)
}

func (m samsung) Drive() drive {
	return samsungDrive(32)
}

type hynix struct{}

func (m hynix) Memory() memory {
	return hynixMemory(1024)
}

func (m hynix) Drive() drive {
	return hynixDrive(64)
}

func NewMaker(name string) maker {
	switch name {
	case "samsung":
		return samsung{}
	case "hynix":
		return hynix{}
	default:
		return nil
	}
}

type memory interface {
	Size() string
}

type samsungMemory int

func (m samsungMemory) Size() string {
	return fmt.Sprint(m, "GB")
}

type hynixMemory int

func (m hynixMemory) Size() string {
	return fmt.Sprint(m, "GB")
}

type drive interface {
	Capacity() string
}

type samsungDrive int

func (d samsungDrive) Capacity() string {
	return fmt.Sprint(d, "TB")
}

type hynixDrive int

func (d hynixDrive) Capacity() string {
	return fmt.Sprint(d, "TB")
}
