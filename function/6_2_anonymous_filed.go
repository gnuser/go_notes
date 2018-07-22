package main

import "sync"

type data struct {
	sync.Mutex
	buf [1024]byte
}

type user struct {

}

type manager struct {
	user
}

func (user) toString() string {
	return "user"
}

func (m manager) toString() string {
	return m.user.toString() + "; manager"
}

func main() {
	d := data {}
	d.Lock()
	defer d.Unlock()

	var m manager
	println(m.toString())
	println(m.user.toString())
}