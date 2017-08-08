package main

import "fmt"
import "./is308interfaces"
import "./course"
import "./comparable"

// Hello Passing an Interface from is308ifc as parameter
func Hello(i is308ifc.I) {
	fmt.Printf("Hi, my name is %s\n", i.M())
}

// TestCourse Passing an Interface from is308course as parameter
func TestCourse(c is308course.C) {

	err, num := c.F2("GetHours")
	fmt.Printf("Course name: %v, number: %f\n", err, num)
}

// CompareInteger32 just testing
func CompareInteger32(i comparable.Comparable, n comparable.Integer32) int8 {
	return i.CompareTo(n)
}

func main() {
	Hello(is308ifc.T{Name: "Michał"})
	TestCourse(is308course.T(1))

	//var n comparable.Integer32
	//n = 4
	var i comparable.Comparable
	//var t comparable.T{2}
	t := comparable.T{2}
	//t.Value = 2
	fmt.Printf("%+v\n", t)
	i = t
	fmt.Println(CompareInteger32(i, 2))

	//CompareInteger32(comparable.Comparable({Value:4}), 4)
}
