package proxy

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestProxy(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	mysql := newProxyMysql()
	go func() {
		for i := 0; i < 100; i++ {
			authCode := rand.Intn(3) - 1
			get := mysql.get(authCode)
			fmt.Printf("1 %d:%s\n", i, get)
		}
	}()

	for i := 0; i < 100; i++ {
		authCode := rand.Intn(3) - 1
		get := mysql.get(authCode)
		fmt.Printf("2 %d:%s\n", i, get)
	}

	group.Done()
}
