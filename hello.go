package main		//определение пакета для текущего файла
import (
	"fmt"		//подключение пакета fmt
	"strconv"
)

//определение функции main
func main(){
	// variables()
	// showConditions()
	// shiftBits()
	// bitOperations()
	// arrayExamples()
	// switchExample()
	// cycleExamples()
	// sumVariables(1, 2, 3, 4.5, 6.7)
	// sumIntNumbers(1, 2, 3, 4, 5)
	
	// fmt.Println(returnFuncResult(7, 8))
	// fmt.Println(returnFuncResultNamedVariable(9, 8))
	fmt.Println(returnSeveralVariables(9, 8, "pen", "apple"))
	fmt.Println(returnSeveralVariablesWithNamedVariables(9, 8, "pen", "apple"))
}

func variables(){
	fmt.Println("Hello Go!");fmt.Println("Hello Golang!")
	fmt.Println("Hello Go!")
	fmt.Println("123456")

	var value1 string = "переменная 1"
	fmt.Println(value1);
	var value2, value3 string
	value2 = "переменная 2"
	value3 = "переменная 3"
	fmt.Println(value2)
	fmt.Println(value3)

	value4 := 123
	fmt.Println(value4)
	value5 := "dadadas"
	fmt.Println(value5)

	var value6, value7 string = "значение 6", "Значение 7"
	fmt.Println(value6 + value7)
	const (
		pi float32 = 3.1415
		b
		e float32 = 2.72
		c
	)
	fmt.Printf("pi %f", pi)
	fmt.Println(b)
	var a1 = 4
	var b1 = 6
	var c1 = a1 + b1
	fmt.Println(c1)
	var c2 = a1 - b1
	fmt.Println(c2)
	var c3 = a1 * b1
	fmt.Println(c3)
	var c4 = b1 / a1
	fmt.Println(c4)
	var c5 = b1 % a1
	fmt.Println(c5)
}

func showConditions(){
	var a int = 10
	var b int = 10
	var c bool = a == b
	fmt.Println(c)
	if a < b {
		fmt.Println("a < b")
	}else if a == b{
		fmt.Println("b = a")
	} else{
		fmt.Println("b > a")
	}
}

func switchExample(){
	var a int = 6
	switch(a){
		case 9:
			fmt.Println("a = 9")
		case 8:
			fmt.Println("a = 8")
		case 7:
			fmt.Println("a = 7")
		case 6, 5, 4:
			fmt.Println("a = 6, 5, 4")
		default:
			fmt.Println("other examples")
	}
}

func shiftBits(){
	var a int = 2 << 2;
	fmt.Println(a)
	var c int = 16 >> 3;
	fmt.Println(c);
}

func bitOperations(){
	var a int = 5 | 2;	//101 | 010 = 111 (7)
	fmt.Println(a)
	var b int = 5 & 2;	//101 & 010 = 000 (0)
	fmt.Println(b)
	var c int = 5 ^ 2;	//101 ^ 010 = 111 (7)
	fmt.Println(c)
	var d int = 5 &^ 2;	//101 &^ 010 = 101 (5)
	fmt.Println(d)
}
func arrayExamples(){
	var numbers [10]int = [10]int{1,3,4,2,5,6,7,9,8}
	fmt.Println(numbers)
	numbers[5] = 12
	fmt.Println(numbers)
	fmt.Println(numbers[7])
	var colors = [3]string{2: "blue", 0:"red", 1:"green"}
	fmt.Println(colors)
	fmt.Println(colors[2])
}

func cycleExamples(){
	for i := 0; i < 10; i++{
		var str string = strconv.Itoa(i) + " "
		fmt.Print(str)
	}
	fmt.Println()

	var counter int = 0
	for ; counter < 10; counter++{
		fmt.Println(counter * counter)
	}

	var c int = 0
	for c < 10{
		fmt.Println(c * c)
		c++
	}

	for i := 1; i < 10; i++{
		for j := 1; j < 10; j++{
			fmt.Print(i * j, "\t")
		}
		fmt.Println()
	}

	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users{
		fmt.Println(index, value)
	}

	for _, value := range users{
		fmt.Println(value)
	}

	for i:=0; i < len(users); i++{
		fmt.Println(users[i])
	}
	
	for j:=0; j < 100; j++{
		if j > 10 {
			fmt.Println("breaked")
			break
		}else{
			fmt.Println(j)
		}
	}
}

func sumVariables(a, b, c int, x, y float32){
	fmt.Println("IntegersSum: ", a + b + c)
	fmt.Println("Float32Sum: ", x + y)
}

func sumIntNumbers(numbers ...int){
	var sum int = 0
	for _, number := range numbers{
		sum += number
	}
	fmt.Println("sum: ", sum)
}

func returnFuncResult(x, y int) int{
	return x + y
}

func returnFuncResultNamedVariable(x, y int) (z int){
	z = x + y
	return
}

func returnSeveralVariables(x, y int, str1, str2 string) (int, string){
	sum := x + y
	str := str1 + str2
	return sum, str
}

func returnSeveralVariablesWithNamedVariables(x, y int, str1, str2 string) (sum int, str string){
	sum = x + y
	str = str1 + str2
	return
}

