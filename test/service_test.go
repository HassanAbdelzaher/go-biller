package sv_tst

import (
	"testing"
)

func Test(t *testing.T) {
	errF := t.Error
	logF := t.Log
	errff := t.Errorf
	logff := t.Logf
	TestService(logF, errF, errff, logff, "test_patter.json")
}
