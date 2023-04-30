package pattern

import "fmt"

type waiter struct { //invoker
	command command
}

func (w *waiter) makeOrder() {
	w.command.execute()
}

type command interface {
	execute()
}

type orderFish struct { //command
	cook cook
}

func (c *orderFish) execute() {
	c.cook.cookFish()
}

type orderMeat struct { //command
	cook cook
}

func (c *orderMeat) execute() {
	c.cook.cookMeat()
}

type cook interface {
	cookFish()
	cookMeat()
}

type chief struct { //receiver
}

func (c *chief) cookFish() {
	fmt.Println("fish was cooked!")
}

func (c *chief) cookMeat() {
	fmt.Println("meat was cooked!")
}

/*
Применимость

Позволяет инкапсулировать действия в объекты.
Основная идея, стоящая за шаблоном — это предоставление средств,
для разделения клиента и получателя.
Шаблон команда может быть использован для реализации системы, основанной на транзакциях,
где сохраняется история команд сразу после выполнения.
Если окончательная команда успешно выполнена, то все хорошо,
иначе алгоритм просто перебирает историю и продолжает выполнять отмену для всех выполненных команд.

Плюсы и минусы

Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
Позволяет реализовать простую отмену и повтор операций.
Расширения для добавления новой команды просты и могут быть выполнены без изменения существующего кода.
Можно определить систему отката с помощью паттерна Command.

Создается большое количество классов и объектов, работающих вместе.
Каждая отдельная команда является ConcreteCommand.

В данном примере с помощью паттерна мы реализуем систему заказа (command) в ресторане
через официанта (invoker), который передает его
шеф-повару (receiver)
*/
