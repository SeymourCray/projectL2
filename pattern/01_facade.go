package pattern

import "fmt"

type carFactoryFacade struct {
	engineFactory  *engineFactory
	bodyFactory    *bodyFactory
	chassisFactory *chassisFactory
}

func newCarFactoryFacade() *carFactoryFacade {
	return &carFactoryFacade{
		newEngineFactory(),
		newBodyFactory(),
		newChassisFactory()}
}

func (factory *carFactoryFacade) makeCar() {
	fmt.Println("start assemble new car...")

	factory.bodyFactory.makeBody()
	factory.chassisFactory.makeChassis()
	factory.engineFactory.makeEngine()

	fmt.Println("car is ready!")
}

type engineFactory struct {
}

func newEngineFactory() *engineFactory {
	return &engineFactory{}
}

func (factory *engineFactory) makeEngine() {
	fmt.Println("Making engine...")
}

type bodyFactory struct {
}

func newBodyFactory() *bodyFactory {
	return &bodyFactory{}
}

func (factory *bodyFactory) makeBody() {
	fmt.Println("Making body...")
}

type chassisFactory struct {
}

func newChassisFactory() *chassisFactory {
	return &chassisFactory{}
}

func (factory *chassisFactory) makeChassis() {
	fmt.Println("Making chassis...")
}

/*
Фасад. Применимость.

Когда нужно представить простой или урезанный интерфейс к сложной подсистеме.
Когда надо уменьшить количество зависимостей между клиентом и сложной системой.
Фасадные объекты позволяют отделить, изолировать компоненты системы от клиента и развивать и работать с ними независимо.
Когда надо разложить подсистему на отдельные слои.

Плюс: изолирует клиентов от компонентов сложной подсистемы.
Минус: фасад рискует стать божественным объектом, привязанным ко всем классам программы.

В данном примере фасадом является цех по производству машин, который скрывает в себе более мелкие цеха по производству
составных частей автомобиля, таких как двигатель, шасса и кузов.
*/
