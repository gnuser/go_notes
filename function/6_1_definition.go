package main

import "fmt"

type N int

func (n N) toString() string{
	return fmt.Sprintf("addr: %p, %#x", &n, n)
}

func (N) test() {
	println("hi")
}

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

// 使用指针前置类型,不会复制对象实例
func (n *N) pointer() {
	(*n)++
	fmt.Printf("p: %p, %v\n", n, *n)
}

func main() {
	var a N = 25
	fmt.Printf("origin addr: %p, %#x\n", &a, a)
	a.test()
	fmt.Println("1", a.toString())
	a.value()
	fmt.Println("2", a.toString())
	a.pointer()
	fmt.Println("3", a.toString())
}
