package my

import "testing"

func TestRndNumber(t *testing.T) {
	s := RndNumber(6)
	if !Test(s, `^\d{6}$`) {
		t.Log(s)
		t.Error("RndNumber failed.")
	}
}

func TestRndAlpha(t *testing.T) {
	s := RndAlpha(6)
	if !Test(s, `^[a-z]{6}$`) {
		t.Log(s)
		t.Error("RndAlpha failed.")
	}
}

func TestRndString(t *testing.T) {
	s := RndString(6)
	if !Test(s, `^[0-9a-z]{6}$`) {
		t.Log(s)
		t.Error("RndString failed.")
	}
}

func TestRndFilename(t *testing.T) {
	s := RndFilename("jpg")
	if !Test(s, `^[0-9]{20}\.jpg$`) {
		t.Log(s)
		t.Error("RndFilename failed.")
	}
}
