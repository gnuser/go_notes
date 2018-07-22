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

type X struct {}

func (x *X) test() {
	println("hi!", x)
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
	fmt.Println("===========")
	// 自动使用对应类型
	p := &a
	p.value()
	fmt.Println("4", a.toString())
	p.pointer()
	fmt.Println("5", a.toString())

	// 不能用多级指针
	//p2 := &p
	//p2.value() // calling method value with receiver p2 (type **N) requires explicit dereference
	//p2.pointer() // calling method pointer with receiver p2 (type **N) requires explicit dereference

	fmt.Println("===========")
	var b *X
	b.test()     // 还没有初始化的指针,也可以调用
	//X{}.test() // 不是指针不行
	(&X{}).test()

}
