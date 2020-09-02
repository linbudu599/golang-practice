package main

import (
	"fmt"
	"math"
	"time"
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

	// 初始化长度 与 容量
	s := make([]int, 5, 10)
	fmt.Println(s) // 均为零值

	arr1 := []int{1, 2, 3, 4, 5}

	// 数组的引用
	s1 := arr1[:]
	fmt.Println(s1)

	// 或者是[:end] / [start:]
	s2 := arr1[1:3]
	fmt.Println(s2)

	var numbers = make([]int, 3, 5)

	printSlice(numbers)

	var numbers1 []int

	if numbers1 == nil {
		fmt.Println("切片是空的")
	}

	fmt.Println("numbers[1:4] ==", numbers[1:4])

	var numbers3 []int

	numbers3 = append(numbers3, 0)
	numbers3 = append(numbers3, 1)
	numbers3 = append(numbers3, 2, 3, 4)

	numbers4 := make([]int, len(numbers3), (cap(numbers3) * 2))
	printSlice(numbers4)

	// 3 -> 4
	copy(numbers4, numbers3)
	printSlice(numbers4)

	var map1 map[string]string

	map1 = make(map[string]string)

	map1["foo1"] = "bar1"
	map1["foo2"] = "bar2"

	for item := range map1 {
		fmt.Println(item, map1[item])
	}

	// ok代表是否存在
	foo1, ok := map1["foo1"]

	fmt.Println(foo1, ok)

	delete(map1, "foo1")

	for idx, item := range numbers3 {
		fmt.Println(idx, item)
	}

	var sum1 int = 17
	var count1 int = 5
	var mean1 float32

	mean1 = float32(sum1) / float32(count1)

	fmt.Println(mean1)

	var phone Phone

	phone = new(NokiaPhone)
	phone.call(1)

	phone = new(IPhone)
	phone.call(2)

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

	go say("wuhu!")
	say("yahu!")

	ch1 := make(chan int)

	s21 := []int{1, 2, 3, 4, -6, 9, 67, 3, 56}

	go sum111(s21[:len(s1)/2], ch1)
	go sum111(s21[len(s1)/2:], ch1)

	x, y := <-ch1, <-ch1

	fmt.Println(x, y, x+7)

	ch := make(chan int, 2)

	// 可以同时发送两个而不用读取
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	c_group := make(chan int, 10)
	go fibonacci(cap(c_group), c_group)
	// 如果c不关闭 range就不会结束
	for i := range c_group {
		fmt.Println(i)
	}

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 接收完10个数据后关闭通道
	close(c)
}

func sum111(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func say(s string) bool {
	for i := 0; i < len(s); i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	return len(s) > 0
}

type Phone interface {
	call(i int)
}

type NokiaPhone struct{}

type IPhone struct{}

func (nokiaPhone NokiaPhone) call(i int) {
	fmt.Println("I am Nokia, I can call you!")
}

func (iPhone IPhone) call(i int) {
	fmt.Println("I am iPhone, I can call you!")
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
