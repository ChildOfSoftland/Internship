package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"golang.org/x/tour/pic"
	"fmt"
	)

//Упражнение: циклы и функции
func Sqrt(x float64) float64 {
	z := 1.0
	var z0 float64
	for i := 0; i < 10; i++ {
		z0 = z
		z = z0 - (z0 * z0 - x)/(2 * z0)
	}
	return z
}

//Упражнение: срезы
func Pic(dx, dy int) [][]uint8 {
	grey := make([][]uint8, dy)
	for y:= 0; y < dy; y++ {
		grey[y] = make([]uint8, dx)
		for x:= 0; x < dx; x++ {
			grey[y][x] = uint8( x*x + y*y)
		}
	}
	return grey
}

//Упражнение: карты
func WordCount(s string) map[string]int {
	count := map[string]int{}
	for _, word := range strings.Fields(s) {
		count[word] = count[word] + 1
	}
	return count
}

//Упражнение: замыкание Фибоначчи
func fibonacci() func() int {
	first, second := -1, 1
	return func() int {
		numb := first + second
		first, second = second, numb
		return numb
	}
}

func main() {
	fmt.Println(Sqrt(2))
	pic.Show(Pic)
	wc.Test(WordCount)
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
