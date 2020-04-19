package tinkoff_test

import (
	"os"
	"testing"

	"github.com/nikita-vanyasin/tinkoff"
)

func helperCreateClient(tb testing.TB) *tinkoff.Client {
	tb.Helper()

	tk := os.Getenv("TERMINAL_KEY")
	tp := os.Getenv("TERMINAL_PASSWORD")
	if tk == "" || tp == "" {
		tb.Fatal("provide TERMINAL_KEY and TERMINAL_PASSWORD env vars to execute tests")
	}
	return tinkoff.NewClient(tk, tp)
}

func assertNotError(tb testing.TB, err error) {
	tb.Helper()

	if err != nil {
		tb.Fatalf("expected no error, got %s", err.Error())
	}
}

func assertEq(tb testing.TB, expected, actual interface{}) {
	tb.Helper()

	if expected != actual {
		tb.Fatalf("expected %v, got %v", expected, actual)
	}
}

func assertNotEmptyString(tb testing.TB, actual string) {
	tb.Helper()

	if actual == "" {
		tb.Fatalf("expected empty string, got %s", actual)
	}
}
