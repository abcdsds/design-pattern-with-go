package proxy

import (
	"fmt"
	"sync"
	"time"
)

type datasource interface {
	create() string
	get() string
	delete() string
}

type mysql struct {
}

func (m mysql) create() string {
	return "create"
}

func (m mysql) get() string {
	time.Sleep(1 * time.Second)
	return "get"
}

func (m mysql) delete() string {
	return "delete"
}

type proxyMysql struct {
	mysql *mysql
	lock  *sync.Mutex
}

func newProxyMysql() *proxyMysql {
	return &proxyMysql{
		mysql: &mysql{},
		lock:  new(sync.Mutex),
	}
}

func (m proxyMysql) create(authCode int) string {
	if authCode < 1 {
		return "failed create"
	}

	fmt.Println("success")
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.mysql.create()
}

func (m proxyMysql) get(authCode int) string {
	if authCode < 1 {
		return "failed get"
	}

	fmt.Println("success")
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.mysql.get()
}

func (m proxyMysql) delete(authCode int) string {
	if authCode < 1 {
		return "failed delete"
	}

	fmt.Println("success")
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.mysql.delete()
}
