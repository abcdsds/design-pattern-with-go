package singleton

import (
	"fmt"
	"sync"
)

type screenConfig struct {
	width  int
	height int
}

var once sync.Once
var config *screenConfig

func setUpScreenConfig(width, height int) {
	once.Do(func() {
		if config == nil {
			config = &screenConfig{
				width:  width,
				height: height,
			}
			fmt.Println("setup screen config")
		}
	})
}
