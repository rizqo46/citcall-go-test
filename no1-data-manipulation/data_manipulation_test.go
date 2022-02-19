package no1datamanipulation_test

import (
	no1datamanipulation "citcall-go-test/no1-data-manipulation"
	"testing"
)

func TestLoadData(t *testing.T) {
	_, err := no1datamanipulation.LoadData()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestMakeTable(t *testing.T) {
	if err := no1datamanipulation.MakeTable(); err != nil {
		t.Error(err)
		t.Fail()
	}
}
