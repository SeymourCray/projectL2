package pattern

import "fmt"

type Builder interface {
	setCake()
	setTomato()
	setCheese()
	setPineapple()
}

type pizza struct {
	cake      bool
	tomato    bool
	cheese    bool
	pineapple bool
}

type pizzaBuilder struct {
	cake      bool
	tomato    bool
	cheese    bool
	pineapple bool
}

func (p *pizzaBuilder) setCake() {
	p.cake = true
	fmt.Println("added cake...")
}

func (p *pizzaBuilder) setTomato() {
	p.tomato = true
	fmt.Println("added tomato...")
}

func (p *pizzaBuilder) setCheese() {
	p.cheese = true
	fmt.Println("added cheese...")
}

func (p *pizzaBuilder) setPineapple() {
	p.pineapple = true
	fmt.Println("added pineapple...")
}

func getPizza(p *pizzaBuilder) pizza {
	return pizza{
		cake:      p.cake,
		tomato:    p.tomato,
		cheese:    p.cheese,
		pineapple: p.pineapple,
	}
}

/*
Строитель. Применимость

Когда надо избавиться от «телескопического конструктора».
Когда код должен создавать разные представления какого-то объекта. Например, пицца с разными видами теста
Когда нужно собирать сложные составные объекты.

Плюсы и минусы

Позволяет создавать объекты пошагово.
Позволяет использовать один и тот же код для создания различных объектов.
Изолирует сложный код сборки объекта от его основной бизнес-логики.

Усложняет код программы из-за введения дополнительных классов.
Клиент может быть привязан к конкретным классам строителей.

В данном примере с помощью строителя мы собираем свою пиццу.
*/
