package main	//определение пакета для текущего файла
import "fmt"	//подключение пакета fmt

//определение функции main
func main(){
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
}