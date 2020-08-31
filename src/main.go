package main

import (
	"fmt"
	"math"
)

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func checkIsCorrectName(name string) bool {
	return len(name) > 1
}

func swap(x, y string) (string, string) {
	return y, x
}

type TCb func(int) int

func testCb(x int, f TCb) {
	f(x)
}

func cb(x int) int {
	fmt.Println("Cb Invoked")
	return x
}

// 方法结构体 类似于接口?
type Circle struct {
	radius float64
}

// 哦哦就是指定从哪里.出来?
//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func main() {

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())

	nextNum := getSequence()

	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())

	nextNum1 := getSequence()

	fmt.Println(nextNum1())
	fmt.Println(nextNum1())
	fmt.Println(nextNum1())
	fmt.Println(nextNum1())
	fmt.Println(nextNum1())

	getSquaRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	testCb(1, cb)
	testCb(2, func(x int) int {
		fmt.Println("Cb Invoked")
		return x
	})

	fmt.Println(getSquaRoot(9))

	fmt.Println(checkIsCorrectName("name"))
	fmt.Println(swap("111", "222"))

	var d bool = true
	var a string = "linbudu"
	var b, c int = 1, 2
	e := false
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("Hello, World!")

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	fmt.Println(Unknown, Female, Male)

	const (
		aa = "abc" + "def"
		bb = len(aa)
		// cc = unsafe.Sizeof(aa)
	)

	const (
		ee = iota
		dd
		ff
	)

	const g bool = false

	if g {
		fmt.Println("芜湖!")
	} else {

	}

	// 暂时跳过 chan 指 channel 是后面并发阶段的只知识
	// var c1, c2, c3 chan int
	// var i1, i2 int
	// select {
	// case i1 = <-c1:
	// 	fmt.Printf("received ", i1, " from c1\n")
	// case c2 <- i2:
	// 	fmt.Printf("sent ", i2, " to c2\n")
	// case i3, ok := (<-c3): // same as: i3, ok := <-c3
	// 	if ok {
	// 		fmt.Printf("received ", i3, " from c3\n")
	// 	} else {
	// 		fmt.Printf("c3 is closed\n")
	// 	}
	// default:
	// 	fmt.Printf("no communication\n")
	// }

	// var loop int = 1

	for loop := 0; loop <= 10; loop++ {
		fmt.Println(loop)
	}

	sum := 1
	// 省略init和post
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 数组?
	strs := []string{"1", "2"}
	for i, s := range strs {
		// index val
		fmt.Println(i, s)
	}

	// 以零值填充多余长度
	nums := [6]int{1, 2, 3, 4}
	for i, x := range nums {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	// 数组语法
	// 不限制数组大小
	var unlimit_arr = [...]int{1, 2, 3, 4}

	fmt.Println(unlimit_arr)

	var nn [10]int

	for i := 0; i < 10; i++ {
		nn[i] = i + 100
	}

	// 指针
	var a1 int = 10

	fmt.Printf("变量的地址: %x\n", &a1)

	var ip *int

	ip = &a1

	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// 空指针
	// var empty_ptr *int

	const max int = 3

	a2 := []int{10, 11, 12}

	var ptr_arr [max]*int

	for i := 0; i < max; i++ {
		// 还是需要手动赋值的
		ptr_arr[i] = &a2[i]
	}

	for i := 0; i < max; i++ {
		fmt.Printf("指针地址 %x\n", ptr_arr[i])
		fmt.Printf("a[%d] = %d\n", i, *ptr_arr[i])
	}

	type Student struct {
		name     string
		age      int
		class_id int
	}

	student := Student{name: "linbudu", age: 10, class_id: 1}

	fmt.Println(student.name)
}
