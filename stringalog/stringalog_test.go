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
