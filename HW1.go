package main

import (
	"fmt"
	"strings"
	"sort"
)

// Вспомогательные функции
func inputArray() []int {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&n)

	numbers := make([]int, n)
	fmt.Println("Введите элементы массива через пробел:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	return numbers
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func minMaxArray(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func mergeSortedArrays(array1, array2 []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(array1) && j < len(array2) {
		if array1[i] < array2[j] {
			result = append(result, array1[i])
			i++
		} else {
			result = append(result, array2[j])
			j++
		}
	}
	result = append(result, array1[i:]...)
	result = append(result, array2[j:]...)
	return result
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Основные задачи
func main() {
	tasks := []struct {
		number int
		desc   string
	}{
		{1, "Привет, мир!"},
		{2, "Сложение чисел"},
		{3, "Четное или нечетное"},
		{4, "Максимум из трех чисел"},
		{5, "Факториал числа"},
		{6, "Проверка символа"},
		{7, "Простые числа"},
		{8, "Строка и её перевёртыш"},
		{9, "Массив и его сумма"},
		{10, "Структуры и методы"},
		{11, "Конвертер температур"},
		{12, "Обратный отсчёт"},
		{13, "Длина строки"},
		{14, "Содержит ли массив?"},
		{15, "Среднее значение массива"},
		{16, "Таблица умножения"},
		{17, "Палиндром"},
		{18, "Найти минимум и максимум"},
		{19, "Удаление элемента из слайса"},
		{20, "Линейный поиск"},
		{21, "Удаление дубликатов"},
		{22, "Сортировка пузырьком"},
		{23, "Фибоначчиева последовательность"},
		{24, "Количество вхождений элемента в массив"},
		{25, "Пересечение двух массивов"},
		{26, "Анаграмма"},
		{27, "Слияние отсортированных массивов"},
		{28, "Хэш-таблица с коллизиями"},
		{29, "Бинарный поиск"},
		{30, "Очередь на основе двух стеков"},
	}

	// Вывод списка задач
	fmt.Println("Список доступных задач:")
	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.number, task.desc)
	}

	// Запрос номера задания
	var choice int
	fmt.Print("Введите номер задания: ")
	fmt.Scan(&choice)

	// Обработка выбора задания
	switch choice {
	case 1:
		fmt.Println("Привет, мир!")
	case 2:
		var a, b int
		fmt.Println("Введите два числа:")
		fmt.Scan(&a, &b)
		fmt.Printf("Сумма: %d\n", a+b)
	case 3:
		var a int
		fmt.Println("Введите число:")
		fmt.Scan(&a)
		if a%2 == 0 {
			fmt.Println("Четное")
		} else {
			fmt.Println("Нечетное")
		}
	case 4:
		var a, b, c int
		fmt.Println("Введите три числа:")
		fmt.Scan(&a, &b, &c)
		max := a
		if b > max {
			max = b
		}
		if c > max {
			max = c
		}
		fmt.Printf("Максимум: %d\n", max)
	case 5:
		var n int
		fmt.Println("Введите число:")
		fmt.Scan(&n)
		fact := 1
		for i := 2; i <= n; i++ {
			fact *= i
		}
		fmt.Printf("Факториал: %d\n", fact)
	case 6:
		fmt.Println("Введите символ:")
		var ch string
		fmt.Scan(&ch)
		vowels := "aeiouyаеёиоуыэюя"
		if strings.Contains(vowels, strings.ToLower(ch)) {
			fmt.Println("Гласная")
		} else {
			fmt.Println("Согласная")
		}
	case 7:
		var max int
		fmt.Println("Введите максимальное число:")
		fmt.Scan(&max)
		primes := make([]bool, max+1)
		for i := 2; i <= max; i++ {
			if !primes[i] {
				fmt.Printf("%d ", i)
				for j := i * i; j <= max; j += i {
					primes[j] = true
				}
			}
		}
		fmt.Println()
	case 8:
		fmt.Println("Введите строку:")
		var str string
		fmt.Scan(&str)
		fmt.Printf("Перевёрнутая строка: %s\n", reverseString(str))
	case 9:
		numbers := inputArray()
		fmt.Printf("Сумма элементов массива: %d\n", sumArray(numbers))
	case 10:
		type Rectangle struct {
			width, height float32
		}
		rectangle := Rectangle{}
		fmt.Println("Введите ширину и высоту прямоугольника:")
		fmt.Scan(&rectangle.width, &rectangle.height)
		area := rectangle.width * rectangle.height
		fmt.Printf("Площадь прямоугольника: %.2f\n", area)
	case 11:
		var celsius float64
		fmt.Println("Введите температуру в градусах Цельсия:")
		fmt.Scan(&celsius)
		fahrenheit := celsius*9/5 + 32
		fmt.Printf("Температура в градусах Фаренгейта: %.2f\n", fahrenheit)
	case 12:
		var n int
		fmt.Println("Введите число:")
		fmt.Scan(&n)
		for i := n; i > 0; i-- {
			fmt.Println(i)
		}
	case 13:
		fmt.Println("Введите строку:")
		var str string
		fmt.Scan(&str)
		length := 0
		for range str {
			length++
		}
		fmt.Printf("Длина строки: %d\n", length)
	case 14:
		numbers := inputArray()
		var target int
		fmt.Println("Введите число для поиска:")
		fmt.Scan(&target)
		found := false
		for _, v := range numbers {
			if v == target {
				found = true
				break
			}
		}
		if found {
			fmt.Println("Число найдено в массиве")
		} else {
			fmt.Println("Число не найдено в массиве")
		}
	case 15:
		numbers := inputArray()
		total := sumArray(numbers)
		average := float64(total) / float64(len(numbers))
		fmt.Printf("Среднее значение массива: %.2f\n", average)
	case 16:
		var n int
		fmt.Println("Введите число для таблицы умножения:")
		fmt.Scan(&n)
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d x %d = %d\n", n, i, n*i)
		}
	case 17:
		fmt.Println("Введите строку:")
		var str string
		fmt.Scan(&str)
		reversed := reverseString(str)
		if str == reversed {
			fmt.Println("Строка является палиндромом")
		} else {
			fmt.Println("Строка не является палиндромом")
		}
	case 18:
		numbers := inputArray()
		min, max := minMaxArray(numbers)
		fmt.Printf("Минимум: %d, Максимум: %d\n", min, max)
	case 19:
		fmt.Println("Введите количество элементов в массиве:")
		var n, index int
		fmt.Scan(&n)
		numbers := make([]int, n)
		fmt.Println("Введите элементы массива:")
		for i := 0; i < n; i++ {
			fmt.Scan(&numbers[i])
		}
		fmt.Println("Введите индекс для удаления:")
		fmt.Scan(&index)
		if index >= 0 && index < len(numbers) {
			numbers = append(numbers[:index], numbers[index+1:]...)
			fmt.Printf("Массив после удаления: %v\n", numbers)
		} else {
			fmt.Println("Недопустимый индекс")
		}
	case 20:
		fmt.Println("Введите массив:")
		array := inputArray()
		fmt.Println("Введите элемент для поиска:")
		var target int
		fmt.Scan(&target)
		index := binarySearch(array, target)
		if index != -1 {
			fmt.Printf("Элемент найден на позиции: %d\n", index)
		} else {
			fmt.Println("Элемент не найден")
		}
	case 21:
		numbers := inputArray()
		uniqueNumbers := removeDuplicates(numbers)
		fmt.Printf("Массив без дубликатов: %v\n", uniqueNumbers)
	case 22:
		numbers := inputArray()
		bubbleSort(numbers)
		fmt.Printf("Отсортированный массив: %v\n", numbers)
	case 23:
		var n int
		fmt.Println("Введите количество элементов в последовательности Фибоначчи:")
		fmt.Scan(&n)
		fib := make([]int, n)
		if n > 0 {
			fib[0] = 0
		}
		if n > 1 {
			fib[1] = 1
		}
		for i := 2; i < n; i++ {
			fib[i] = fib[i-1] + fib[i-2]
		}
		fmt.Printf("Последовательность Фибоначчи: %v\n", fib)
	case 24:
		var n int
		fmt.Println("Введите число N:")
		fmt.Scan(&n)
		a, b := 0, 1
		for a <= n {
			fmt.Print(a, " ")
			a, b = b, a+b
		}
		fmt.Println()
	case 25:
		fmt.Println("Введите массив:")
		array := inputArray()
		fmt.Println("Введите элемент для подсчета:")
		var target int
		fmt.Scan(&target)
		count := 0
		for _, v := range array {
			if v == target {
				count++
			}
		}
		fmt.Printf("Количество вхождений %d: %d\n", target, count)
	case 26:
		fmt.Println("Введите первую строку:")
		var str1 string
		fmt.Scan(&str1)
		fmt.Println("Введите вторую строку:")
		var str2 string
		fmt.Scan(&str2)
		sortedStr1 := strings.Split(strings.ToLower(str1), "")
		sortedStr2 := strings.Split(strings.ToLower(str2), "")
		sort.Strings(sortedStr1)
		sort.Strings(sortedStr2)
		if strings.Join(sortedStr1, "") == strings.Join(sortedStr2, "") {
			fmt.Println("Строки являются анаграммами")
		} else {
			fmt.Println("Строки не являются анаграммами")
		}
	case 27:
		fmt.Println("Введите первый отсортированный массив:")
		array1 := inputArray()
		fmt.Println("Введите второй отсортированный массив:")
		array2 := inputArray()
		merged := mergeSortedArrays(array1, array2)
		fmt.Printf("Объединённый массив: %v\n", merged)
	case 28:
		type HashTable struct {
			table map[string]string
		}

		hashTable := HashTable{table: make(map[string]string)}

		for {
			fmt.Println("1. Добавить пару (ключ-значение)")
			fmt.Println("2. Получить значение по ключу")
			fmt.Println("3. Выйти")

			var choice int
			fmt.Print("Выберите действие: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				var key, value string
				fmt.Print("Введите ключ: ")
				fmt.Scan(&key)
				fmt.Print("Введите значение: ")
				fmt.Scan(&value)
				hashTable.table[key] = value
				fmt.Println("Пара (ключ-значение) добавлена")
			case 2:
				var key string
				fmt.Print("Введите ключ: ")
				fmt.Scan(&key)
				if value, exists := hashTable.table[key]; exists {
					fmt.Printf("Значение для ключа '%s': %s\n", key, value)
				} else {
					fmt.Println("Ключ не найден")
				}
			case 3:
				return
			default:
				fmt.Println("Неверный выбор")
			}
		}
	case 29:
		fmt.Println("Введите отсортированный массив:")
		array := inputArray()
		fmt.Println("Введите элемент для поиска:")
		var target int
		fmt.Scan(&target)
		index := binarySearch(array, target)
		if index != -1 {
			fmt.Printf("Элемент найден на позиции: %d\n", index)
		} else {
			fmt.Println("Элемент не найден")
		}
	case 30:
		type Queue struct {
			stack1 []int
			stack2 []int
		}

		queue := Queue{}

		for {
			fmt.Println("1. Добавить элемент в очередь")
			fmt.Println("2. Удалить элемент из очереди")
			fmt.Println("3. Выйти")

			var choice int
			fmt.Print("Выберите действие: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				var element int
				fmt.Print("Введите элемент: ")
				fmt.Scan(&element)
				queue.stack1 = append(queue.stack1, element)
				fmt.Println("Элемент добавлен в очередь")
			case 2:
				if len(queue.stack2) == 0 {
					for len(queue.stack1) > 0 {
						queue.stack2 = append(queue.stack2, queue.stack1[len(queue.stack1)-1])
						queue.stack1 = queue.stack1[:len(queue.stack1)-1]
					}
				}
				if len(queue.stack2) == 0 {
					fmt.Println("Очередь пуста")
				} else {
					element := queue.stack2[len(queue.stack2)-1]
					queue.stack2 = queue.stack2[:len(queue.stack2)-1]
					fmt.Printf("Удалённый элемент: %d\n", element)
				}
			case 3:
				return
			default:
				fmt.Println("Неверный выбор")
			}
		}
	default:
		fmt.Println("Неверный номер задания")
	}
}
