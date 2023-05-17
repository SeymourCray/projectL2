package main

import "testing"

func TestUnpackString(t *testing.T) {
	if res, _ := unpackString("a4bc2d5e"); res != "aaaabccddddde" {
		t.Error(`res, _ := unpackString("a4bc2d5e"); res != "aaaabccddddde"`)
	}

	if res, _ := unpackString("abcd"); res != "abcd" {
		t.Error(`res, _ := unpackString("abcd"); res != "abcd"`)
	}

	if _, err := unpackString("45"); err == nil {
		t.Error(`_, err := unpackString("45"); err == nil`)
	}

	if res, _ := unpackString(""); res != "" {
		t.Error(`res, _ := unpackString(""); res != ""`)
	}

	if res, _ := unpackString(`qwe\4\5`); res != `qwe45` {
		t.Error(`res, _ := unpackString("qwe\4\5"); res != "qwe45"`)
	}

	if res, _ := unpackString(`qwe\45`); res != `qwe44444` {
		t.Error(`res, _ := unpackString("qwe\45"); res != "qwe44444"`)
	}

	if res, _ := unpackString(`qwe\\5`); res != `qwe\\\\\` {
		t.Error(`res, _ := unpackString("qwe\\5"); res != "qwe\\\\\"`)
	}

	if _, err := unpackString(`aaa\`); err == nil {
		t.Error(`_, err := unpackString("aaa\"); err == nil`)
	}

	if _, err := unpackString(`aaa\q`); err == nil {
		t.Error(`_, err := unpackString("aaa\q"); err == nil`)
	}
}
