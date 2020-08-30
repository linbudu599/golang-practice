package main

import "fmt"

func main() {
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
		cc = unsafe.Sizeof(aa)
   )

   const (
      ee = iota
      dd
      ff
   )
   


}
