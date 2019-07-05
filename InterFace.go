package main
import "fmt"

// interface general
type Animal interface {
	WhoIam() string
}

type Dog struct {
	d string
}

func (d Dog) WhoIam() string {
	return d.d
}

type Cat struct {
	c string
}

func (c Cat) WhoIam() string {
	return c.c
}

type Llama struct {
	l string
}

func (l Llama) WhoIam() string {
	return l.l
}

type JavaProgrammer struct {
	j string
}

func (j JavaProgrammer) WhoIam() string {
	return j.j
}

// Empty interface

type AnimalEmpty interface {
	WhoIam() interface{}
}
type DogEmpty struct {
	d interface{}
}

func (d DogEmpty) WhoIam() interface{} {
	return d.d
}

type CatEmpty struct {
	c interface{}
}

func (c CatEmpty) WhoIam() interface{} {
	return c.c
}

type LlamaEmpty struct {
	l interface{}
}

func (l LlamaEmpty) WhoIam() interface{} {
	return l.l
}

type JavaProgrammerEmpty struct {
	j interface{}
}

func (j JavaProgrammerEmpty) WhoIam() interface{} {
	return j.j
}

//////main function
func InterFaceMain(isEmpty bool) {
	if !isEmpty {
		str := "I am a "
		animals := []Animal{Dog{str}, Cat{str}, Llama{str}, JavaProgrammer{str}}
		for _, animal := range animals {
			fmt.Printf("%v %T \n",animal.WhoIam(), animal)
		}
	}else {
		var str interface{}
		str = "I am a "
		animals := []AnimalEmpty{DogEmpty{}, CatEmpty{}, LlamaEmpty{}, JavaProgrammerEmpty{}}

		for key, animal := range animals {

			fmt.Printf("%v,%T, %v %T ,%s\n",key,key,animal.WhoIam(), animal,str)
		}
	}

}