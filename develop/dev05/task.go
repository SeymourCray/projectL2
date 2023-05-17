package main

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Код должен проходить проверки go vet и golint.
*/

func main() {

	var (
		found               bool
		err                 error
		count, stringsAfter int          //N, для подсчета строк и подсчета строк после совпадения
		queue               = list.New() //храним прошлые строки в очереди
	)

	ignoreCase := flag.Bool("i", false, "")
	printAfter := flag.Int("A", 0, "")
	printLineNumber := flag.Bool("n", false, "")
	printBefore := flag.Int("B", 0, "")
	printContext := flag.Int("C", 0, "")
	printCount := flag.Bool("c", false, "")
	fixedString := flag.Bool("F", false, "")
	invertOutput := flag.Bool("v", false, "")
	flag.Parse()

	var file, pattern string
	if flag.NArg() < 2 {
		log.Fatalln("not enough values")
	} else {
		pattern, file = flag.Args()[0], flag.Args()[1]
	}

	var re *regexp.Regexp //подготавливаем выражение
	if !*fixedString {
		if *ignoreCase {
			re, err = regexp.Compile("(?i)" + pattern)
		} else {
			re, err = regexp.Compile(pattern)
		}
	}

	if err != nil {
		log.Fatalln("Invalid pattern:", err.Error())
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatalln("Error opening file:", err.Error())
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if (*printBefore != 0 || *printContext != 0) && stringsAfter < 1 { //если есть ключи printBefore || printContext и мы не выводим строки после совпадения
			if *printBefore+*printContext+1 > queue.Len() { //ограничиваем очередь длиной в num+1
				queue.PushBack(line)
			} else {
				queue.Remove(queue.Front())
				queue.PushBack(line)
			}
		}

		if *fixedString { //сравниваем либо по фиксированной строке, либо по паттерну
			if *ignoreCase {
				found = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			} else {
				found = strings.Contains(line, pattern)
			}
		} else {
			found = re.MatchString(line)
		}

		if found != *invertOutput { //если совпало, инверсия учитывается
			count++
			if *printCount {
				continue
			}

			if *printBefore != 0 || *printContext != 0 { //выводим все строки до
				l := queue.Len()
				for i := 0; i < l-1; i++ {
					v := queue.Front()
					printLine(*printLineNumber, lineNum-l+i+1, v.Value.(string), "-")
					queue.Remove(v)
				}
				queue.Remove(queue.Front())
			}

			printLine(*printLineNumber, lineNum, line, ":") //выводим текущую строку

			if *printAfter != 0 || *printContext != 0 { //делаем счетчик для вывода строк после
				stringsAfter = *printAfter + *printContext + 1
			}

		} else if stringsAfter > 0 { //печатаем строки после
			printLine(*printLineNumber, lineNum, line, "-")
		}

		stringsAfter--
	}

	if *printCount { //если есть ключ для кол-ва строк - выводим
		fmt.Println(count)
	}
}

func printLine(f bool, n int, s, sep string) { //функция для вывода строки с учетом номера
	if f {
		fmt.Printf("%d%s%v\n", n, sep, s)
	} else {
		fmt.Println(s)
	}
}
