package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fieldsPtr := flag.String("f", "", "выбрать поля (колонки)")
	delimiterPtr := flag.String("d", "\t", "использовать другой разделитель")
	separatedPtr := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	Cut(*separatedPtr, *delimiterPtr, *fieldsPtr)
}

func Cut(separated bool, delimiter string, fields string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// если необходимо выбрать только строки с разделителем
		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		// разбиваем строку на колонки по разделителю
		columns := strings.Split(line, delimiter)

		// если необходимо выбрать только определенные колонки
		if fields != "" {
			fieldIndexes := strings.Split(fields, ",")
			var selectedColumns []string
			for _, indexString := range fieldIndexes {
				// преобразуем индекс колонки в число
				index := int(parseUint32(indexString) - 1)
				if index < len(columns) {
					selectedColumns = append(selectedColumns, columns[index])
				}
			}
			// выводим только выбранные колонки
			fmt.Println(strings.Join(selectedColumns, delimiter))
		} else {
			// выводим все колонки
			fmt.Println(line)
		}
	}
}

// parseUint32 - функция для преобразования строки в uint32.
// Возвращает 0, если преобразование не удалось.
func parseUint32(str string) uint32 {
	value, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(value)
}
