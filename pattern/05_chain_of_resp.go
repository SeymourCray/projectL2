package pattern

import "fmt"

type stage interface {
	execute(*student)
	setNext(stage)
}

type ophthalmologist struct {
	next stage
}

func (o *ophthalmologist) execute(s *student) {
	if s.ophthalmologistDone {
		fmt.Println("ophthalmologist has already examined student")
	} else {
		fmt.Println("ophthalmologist examined student")
		s.ophthalmologistDone = true
	}

	o.next.execute(s)
}

func (o *ophthalmologist) setNext(next stage) {
	o.next = next
}

type cardiologist struct {
	next stage
}

func (c *cardiologist) execute(s *student) {
	if s.cardiologistDone {
		fmt.Println("cardiologist has already examined student")
	} else {
		s.cardiologistDone = true
		fmt.Println("cardiologist examined student")
	}

	c.next.execute(s)
}

func (c *cardiologist) setNext(next stage) {
	c.next = next
}

type dentist struct {
	next stage
}

func (d *dentist) execute(s *student) {
	if s.dentistDone {
		fmt.Println("dentist has already examined student")
	} else {
		s.dentistDone = true
		fmt.Println("dentist examined student")
	}

	d.next.execute(s)
}

func (d *dentist) setNext(next stage) {
	d.next = next
}

type student struct {
	name                string
	ophthalmologistDone bool
	cardiologistDone    bool
	dentistDone         bool
}

/*
Применимость
Когда программа должна обрабатывать разнообразные запросы несколькими способами,
но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
Когда набор объектов, способных обработать запрос, должен задаваться динамически.

Плюсы и минусы

Уменьшает зависимость между клиентом и обработчиками.
Реализует принцип единственной ответственности.
Реализует принцип открытости/закрытости.

Запрос может остаться никем не обработанным.

В данном примере мы используем паттерн, чтобы студент поэтапно прошел обследование у разных врачей.
*/
