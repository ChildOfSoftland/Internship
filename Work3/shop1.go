package main

import (
	"fmt"
	"sort"
	"strings"
)

func AddProduct(prod map[string]float64, name string, price float64) {
	if _, ok := prod[name]; ok {
		fmt.Println("Товар с названием ", name, "уже существует. Повторное создание не возможно." +
			" Пожалуйста, используйте другое имя для создания нового товара.")
		return
	}
	prod[name] = price
}

func DelProduct(prod map[string]float64, name string) {
	if _, ok := prod[name]; !ok {
		fmt.Println("Товар с названием", name, "не существует. Удалить товар невозможно.")
		return
	}
	delete(prod, name)
}

func ChanNameProduct(prod map[string]float64, nameOld string, nameNew string) {
	if _, ok := prod[nameOld]; !ok {
		fmt.Println("Товар с названием", nameOld, "не существует. Изменить название товара невозможно.")
		return
	}
	prod[nameNew] = prod[nameOld]
	delete(prod, nameOld)
}

func ChanPriceProduct(prod map[string]float64, name string, priceNew float64) {
	if _, ok := prod[name]; !ok {
		fmt.Println("Товар с названием", name, "не существует. Изменить цену товара невозможно.")
		return
	}
	prod[name] = priceNew
}


type Order struct {
	Products []string
	//Bundles  []string
}

func PriceList() func(order Order, prod map[string]float64) (priceList float64) {
	cacheOrder := map[string]float64{}
	return func(order Order, prod map[string]float64) (priceList float64) {
		st := strings.Join(order.Products, " ")
		if sum, ok := cacheOrder[st]; ok {
			fmt.Println("From cache")
			return sum
		}
		for _, name := range order.Products {
			if _, ok := prod[name]; !ok {
				fmt.Println("Товар с названием", name, "не существует. Данный товар будет проигнорирован.")
				continue
			}
			priceList += prod[name]
		}
		cacheOrder[st] = priceList
		fmt.Println("Not from cache")
		return priceList
	}
}

func AccSort(acc map[string]float64) {
	keys := make([]string, 0, len(acc))
	for k := range acc {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Прямая сортировка по именам: ")
	for _, k := range keys {
		fmt.Println( k, ": ", acc[k])
	}
}

func AccReverseSort(acc map[string]float64) {
	keys := make([]string, 0, len(acc))
	for k := range acc {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	fmt.Println("Обратная сортировка по именам: ")
	for _, k := range keys {
		fmt.Println( k, ": ", acc[k])
	}
}

func AccCashSort(acc map[string]float64) {
	type key_value struct {
		Key   string
		Value float64
	}
	var sorted_struct []key_value
	for k, v := range acc {
		sorted_struct = append(sorted_struct, key_value{k, v})
	}
	sort.Slice(sorted_struct , func(i, j int) bool {
		return sorted_struct[i].Value > sorted_struct[j].Value
	})
	fmt.Println("Обратная сортировка по количеству денег: ")
	for _, key_value := range sorted_struct {
		fmt.Println( key_value.Key, key_value.Value)
	}
}

func main() {
	// Пункт 2
	product := map[string]float64{}
	AddProduct(product, "Молоко", 150.5)
	AddProduct(product, "Молоко", 155.5)
	AddProduct(product, "Хлеб", 55.2)
	AddProduct(product, "Яблоко", 210.0)
	AddProduct(product, "Порше", 20000000.0)
	fmt.Println(product)
	DelProduct(product, "Порше")
	DelProduct(product, "Ламборгини")
	fmt.Println(product)
	ChanNameProduct(product, "Хлеб", "Бородинский Хлеб")
	ChanNameProduct(product, "Молоко", "Коровье молоко")
	ChanNameProduct(product, "Машина", "Стиральная машина")
	fmt.Println(product)
	ChanPriceProduct(product, "Яблоко", 180.4)
	ChanPriceProduct(product, "Кот", 2000.9)
	fmt.Println(product)
	// Пункты 3 - 4
	list1 := Order{Products: []string{"Коровье молоко", "Коровье молоко", "Яблоко"}}
	list2 := Order{Products: []string{"Бородинский Хлеб", "Коровье молоко", "Яблоко"}}
	list3 := Order{Products: []string{"Коровье молоко", "Коровье молоко", "Яблоко"}}
	list4 := Order{Products: []string{"Хлеб", "Молоко", "Яблоко"}}
	summa := PriceList()
	fmt.Println(summa(list1, product))
	fmt.Println(summa(list2, product))
	fmt.Println(summa(list3, product))
	fmt.Println(summa(list4, product))
	// Пункты 5 - 6
	accounts := map[string]float64{"Василий":3000.5, "Генадий":2600.24, "Евгения":3400.7, "Мария":400.67}
	AccSort(accounts)
	AccReverseSort(accounts)
	AccCashSort(accounts)
}
