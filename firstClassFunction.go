package main

import "fmt"
type add func(a int, b int) int

func funcAcceptfunc(a func(a, b int) int, value1 ,value2 int) {
	fmt.Println(a(value1, value2))
}


func firstclassFunction() {
		a := func(a1 int,a2 int) int {
			fmt.Println("I m in \" a := func() int  \" first class function")
			return a1+a2
		}
		b := a(2,3)
		fmt.Printf("%T Return from %d \n\n ", a,b)

	func(n string) {
		fmt.Println("In func(n string) (\"Paramananda\")Welcome", n)
	}("Paramananda")


	fmt.Println("\nUser defined function types Ex. type add func(a int, b int) int \n" +
		"Hete \"add\" is a func type where we can declar a var and call it \n" +
		"Ex.     var funcp add = func(a int, b int) int { return a + b} ; s := funcp(5, 6)")
	var funcp add = func(a int, b int) int {
		return a + b
	}
	s := funcp(5, 6)
	fmt.Println("s := ",s)

	fmt.Println("\n Passing functions as arguments to other functions. \n Ex: funcAcceptfunc(funcp)")

	funcAcceptfunc(funcp,10,20)
	funcAcceptfunc(a,100,200)

}
