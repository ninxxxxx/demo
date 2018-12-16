package main

import (
	"fmt"

	"github.com/ninxxxxx/demo/example"
)

func main() {

	println("Hello, Go")

	fmt.Println(func() string {
		return "data from anonymous function"
	}())

	//firstclass function
	counter := counterFactory()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	//benedic function
	fmt.Println("hello", "world", "guys")

	//exercise 1
	example.CombineStr("hello", "world")

	a, b := example.SwapStr("hi", "guys")
	fmt.Println(a, b)

	//array is an immutable
	var r [2]string
	r[0] = "arnon"
	r[1] = "kaewprasert"
	fmt.Println(r[0])
	fmt.Println(r[1])
	fmt.Println(r)

	primes := [...]int{1, 3, 4, 2, 5, 6, 6}
	fmt.Println(primes)

	//slice
	var s []string
	s = make([]string, 0, 0) //init val
	// s := make([]string, 0, 0) //init val
	// s := []string{}
	// s := []string{"1","2","3"}
	s = append(s, "1", "2", "3")
	s = append(s, "4")
	fmt.Println(s)
	fmt.Println(s[1:])
	fmt.Println(s[1:3])
	fmt.Println(s[:3])
	fmt.Println("len is", len(s))
	fmt.Println("cap is", cap(s))

	//use _ for ignor
	// for i, v := range s {
	for _, v := range s {
		fmt.Println(v)
	}

	//map
	var m map[string]string
	m = make(map[string]string)
	if m == nil {
		fmt.Println("it's a nill")
	}
	fmt.Println(m)

	m["a"] = "aaa"
	m["b"] = "bbb"
	m["c"] = "ccc"
	fmt.Println(m)

	delete(m, "b")
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	if v, ok := m["z"]; ok {
		fmt.Println(v)

	}

	//defer จงทำสิ่งนี้ก่อนจบ func
	//สั่งทีหลังจะทำก่อน
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	doSomething(4)

	//struct

	rec := rectangle{
		Width:  3,
		Length: 1,
	}
	rec2 := rectangle{3, 4}
	fmt.Println(rec.area())
	fmt.Println(rec2.area())
	// rec.

	//naming interface with ...er
	var i interface{} = 2
	val, yes := i.(int)
	fmt.Println(i, val, yes)

	//GoRoutine

}

type rectangle struct {
	Width  int
	Length int
}

func (rec rectangle) area() int {
	return rec.Width * rec.Length
}

//implicit implement , rectangle auto implement areaer
//usecase for test database can implement interface to be like data from database
type areaer interface {
	area() int
}

//new type
type getterFunc func() int

//just alias
type aliasgetterFunc = func() int

//cloture functino
func counterFactory() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func init() {
	fmt.Println("initial1")
}

func init() {
	fmt.Println("initial0")
}

func doSomething(n int) {
	//assign value first but show before
	defer fmt.Println(n)
	defer func() {
		fmt.Println(n)
	}()

	n = n * n
	fmt.Println(n)
	//16
	//16
	//4
}
