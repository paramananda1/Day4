package main

import (
	"fmt"
	"reflect"
)

type S struct {
	str string
	i int
	f float64
	b bool
}
func DeepVsShallowCopy() {
	s := new(S)
	s.str = "paramananda"
	s.i = 10
	s.f = 10.9
	s.b = true

	var sp *S
	sp = s
	fmt.Println("s &s sp &sp", s, &s, sp, &sp)

	var spp **S
	spp = &s
	fmt.Println("s &s spp &spp", s, &s, spp, &spp)

	fmt.Println("Check struct (s,sp) are equal ",reflect.DeepEqual(s, sp))
	fmt.Println("Check struct (s,spp) are equal ",reflect.DeepEqual(s, spp))

	s1 := new(S)
	s1=s
	fmt.Println("Check struct (s,s1) are equal ",reflect.DeepEqual(s, s1))

}
