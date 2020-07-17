package main

import (
	"fmt"
	"sort"
	"strconv"
)

func SlicePlus1(array []int) {
	fmt.Println("1. Массив в начале", array)
	for i, numb := range array {
		array[i] = numb + 1
	}
	fmt.Println("1. Массив в конце", array)
}

func SliceInsertEnd(array []int) []int {
	fmt.Println("2. Массив в начале", array)
	array = append(array,5)
	fmt.Println("2. Массив в конце", array)
	return array
}

func SliceInsertBegin(array []int) []int {
	fmt.Println("3. Массив в начале", array)
	five := []int {5}
	array = append(five, array...)
	fmt.Println("3. Массив в конце", array)
	return array
}

func SliceDeleteLast(array []int) (res int, ar []int) {
	fmt.Println("4. Массив в начале", array)
	n := len(array)
	res = array[n-1]
	ar = array[:n-1]
	fmt.Println("4. Последний элемент", res)
	fmt.Println("4. Массив в конце", ar)
	return res, ar
}

func SliceDeleteFirst(array []int) (res int, ar []int) {
	fmt.Println("5. Массив в начале", array)
	res = array[0]
	ar = array[1:]
	fmt.Println("5. Первый элемент", res)
	fmt.Println("5. Массив в конце", ar)
	return res, ar
}

func SliceDelete(array []int, i int) (res int, ar []int) {
	fmt.Println("6. Массив в начале", array)
	n := len(array)
	if i >= n {
		panic("Взять i-ый элемент, превышающий размер массива, невозможно!" +
			" Введите число поменьше.")
	}
	res = array[i]
	ar = append(array[:i], array[i+1:]...)
	fmt.Println("6. i-ый элемент", res)
	fmt.Println("6. Массив в конце", ar)
	return res, ar
}

func SlicesTogether(array1 []int, array2 []int) []int {
	fmt.Println("7. Массивы в начале", array1, array2)
	array1 = append(array1, array2...)
	fmt.Println("7. Массив в конце", array1)
	return array1
}

func SlicesRemoveDifference(array1 []int, array2 []int) []int {
	fmt.Println("8. Массивы в начале", array1, array2)
	for _, numb2 := range array2 {
		for i, numb1 := range array1 {
			if numb1 == numb2 {
				array1 = append(array1[:i], array1[i+1:]...)
			}
		}
	}
	fmt.Println("8. Массив в конце", array1)
	return array1
}

func SliceShiftLeft1(array []int) {
	fmt.Println("9. Массив в начале", array)
	first := array[0]
	_ = copy(array[:(len(array) - 1)], array[1:])
	array[len(array)-1] = first
	fmt.Println("9. Массив в конце",array)
}

func SliceShiftLeft(array []int, k int) {
	fmt.Println("10. Массив в начале", array)
	n := len(array)
	if k > n {
		panic("Сдвиг на число, превышающее размер массива, невозможен!" +
			 	 " Введите число поменьше.")
	}
	rem := make([]int, k)
	_ = copy(rem, array[:k])
	_ = copy(array[:(n-k)], array[k:])
	_ = copy(array[(n-k):], rem)
	fmt.Println("10. Массив в конце", array)
}

func SliceShiftRight1(array []int) {
	fmt.Println("11. Массив в начале", array)
	n := len(array)
	last := array[n-1]
	_ = copy(array[1:], array[:(n-1)])
	array[0] = last
	fmt.Println("11. Массив в конце",array)
}

func SliceShiftRight(array []int, k int) {
	fmt.Println("12. Массив в начале", array)
	n := len(array)
	if k > n {
		panic("Сдвиг на число, превышающее размер массива, невозможен!" +
			" Введите число поменьше.")
	}
	rem := make([]int, k)
	_ = copy(rem, array[(n-k):])
	_ = copy(array[k:], array[:(n-k)])
	_ = copy(array[:k], rem)
	fmt.Println("12. Массив в конце", array)
}

func SliceCopy(array []int) []int {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)
	fmt.Println("13. Массив в конце", arrayCopy)
	return arrayCopy
}

func SliceSwap(array []int) {
	fmt.Println("14. Массив в начале", array)
	n := len(array)
	if n % 2 != 0 {
		n--
	}
	for i := 0; i < n; i += 2 {
		array[i], array[i+1] = array[i+1], array[i]
	}
	fmt.Println("14. Массив в конце", array)
}

func SliceSorts(array []int) []string{
	fmt.Println("15. Массив в начале", array)
	sort.Ints(array)
	fmt.Println("15. Прямая сортировка", array)
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	fmt.Println("15. Обратная сортировка", array)
	var arrayText []string
	for _, v := range array {
		numb := v
		st := strconv.Itoa(numb)
		arrayText = append(arrayText, st)
	}
	sort.Strings(arrayText)
	fmt.Println("15. Лексикографическая сортировка", arrayText)
	return arrayText
}

func main() {
	array := []int{0, 1, 2, 3, 4}
	arrayToo := []int{5, 1, 2, 7, 4}
	SlicePlus1(array)
	_ = SliceInsertEnd(array)
	_ = SliceInsertBegin(array)
	_, _ = SliceDeleteLast(array)
	_, _ = SliceDeleteFirst(array)
	_, _ = SliceDelete(array, 4)
	_ = SlicesTogether(array, arrayToo)
	_ = SlicesRemoveDifference(array, arrayToo)
	array = []int{0, 1, 2, 3, 4}
	SliceShiftLeft1(array)
	SliceShiftLeft(array, 3)
	SliceShiftRight1(array)
	SliceShiftRight(array, 3)
	_ = SliceCopy(array)
	SliceSwap(array)
	array = []int{0, 1, 5, 11, 15}
	SliceSorts(array)
}