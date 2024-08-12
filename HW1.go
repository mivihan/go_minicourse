package main

import (
	"fmt"
	"strings"
)

type Exercise struct {
	name   string
	action func()
}

var exercises = []Exercise{
	{
		name: "Привет, мир!",
		action: func() {
			fmt.Println("Привет, мир!")
		},
	},
	{
		name: "Сложение чисел",
		action: func() {
			var x, y int
			fmt.Scan(&x, &y)
			fmt.Printf("Сумма: %d\n", add(x, y))
		},
	},
	{
		name: "Четное или нечетное",
		action: func() {
			var num int
			fmt.Scan(&num)
			if isEven(num) {
				fmt.Println("Четное")
			} else {
				fmt.Println("Нечетное")
			}
		},
	},
	{
		name: "Максимум из трех чисел",
		action: func() {
			numbers := inputThreeNumbers()
			fmt.Printf("Максимальное: %d\n", max(numbers...))
		},
	},
	{
		name: "Факториал числа",
		action: func() {
			var n int
			fmt.Scan(&n)
			fmt.Printf("Факториал: %d\n", factorial(n))
		},
	},
	{
		name: "Проверка символа",
		action: func() {
			var char string
			fmt.Scan(&char)
			if isVowel(char) {
				fmt.Println("Гласная")
			} else {
				fmt.Println("Согласная")
			}
		},
	},
	{
		name: "Простые числа",
		action: func() {
			var max int
			fmt.Scan(&max)
			primes := sieveOfEratosthenes(max)
			fmt.Printf("Простые числа: %v\n", primes)
		},
	},
	{
		name: "Строка и ее перевертыш",
		action: func() {
			var input string
			fmt.Scan(&input)
			fmt.Printf("Перевернутая строка: %s\n", reverseString(input))
		},
	},
	{
		name: "Массив и его сумма",
		action: func() {
			numbers := inputArray()
			fmt.Printf("Сумма: %d\n", sumArray(numbers))
		},
	},
	{
		name: "Структуры и методы",
		action: func() {
			var width, height float32
			fmt.Scan(&width, &height)
			r := NewRectangle(width, height)
			fmt.Printf("Площадь: %.2f\n", r.Area())
		},
	},
	{
		name: "Конвертер температур",
		action: func() {
			fmt.Println("Конвертер в разработке...")
		},
	},
}

func add(a, b int) int {
	return a + b
}

func isEven(n int) bool {
	return n%2 == 0
}

func inputThreeNumbers() []int {
	var numbers [3]int
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}
	return numbers[:]
}

func max(nums ...int) int {
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func isVowel(char string) bool {
	vowels := "aeiouyаеёиоуыэюя"
	return strings.Contains(vowels, strings.ToLower(char))
}

func sieveOfEratosthenes(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func inputArray() []int {
	var size int
	fmt.Scan(&size)
	array := make([]int, size)
	for i := range array {
		fmt.Scan(&array[i])
	}
	return array
}

func sumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

type Rectangle struct {
	width  float32
	height float32
}

func NewRectangle(w, h float32) Rectangle {
	return Rectangle{width: w, height: h}
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

func main() {
	fmt.Println("Доступные задания:")

	for i, ex := range exercises {
		fmt.Printf("%d: %s\n", i+1, ex.name)
	}

	fmt.Print("Введите номер задания: ")
	var taskNum int
	fmt.Scan(&taskNum)
	taskNum--

	if taskNum < 0 || taskNum >= len(exercises) {
		fmt.Println("Некорректный номер задания")
		return
	}

	fmt.Printf("Выполнение задания: \"%s\"\n", exercises[taskNum].name)
	exercises[taskNum].action()
}
