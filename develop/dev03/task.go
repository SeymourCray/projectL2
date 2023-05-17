package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Код должен проходить проверки go vet и golint.
*/

func main() {
	// Определение флагов командной строки
	reversePtr := flag.Bool("r", false, "Reverse the result")
	numericPtr := flag.Bool("n", false, "Sort numerically")
	uniquePtr := flag.Bool("u", false, "Output only unique lines")
	columnNumber := flag.Int("k", 0, "Column to sort")
	flag.Parse()
	// Чтение входных данных из файла
	var input []string
	var file string

	if flag.NArg() != 1 {
		log.Fatalln("wrong format")
	} else {
		file = flag.Args()[0]
		f, err := os.Open(file)
		if err != nil {
			log.Fatalln(err.Error())
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			input = append(input, scanner.Text())
		}
	}
	// Сортировка входных данных
	if *numericPtr {
		sort.Slice(input, func(i, j int) bool {
			return parseNum(input[i]) < parseNum(input[j])
		})
	} else if *columnNumber != 0 {
		sort.Slice(input, func(i, j int) bool {
			return parseStr(input[i], *columnNumber-1) < parseStr(input[j], *columnNumber-1)
		})
	} else {
		sort.Strings(input)
	}
	// Удаление дубликатов, если указан флаг -u
	if *uniquePtr {
		input = removeDuplicates(input)
	}
	// Обратный порядок сортировки, если указан флаг -r
	if *reversePtr {
		reverse(input)
	}
	// nВывод отсортированных данных в новый файл
	f, err := os.Create("new_" + file)
	if err != nil {
		log.Fatalln(err.Error())
	}

	w := bufio.NewWriter(f)

	for _, s := range input {
		w.WriteString(s)
		w.WriteString("\n")
	}

	w.Flush()
}

// Функция для парсинга чисел из строк
func parseNum(s string) float64 {
	var num float64
	fmt.Sscanf(s, "%f", &num)
	return num
}

// Функция для извлечения слова по номеру
func parseStr(s string, p int) string {
	return strings.Split(s, " ")[p]
}

// Функция для удаления дубликатов из отсортированного списка строк
func removeDuplicates(input []string) []string {
	if len(input) == 0 {
		return input
	}
	j := 0
	for i := 1; i < len(input); i++ {
		if input[j] == input[i] {
			continue
		}
		j++
		input[j] = input[i]
	}
	return input[:j+1]
}

// Функция для обратной сортировки списка строк
func reverse(input []string) {
	for i := 0; i < len(input)/2; i++ {
		j := len(input) - i - 1
		input[i], input[j] = input[j], input[i]
	}
}
