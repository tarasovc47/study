package main

import (
	"fmt"
	"strconv"
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
	//bancomat()
	//recursion()
	//comparsion([]int{1, 2, 3, 4}, []int{1, 2, 3, 4})
	//var someVar interface{} = 224
	//definer(someVar)
	//cutter()
	//concat()
	//creator()
	//checker()
	//initialSlice := []int{7, 77, 6, 1, 2, 3, 5, 8, 12, 13, 15, 140, 200, 201, 202, 201, 203, 204, 1110, 500, 502, 503, 504, 505}
	//average_value(initialSlice)
	//StrConvToInt("8н")
	//subseq(initialSlice)
	//csv_value := [][]string{
	//	{"Alice", "30", "alice@example.com"},
	//	{"Bob", "35", "bob@example.com"},
	//	{"Charlie", "40", "charlie@example.com"},
	//}
	//csv_writer(csv_value)
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

/*func bancomat() {
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
}*/
// ЗАДАНИЕ 3
/*func recursion() bool {
	var input int
	fmt.Print("Введите 0: ")
	fmt.Scan(&input)
	if input == 0 {
		fmt.Println("поздравляю, вы ввели \"ноль\"")
		return true
	} else if input < 1 || input > 9 {
		fmt.Println("попробуйте ещё раз...")
		return recursion()
	}
	return recursion()
}*/

// ЗАДАНИЕ 4
/*func comparsion(var1 []int, var2 []int) bool {
	//fmt.Println(var1)
	//fmt.Println(var2)
	for i := 0; i < len(var1); i++ {
		if var1[i] == var2[i] {
			//fmt.Println(var1[i], "равен", var2[i])
			continue
		}
		fmt.Println("Массивы не совпадают")
		return false
	}
	fmt.Println("Массивы совпадают")
	return true
}*/

// ЗАДАНИЕ 5
/*func definer(someVar interface{}) {
	//if reflect.TypeOf(someVar).Kind() == reflect.Int {
	//	fmt.Println("someVar is an integer")
	//} else if reflect.TypeOf(someVar).Kind() == reflect.Slice {
	//	fmt.Println("someVar is an slice")
	//} else {
	//	fmt.Println("someVar is not an integer and not a slice")
	//}
	_, ok := someVar.(int)
	if !ok {
		fmt.Println("someVar is not an integer")
	} else {
		fmt.Println("someVar is an integer")
	}
}*/

// ЗАДАНИЕ 6
/*func cutter() {
	myString := "Hello_World!_This.sentence.has_chars_I.want_to.remove!"
	replacer := strings.NewReplacer("_", " ", ".", " ") // ну это прям вообще дичь дичная
	myString = replacer.Replace(myString)
	fmt.Println(myString)
}*/

// ЗАДАНИЕ 7
// Слава ChatGPT! Гугл на тему конкатенации слайсов вообще околесицу выдаёт
/*func concat() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println(c)
}*/

// ЗАДАНИЕ 8
/*func StrConvToInt(str string) (int, error) {
	myInt, err := strconv.Atoi(str)
	if err != nil {
		//fmt.Println("Переменная str - не число!")
		fmt.Println(err)
		return 0, err
	}
	fmt.Println(myInt)
	return myInt, nil
}*/

// ЗАДАНИЕ 9
/*func creator() {
	file, err := os.Create("hello.txt") // создаем файл
	if err != nil {                     // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}
	defer file.Close()       // закрываем файл
	fmt.Println(file.Name()) // hello.txt
}*/
/*func checker() {
	fileInfo, err := os.Stat("hello.txt")
	if os.IsNotExist(err) {
		fmt.Println("Файла нет")
	} else {
		fmt.Println("Файл есть, его размер: ", fileInfo.Size())
	}
}*/

// ЗАДАНИЕ 10
/*func average_value(initialSlice []int) {
	sum := 0
	for _, value := range initialSlice {
		sum += value
	}
	aver := sum / len(initialSlice)
	fmt.Println(aver)
}*/

// ЗАДАНИЕ 11
/*func subseq(nums []int) {
	// подаётся это ----> initialSlice := []int{7, 77, 6, 1, 2, 3, 5, 8, 12, 13, 15, 140, 200, 201, 203, 204, 1110, 500, 502, 503, 504}
	summa := []int{}
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i+1]-nums[i] == 1 {
			if len(summa) > 0 {
				if summa[len(summa)-1] != nums[i] {
					summa = append(summa, nums[i])
					summa = append(summa, nums[i+1])
				}
			} else if len(summa) == 0 {
				summa = append(summa, nums[i])
				//summa = append(summa, nums[i+1])
			}
		}
	}
	fmt.Println(summa) // на выходе [1 2 3 12 13 200 201 203 204 502 503 504 505]
}*/

// ЗАДАНИЕ 12
/*func csv_writer(data [][]string) {
	file, err := os.Create("file.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Name", "Age", "Email"}
	writer.Write(header)

	for _, row := range data {
		writer.Write(row)
	}
}*/
