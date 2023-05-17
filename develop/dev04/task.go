package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:

'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	a := []string{"пятак", "пятак", "Пятка", "тяпка", "листок", "слиток", "столик", "стол", "лост", "море"}

	b := searchAnagrams(&a)

	for k, v := range *b {
		fmt.Println(k, *v)
	}
}

func searchAnagrams(a *[]string) *map[string]*[]string {
	var (
		res          = make(map[string]*[]string) //итоговое значение
		visited      = make(map[int]struct{})     //сохраняем слова которые посетили
		word1, word2 string                       //чтобы привести слова к нижнему регистру
	)

	for i := range *a {

		if _, ok := visited[i]; ok {
			continue
		}

		visited[i] = struct{}{}
		word1 = strings.ToLower((*a)[i])
		res[word1] = &[]string{word1} //создаем пару ключ значение с первым словом

		for j := i + 1; j < len(*a); j++ {

			if _, ok := visited[j]; !ok { //если еще не посещали слово, то проходим
				word2 = strings.ToLower((*a)[j])

				if compareWords(word1, word2) { //если слова анаграммы, то добавляем новое слово в массив (слайс)
					visited[j] = struct{}{}
					*res[word1] = append(*res[word1], word2)
				}
			}
		}

		*res[word1] = removeDuplicates(*res[word1]) //удаляем дубликаты

		if len(*res[word1]) == 1 { //если слов в множество только 1, то удаляем пару
			delete(res, word1)
			continue
		}

		sort.Strings(*res[word1]) //сортируем массив

	}

	return &res
}

func compareWords(s1, s2 string) bool { //сравниваем слова через мапы
	a := make(map[rune]int, 33)

	for _, i := range s1 {
		a[i]++
	}

	for _, i := range s2 {
		a[i]--
	}

	for _, i := range a {
		if i != 0 {
			return false
		}
	}

	return true
}

func removeDuplicates(s []string) []string { //удаляем дубликаты слов через мапы и пустые структуры
	length := len(s)
	set := make(map[string]struct{}, length)
	var res []string

	for _, i := range s {
		if _, ok := set[i]; !ok {
			set[i] = struct{}{}
			res = append(res, i)
		}
	}

	return res
}

/*
Проще было бы сделать так, чтобы значением в мапе было множество, тогда не надо будет удалять дубликаты.
Но я не совсем понял дано, обязательно ли чтобы был массив (ссылка на массив)
*/
