package pattern

import "fmt"

type state interface {
	freeze()
	boil()
	transfuse()
}

type water struct {
	ice          state
	liquid       state
	steam        state
	currentState state
}

func (w *water) setState(s state) {
	w.currentState = s
}

func (w *water) freeze() {
	w.currentState.freeze()
}

func (w *water) boil() {
	w.currentState.boil()
}

func (w *water) transfuse() {
	w.currentState.transfuse()
}

type ice struct {
	w *water
}

func (i *ice) freeze() {
	fmt.Println("ice cannot be frozen!")
}

func (i *ice) boil() {
	fmt.Println("got water!")
	i.w.setState(i.w.liquid)
}

func (i *ice) transfuse() {
	fmt.Println("ice cannot be poured!")
}

type liquid struct {
	w *water
}

func (l *liquid) freeze() {
	l.w.setState(l.w.ice)
	fmt.Println("got ice!")
}

func (l *liquid) boil() {
	l.w.setState(l.w.steam)
	fmt.Println("got steam!")
}

func (l *liquid) transfuse() {
	fmt.Println("transfused water into a cup")
}

type steam struct {
	w *water
}

func (s steam) freeze() {
	fmt.Println("got water!")
}

func (s steam) boil() {
	fmt.Println("steam cannot be boiled!")
}

func (s steam) transfuse() {
	fmt.Println("steam cannot be poured!")
}

/*
Применимость

Когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
причём типов состояний много, и их код часто меняется
Когда код класса содержит множество больших, похожих друг на друга, условных операторов,
которые выбирают поведения в зависимости от текущих значений полей класса

Плюсы и минусы

Избавляет от множества больших условных операторов машины состояний
Концентрирует в одном месте код, связанный с определённым состоянием
Упрощает код контекста

Может неоправданно усложнить код, если состояний мало и они редко меняются

В данном примере мы используем паттерн состояния, чтобы показать как выполняется различные действия с водой
в ее разных агрегатных состояниях
*/
