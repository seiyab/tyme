package tyme_test

import (
	"testing"
	"time"

	"github.com/seiyab/tyme"
)

func TestDate(t *testing.T) {
	date := tyme.NewLocalDate(2020, time.January, 10)

	if date.Year() != 2020 {
		t.Errorf("date.Year() = %q, want %q", date.Year(), 2020)
	}
	if date.Month() != time.January {
		t.Errorf("date.Month() = %q, want %q", date.Month(), time.January)
	}
	if date.Date() != 10 {
		t.Errorf("date.Date() = %q, want %q", date.Date(), 10)
	}
}
