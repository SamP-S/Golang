package main

import "fmt"
import "math"
import "reflect"	// type as value

func main() {
	// printing to stdout
	fmt.Println("hello world")
	fmt.Println("C syntax not allowed")
	fmt.Println("hello" + "world")
	fmt.Println("1 + 1 =", 1+1)

	// boolean operations
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!false =", !false)

	// variables
	var x = 7.0
	y := float64(3)
	fmt.Println("x / y =", x / y)
	var a, b, c int = 1, 2, 3
	fmt.Println("a, b, c =", a, b, c)
	f := "apple"
	fmt.Println(f)

	// const
	const n = 1000
	const k = 1234567890 / n
	fmt.Println("1234567890 / n =", k)
	fmt.Println("int64(^^) =", int64(k))
	fmt.Println("sin(n) =", math.Sin(n))

	// for loops
	for i := 5; i < 15; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i := range []int{0, 1, 2} {
		fmt.Println("range", i)
	}

	// while loop
	for {
		if true {
			break
		}
	}

	// if/else
	if "hello" == "hello" {
		fmt.Println("direct string equivalency")
	} else {
		fmt.Println("indirect string equivalency")
	}

	if false && true || true {
		fmt.Println("boolean ops are applied forwards")
	} else {
		fmt.Println("boolean ops are applied backwards")
	}

	// switch
	g := -1
	switch g {
	case 1:
		fmt.Print("a")
		fmt.Println("b")
	case 2:
		fmt.Print("c")
		fmt.Println("d")
	default:
		fmt.Println("default")
	}

	// template function
	templateFunc := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println(i, "= boolean")
		case int:
			fmt.Println(i, "= int")
		default:
			fmt.Println("t =", t)
			fmt.Println(i, "= unspecified type (", reflect.TypeOf(i), ")")
		}
	}
	templateFunc(true)
	templateFunc(5)
	templateFunc("a string")

	// arrays
	var staticList [4]int
	fmt.Println("static list empty =", staticList, "; len =", len(staticList))

	dynamicList := [...]int{1, 2, 5:100, 101}
	fmt.Println("dynamic list =", dynamicList, "; len =", len(dynamicList))

	var mat2 [2][2]int
	mat2 = [2][2]int{{1, 2}, {3, 4}}
	for rowIdx := range mat2 {
		for elemIdx := range mat2[rowIdx] {
			fmtStr := fmt.Sprintf("(%d, %d) = %d", rowIdx, elemIdx, mat2[rowIdx][elemIdx])
			fmt.Println(fmtStr)
		}
	}
	fmt.Printf("mat2 = %v\n", mat2)
	fmt.Printf("mat2 = %#v\n", mat2)
	fmt.Println("program end")

	/// TODO:
	// - make, map
	// - rune
	// - io, Reader, ReadFull
	// - binary, endian
	// - base64, StdEncoding
}
