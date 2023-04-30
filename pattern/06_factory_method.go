package pattern

import "fmt"

type delivery interface {
	complete()
}

type shipDelivery struct {
}

func (s *shipDelivery) complete() {
	fmt.Println("make ship delivery")
}

type truckDelivery struct {
}

func (t *truckDelivery) complete() {
	fmt.Println("make truck delivery")
}

type iLogistic interface { //Abstract Interface
	makeDelivery() delivery
	completeDelivery(iLogistic)
}

type logistic struct { //Abstract Concrete Type
}

func (l *logistic) completeDelivery(i iLogistic) {
	d := i.makeDelivery()
	d.complete()
}

type seaLogistic struct {
	logistic
}

func (l *seaLogistic) makeDelivery() delivery {
	return &shipDelivery{}
}

type truckLogistic struct {
	logistic
}

func (l *truckLogistic) makeDelivery() delivery {
	return &truckDelivery{}
}

/*
Применимость

Фабричный метод изпользуется, когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код.
Он позволяет сэкономить системные ресурсы, повторно используя уже существующие объекты,
вместо повторного создания их каждый раз.

Плюсы и минусы

Нет тесной связи между классом создателя и конкретными классами объектов.
Принцип единственной ответственности. Вы можете переместить код создания продукта в одно место в программе,
что упростит поддержку кода.
Принцип открытости/закрытости. Вы можете вводить в программу новые типы продуктов, не нарушая существующий клиентский код.

Код может стать более сложным, поскольку нужно ввести много новых подклассов для реализации шаблона.
В идеале этот паттерн вводится в существующую иерархию классов-создателей.

В данном примере выделается два вида доставки: по морю и грузовиком. Каждый из этих способов
доставки "создается" в компаниях доставки и затем используется.
Abstract Interface и Abstract Concrete Type для создания абстрактного класса компании доставки
*/
