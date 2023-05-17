Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Программа выведет на экран числа 1-8 в произвольном порядке, так как
два канала a и b объединяются в канал c через функцию merge, которая
читает значения из обоих каналов и отправляет их в канал c. Затем значения из
канала c выводятся на экран в цикле for. Так как порядок чтения значений из каналов не определен,
то порядок вывода чисел на экран будет случайным.
```