package main

import "fmt"

func main() {
	//	fToC()
	ftToM()
}

func fToC() {
	var c, f float64
	fmt.Println("insert Fahrenheit value: ")
	fmt.Scanf("%f", &f)
	c = (f - 32) * 5 / 9
	fmt.Println(c)
}

func ftToM() {
	var ft, m float64
	fmt.Println("insert distance in feet")
	fmt.Scanf("%f", &ft)
	m = .3048 * ft
	fmt.Println("distance in m: ", m)
}
