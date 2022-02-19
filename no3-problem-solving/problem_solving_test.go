package no3problemsolving

import (
	"testing"
)

func TestProblemSolving(t *testing.T) {
	bebek := NewBebek(120, true, true, "petok")
	if bebek.energi != 120 ||
		!bebek.hidup ||
		!bebek.bisaTerbang ||
		bebek.suaraTerbang != "petok" {
		t.Log("gagal fungsi generate bebek")
		t.Fail()
	}
	if Makan(&bebek); bebek.energi != 121 {
		t.Log("gagal fungsi makan")
		t.Fail()
	}

	if Terbang(&bebek); bebek.energi != 120 {
		t.Log("gagal fungsi terbang")
		t.Fail()
	}
	if Mati(&bebek); bebek.hidup {
		t.Log("gagal fungsi mati")
		t.Fail()
	}
}
