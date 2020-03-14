package converttyme_test

import (
	"testing"
	"time"

	"github.com/seiyab/tyme"
	"github.com/seiyab/tyme/converttyme"
)

func TestConvertFromLocalYear(t *testing.T) {
	year := tyme.NewLocalYear(2020)

	actualMonth := converttyme.FromLocalYear(year).ToLocalMonth()
	expectedMonth := tyme.NewLocalMonth(2020, time.January)
	if actualMonth != expectedMonth {
		t.Errorf("actualMonth = %q, want %q", actualMonth, expectedMonth)
	}
}
