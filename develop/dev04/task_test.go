package main

import (
	"reflect"
	"testing"
)

func TestSearchAnagrams(t *testing.T) {
	trueMap := &map[string]*[]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

	res := searchAnagrams(&[]string{"Пятак", "пятак", "пятка", "тяпка", "листок", "слиток", "столик", "море"})

	if !reflect.DeepEqual(res, trueMap) {
		t.Error(`res != trueMap`)
	}
}
