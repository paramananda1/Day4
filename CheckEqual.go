package main

import (
	"fmt"
	"reflect"
)

type STRUCT struct {
	X int
	Y string
	Z []int
	M map[string]int
}

func CheckEqualUseReflect() {
	t1 := STRUCT{
		X:1,
		Y:"lei",
		Z:[]int{1,2,3},
		M:map[string]int{
			"a":1,
			"b":2,
		},
	}

	fmt.Println("Struct 1 :",t1)
	t2 :=  STRUCT{
		X:1,
		Y:"lei",
		Z:[]int{1,2,3},
		M:map[string]int{
			"a":1,
			"b":2,
		},
	}
	fmt.Println("Struct 2 :",t2)

	//fmt.Println(t2 == t1)
	//error - invalid operation: t2 == t1 (struct containing []int cannot be compared)


	fmt.Println( "Check reflect.ValueOf(t2) == reflect.ValueOf(t1) : ", reflect.ValueOf(t2) == reflect.ValueOf(t1))
	//false
	fmt.Println("Check reflect.TypeOf(t2) == reflect.TypeOf(t1) : ", reflect.TypeOf(t2) == reflect.TypeOf(t1))
	//true
	//var t3 STRUCT
	//reflect.AppendSlice(t2.Z,5)


	//Update: slice or map
	s1 := []int{1,2,3}
	s2 := []int{1,2,3,4}
	fmt.Println("Slice 1 : Slice 2", s1,s2)
	fmt.Println("Check map are equal ",reflect.DeepEqual(s1, s2))
	fmt.Println("copy s2 to s1")
	copy(s1, s2)
	fmt.Println("[]a1 []a2",s1,s2)

	fmt.Println("Slice 1 : Slice 2", s1,s2)
	fmt.Println("Check map are equal ",reflect.DeepEqual(s1, s2))

	m1 := map[string]int{
		"a":1,
		"b":2,
	}
	m2 := map[string]int{
		"a":1,
		"b":2,
	}
	fmt.Println("map1 : map 2", m1,m2)
	fmt.Println("Check map are equal ",reflect.DeepEqual(m1, m2))
	// m1 == m2 (map can only be compared to nil)
}