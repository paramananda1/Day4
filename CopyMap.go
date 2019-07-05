package main

import (
	"fmt"
	"reflect"
)

func CopyMap() {
	var MAP1,MAP2 map[string]int
	MAP1= map[string]int{"A":1 , "B": 2, "C": 3}

	fmt.Println("Test Copy Map to other:")
	fmt.Println("Map 1",MAP1)

	MAP2 = make(map[string]int)
	for k,v := range MAP1 {
		MAP2[k] = v
	}
	fmt.Println("Map 2 Afrter copy",MAP2)
	fmt.Println("are they eqal : ",reflect.DeepEqual(MAP1, MAP2))

}
