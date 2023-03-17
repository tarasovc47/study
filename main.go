package main

import (
	"fmt"
)

const CONST_EXAMPLE_TEXT = "CONST_EXAMPLE1"
const CONST_EXAMPLE_INT = 12

func main() {
	//res0 := summator(4,6,7)
	//fmt.Println("res0 = " + strconv.FormatInt(res0, 10))

	//res1 := without_param()
	//fmt.Println("res1 = " + strconv.FormatInt(res1, 10))

	//example0_3()

	slicer("1 2 3 5 6 7 8 0")
}

// ЗАДАНИЕ 0
func summator(num1 int64, num2 int64, num3 int64) int64 {
	sum := num1 + num2 + num3
	return sum
}

func without_param() int64 {
	sum2 := int64(CONST_EXAMPLE_INT + 4)
	return sum2
}

func example0_3() {
	ex1 := "Я "
	ex2 := "ничего "
	ex3 := "не "
	ex4 := "принимающая "
	ex5 := "и "
	ex6 := "ничего "
	ex7 := "отдающая "
	ex8 := "функция "
	fmt.Println(ex1 + ex2 + ex3 + ex4 + ex5 + ex6 + ex3 + ex7 + ex8)
}

// ЗАДАНИЕ 1
func slicer(row string) {
	//сначала нарезаем мтроку на массив
	arr := row[1]
	fmt.Println(arr)
}
