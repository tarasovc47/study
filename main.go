package main

import (
	"fmt"
	"strconv"
	"strings"
)

const CONST_EXAMPLE_TEXT = "CONST_EXAMPLE1"
const CONST_EXAMPLE_INT = 12

func main() {
	//res0 := summator(4,6,7)
	//fmt.Println("res0 = " + strconv.FormatInt(res0, 10))

	//res1 := without_param()
	//fmt.Println("res1 = " + strconv.FormatInt(res1, 10))

	//example0_3()

	//slicer("1 2 3 4 5 6 7 8 9 0")
	bancomat()
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
func slicer(data string) {
	//fmt.Println("Тут просто печатаем входные данные как есть")
	//fmt.Println(data)
	//fmt.Println("Тут создаём пустой слайс из 10 значений")
	numbers := make([]int, 0, 10)
	for _, value := range data {
		if value != 32 {
			//fmt.Println("[внутри цикла] печатаем значение из DATA кроме 32")
			//fmt.Println(value)
			//fmt.Println("[внутри цикла] делаем i, _ := strconv.ParseInt(string(value), 10, 64) из DATA кроме 32")
			i, _ := strconv.ParseInt(string(value), 10, 64)
			// затем добавляем значение i в слайс
			numbers = append(numbers, int(i))
		}
	}
	fmt.Println("просто выводим весь массив (уже состоящий из чисел)")
	fmt.Println(numbers)
	even := make([]int, 0, 10)
	odd := make([]int, 0, 10)
	for nums := range numbers {
		if nums%2 == 0 {
			even = append(even, nums)
		} else {
			odd = append(odd, nums)
		}
	}
	fmt.Println("Чётные")
	fmt.Println(even)
	fmt.Println("Нечётные")
	fmt.Println(odd)
}

// ЗАДАНИЕ 2
// сгенерировано chatGPT, исправлены недостатки в формировании "тысяч" и правила для окончания доллара/доллр/долларов, оригинал ниже
/*package main

import (
"fmt"
"strings"
)

func main() {
	var amount int
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	if amount < 1 || amount > 9999 {
		fmt.Println("Некорректный ввод")
		return
	}

	ones := []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять", "десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	tens := []string{"", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	hundreds := []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
	thousands := []string{"", "тысяча", "тысячи", "тысяч"}

	numStr := fmt.Sprintf("%04d", amount)
	th := numStr[0] - '0'
	h := numStr[1] - '0'
	t := numStr[2] - '0'
	o := numStr[3] - '0'

	var sb strings.Builder

	if th > 0 {
		sb.WriteString(hundreds[th])
		sb.WriteByte(' ')
		if h == 0 && t == 0 && o == 0 {
			sb.WriteString(thousands[3])
		} else if h == 1 {
			sb.WriteString(thousands[1])
		} else if h >= 2 && h <= 4 {
			sb.WriteString(thousands[2])
		} else {
			sb.WriteString(thousands[3])
		}
		sb.WriteByte(' ')
	}

	if h > 0 {
		sb.WriteString(hundreds[h])
		sb.WriteByte(' ')
	}

	if t > 0 {
		if t == 1 {
			sb.WriteString(ones[t*10+o])
		} else {
			sb.WriteString(tens[t])
			if o > 0 {
				sb.WriteByte(' ')
				sb.WriteString(ones[o])
			}
		}
	} else if o > 0 {
		sb.WriteString(ones[o])
	}

	currency := "доллар"
	if amount != 1 {
		currency += "ов"
	}

	fmt.Println(sb.String() + " " + currency)
}*/

func bancomat() {
	var amount int
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	// вот здесь уже 090 превращается в 0 --->
	// fmt.Printf("%d\n", amount)
	if amount < 1 || amount > 9999 {
		fmt.Println("Некорректный ввод")
		return
	}

	ones := []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять", "десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	tens := []string{"", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	hundreds := []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
	thousands := []string{"", "одна тысяча", "две тысячи", "три тысячи", "четыре тысячи", "пять тысяч", "шесть тысяч", "семь тысяч", "восемь тысяч", "девять тысяч", "десять тысяч"}

	numStr := fmt.Sprintf("%04d", amount)
	th := numStr[0] - '0'
	h := numStr[1] - '0'
	t := numStr[2] - '0'
	o := numStr[3] - '0'

	var sb strings.Builder

	if th > 0 {
		sb.WriteString(thousands[th])
		sb.WriteByte(' ')
	}

	if h > 0 {
		sb.WriteString(hundreds[h])
		sb.WriteByte(' ')
	}

	if t > 0 {
		if t == 1 {
			sb.WriteString(ones[t*10+o])
		} else {
			sb.WriteString(tens[t])
			if o > 0 {
				sb.WriteByte(' ')
				sb.WriteString(ones[o])
			}
		}
	} else if o > 0 {
		sb.WriteString(ones[o])
	}

	currency := "доллар"
	if o >= 1 && o <= 9 && t == 1 {
		currency += "ов"
	} else if o == 2 ||
		o == 3 ||
		o == 4 {
		currency += "a"
	} else if o == 1 {
		currency += ""
	} else {
		currency += "ов"
	}

	fmt.Println(sb.String() + " " + currency)
}
