package main

import (
	"fmt"
	"strings"
)

func MapWordCount(st string) map[string]int {
	fmt.Println("1. Строка", st)
	count := map[string]int{}
	for _, word := range strings.Fields(st) {
		count[word] = count[word] + 1
	}
	fmt.Println("1. Словарь с количеством слов в строке", count)
	return count
}

type EmptyStr struct {}

func MapFindNumb(array []int) map[int]EmptyStr {
	fmt.Println("2. Массив чисел", array)
	find := map[int]EmptyStr{}
	for _, numb := range array {
		find[numb] = EmptyStr{}
	}
	fmt.Println("2. Словарь с найденными числами", find)
	return find
}

func MapFindNumbTwo(array1 []int, array2 []int) map[int]EmptyStr {
	fmt.Println("3. Массивы чисел", array1, "\n", array2)
	find := map[int]EmptyStr{}
	for _, numb1 := range array1 {
		for _, numb2 := range array2{
			if numb1 == numb2 {
				find[numb1] = EmptyStr{}
			}
		}
	}
	fmt.Println("3. Словарь с найденными числами", find)
	return find
}

func MapFibonach() func(n int) int {
	fib := map[int]int{0:0, 1:1}
	return func(n int) int {
		fmt.Println("4. Переданный номер числа Фибоначчи", n)
		if len(fib) < n {
			for i := len(fib); i < n; i++{
				fib[i] = fib[i-1] + fib[i-2]
			}
		}
		fmt.Println("4. Число Фибонначи, под номером ", n, "равно ", fib[n-1])
		fmt.Println("4. Словарь всех кэшированных чисел Фибоначчи ", fib)
		return fib[n-1]
	}
}

func main() {
	_ = MapWordCount("I ate a donut. Then I ate another donut.")
	array := []int{0, 1, 1, 3, 7, 5, 6, 7, 4, 9, 8, 21, 12, 13, 17, 15, 14, 17, 18, 24, 20, 24, 22, 23, 24, 25}
	arrayToo := []int{0, 3, 5, 33, 70, 4, 46, 12, 12, 43, 15, 16, 17, 19, 34, 20, 63, 52, 15, 20, 45}
	_ = MapFindNumb(array)
	_ = MapFindNumbTwo(array, arrayToo)
	f := MapFibonach()
	_ = f(50)
	_ = f(21)
	_ = f(13)
	_ = f(45)
}
