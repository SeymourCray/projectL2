package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	a := `qwe\\5`

	fmt.Println(unpackString(a))
}

func unpackString(s string) (string, error) {
	var (
		res           strings.Builder              //итоговое значение собираем из strings.Builder
		count         []rune                       //тут сохраняем все цифры, на которые потом будем домножать символ
		lastCharacter rune                         //тут сохраняем последний символ в строке
		err           = errors.New("wrong format") //заранее создаем ошибку неверного формата
		escape        bool                         //экранирован ли следующий символ или нет
	)

	if len(s) == 0 { //если длина строки 0, возвращаем пустую строку
		return "", nil
	}

	if s[0] > 47 && s[0] < 58 || s[len(s)-1] == 92 { //если первый символ цифра или последний символ \ - ошибка
		return "", err
	}

	for _, ch := range s {
		if escape {
			if ch != 92 && (ch < 48 || ch > 57) { // если экранированный символ не \ или число - ошибка
				return "", err
			}
			lastCharacter = ch
			escape = false
		} else {
			switch true {
			case ch > 47 && ch < 58: //если элемент - число, то сохраняем в слайс
				count = append(count, ch)
			case ch == 92: //если элемент - \, ставим флаг, что следующий символ экранирован
				escape = true
				fallthrough
			default:
				n, ok := strconv.Atoi(string(count)) //считаем сколько нам надо повторить последний символ
				if ok != nil {                       //минимальное значение повтора - 1
					n = 1
				}
				count = nil             //очищаем слайс
				if lastCharacter != 0 { //если последний символ сохранен
					res.Grow(n)
					res.WriteString(strings.Repeat(string(lastCharacter), n)) //потовряем и записываем
				}
				lastCharacter = ch
			}
		}
	}

	n, ok := strconv.Atoi(string(count)) //повторяем операцию для последнего сохраненного символа
	if ok != nil {
		n = 1
	}
	res.Grow(n)
	res.WriteString(strings.Repeat(string(lastCharacter), n))

	return res.String(), nil
}
