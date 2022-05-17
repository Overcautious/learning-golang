package main

import (
	"fmt"
	"unsafe"
)

func funcVec() {
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 当出现 "a : b "的时候，a表示索引，b表示值，未表示的地方自动赋0元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 下标4的元素为5，后续跟着下标4赋值，所以元素为 1, 2, 0, 0, 5, 6
	// var b = [...]int{5: 1, 2: 2, 2, 3, 4} // 错误的赋值方式，因为下标2往后的第三个数据，也就是下标5的位置，在前面已经赋值过
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Println(unsafe.Sizeof(a))

	var times [5][0]int
	fmt.Println(len(times))
	fmt.Println(unsafe.Sizeof(times))

}

func funcString() {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println(hello)
	fmt.Println(world)
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("%#v\n", []byte("Hello, 世界"))
	fmt.Println("\xe4\xb8\x96") // 打印: 世
	fmt.Println("\xe7\x95\x8c") // 打印: 界

	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

}

func main() {
	funcVec()
	funcString()

	var str = []byte{'2', '2', ' ', '4', '5'}
	b := str[:0]
	fmt.Printf("len(b)=%d, cap(b)=%d", len(b), cap(b))
	for _, x := range str {
		if x != ' ' {
			b = append(b, x)
		}
	}

	fmt.Printf("%v\n", str)
	fmt.Printf("%v\n", b)

	for x := 0; x <= 9; x++ {
		for y := x + 1; y <= 9; {
			for m := y; m <= 9; {
				for n := m; n <= 9; {
					if x+y+m+n == 30 {
						fmt.Println(x, y, m, n)
					}
					n++
				}
				m++
			}
			y++
		}

	}

	fmt.Println("Hello, world")
}
