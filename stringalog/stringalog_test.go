package stringalog

import (
	"testing"
)

func TestVariantWord(t *testing.T) {
	str10 := "123"
	str20 := "321"
	if !VariantWord(str10, str20) {
		t.Errorf("VaroamtWprd exect true. str %s %s", str10, str20)
	}

	str11 := "1234"
	str21 := "1235"

	if VariantWord(str11, str21) {
		t.Errorf("VaroamtWprd exect true. str %s %s", str11, str21)
	}

	str12 := "1234"
	str22 := "123115"

	if VariantWord(str12, str22) {
		t.Errorf("VaroamtWprd exect true. str %s %s", str12, str22)
	}
}

func TestSubNumberSum(t *testing.T) {
	exampleArray := []struct {
		Val string
		Res int
	}{
		{"A1CD2E33", 36}, {"A-1B--2C-D6E", 7},
	}

	for _, v := range exampleArray {
		r := SubNumberSum(v.Val)
		if r != v.Res {
			t.Errorf(" subnumbersum fail, example %s expect %d but return %d.", v.Val, v.Res, r)
		}
	}

}

func TestKmp(t *testing.T) {
	exampleArray := []struct {
		Val string
		Sub string
		Res int
	}{
		{"123456789", "789", 6}, {"acbc", "bc", 2}, {"acbc", "bcc", -1},
	}

	for i, _ := range exampleArray {
		r := Kmp(exampleArray[i].Val, exampleArray[i].Sub)
		if r != exampleArray[i].Res {
			t.Errorf("Kmp  src %s sub %s expect %d bug return %d.", exampleArray[i].Val, exampleArray[i].Sub, exampleArray[i].Res, r)
		}
	}
}

func TestCharUniq(t *testing.T) {
	exampleArray := []struct {
		Val string
		Res bool
	}{
		{"1234567", true}, {"111123454", false},
	}

	for i, _ := range exampleArray {
		r := CharUniq(exampleArray[i].Val)
		if r != exampleArray[i].Res {
			t.Errorf("charunip val %s expect %v but %v.", exampleArray[i].Val, exampleArray[i].Res, r)
		}
	}
}
