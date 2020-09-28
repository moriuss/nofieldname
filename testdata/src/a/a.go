package a

import "fmt"

type T1 struct {
	Pub string
	pri string
}

func okT1() T1 {
	return T1{Pub: "Pub", pri: "pri"}
}

func ng1T1() T1 {
	return T1{
		"Pub", // want "field name is missing"
		"pri", // want "field name is missing"
	}
}

func ng2T1() {
	t := T1{"Pub", "pri"} // want "field name is missing" "field name is missing"
	fmt.Print(t)
}

type T2 struct {
	One T1
	two int
}

func okT2() T2 {
	return T2{
		One: T1{Pub: "Pub", pri: "pri"},
		two: 2,
	}
}

func ngT2() T2 {
	return T2{
		One: T1{"Pub", "pri"}, // want "field name is missing" "field name is missing"
		two: 2,
	}
}

type T3 struct {
	T1
	three bool
}

func okT3() {
	t3 := T3{
		T1: T1{
			Pub: "Pub",
			pri: "pri",
		},
		three: true,
	}
	fmt.Println(t3)
}

func ng1T3() {
	t3 := T3{
		T1{ // want "field name is missing"
			Pub: "Pub",
			pri: "pri",
		},
		true, // want "field name is missing"
	}
	fmt.Println(t3)
}

func ng2T3() {
	t3 := T3{
		T1{"Pub", "pri"}, // want "field name is missing" "field name is missing" "field name is missing"
		true,             // want "field name is missing"
	}
	fmt.Println(t3)
}

type ss []string

func okSlice() {
	a := ss{
		"a",
		"b",
		"c",
	}
	fmt.Println(a)
}

func ok2Slice() {
	b := []int{1, 2, 3}
	fmt.Println(b)
}

func ok3Slice() {
	f := func(ss []string) {
		fmt.Println(ss)
	}
	f([]string{
		"a",
		"b",
		"c",
	})
}

var ok4Slice = [][]string{
	{"a", "b"},
	{"c", "d"},
}
