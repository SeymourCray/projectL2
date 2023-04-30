package pattern

import "fmt"

type activationFunc interface {
	activate(*tensor)
}

type reLU struct {
}

func (r reLU) activate(t *tensor) {
	fmt.Println("use ReLu")
}

type preLU struct {
}

func (r preLU) activate(t *tensor) {
	fmt.Println("use PReLu")
}

type rreLU struct {
}

func (r rreLU) activate(t *tensor) {
	fmt.Println("use RReLu")
}

type tensor struct {
	values         []float64
	activationFunc activationFunc
}

func (t *tensor) setActivationFunc(a activationFunc) {
	t.activationFunc = a
}

func (t *tensor) activate() {
	t.activationFunc.activate(t)
}

/*
Применимость

Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта
Когда есть множество похожих классов, отличающихся только некоторым поведением
Когда надо закрыть детали реализации алгоритмов для других классов

Плюсы и минусы

Горячая замена алгоритмов на лету
Изолирует код и данные алгоритмов от остальных классов
Уход от наследования к делегированию
Реализует принцип открытости/закрытости

Усложняет программу за счёт дополнительных классов.
Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую

В данном примере мы использовали паттерн стратегии чтобы описать применение
различных функций активации для входных значений нейронной сети
*/
