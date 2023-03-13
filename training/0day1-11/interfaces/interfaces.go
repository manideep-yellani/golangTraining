package main

import (
	"fmt"
	"math"
	"reflect"
)

type Abser interface {
	Abs() float64
	Abs1() float64
}

type Abser1 interface {
	Abs() float64
}

func main() {
	var a Abser
	fmt.Printf("%T %v\n", a, a)

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Printf("%T  %v \n", a, a)
	//a = &v // a *Vertex implements Abser
	fmt.Printf("%T  %v \n", a, a)

	//fmt.Printf("%T  %v \n",a.Abs1(),a.Abs1()) //this will give error

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	var g Vertex
	fmt.Println(g)
	g.X = 9
	fmt.Printf("%v \n", g)

	var hh MyFloat
	hh = 98
	vv := MyFloat(99)
	fmt.Println(vv.Abs(), v.Abs(), hh.Abs())

	fmt.Println(a.Abs())

	takesInterface(f)
	takesInterface(&v)

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs1() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func takesInterface(i Abser) {
	fmt.Printf("from fuction that takes interface %T %v %v\n", i, i, reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(i))
	j := Vertex{4, 5}
	takesInterface111(i)
	takesInterface111(&j)

	takesInterface222(i)
	takesInterface222(&j)

}

func takesInterface111(i Abser) {
	fmt.Printf("from fuction that takes interface111 %T %v\n", i, i)
}

func takesInterface222(i Abser1) {
	i.Abs()
	//i.Abs() gives error
	fmt.Printf("from fuction that takes interface222 %T %v\n", i, i)
}
