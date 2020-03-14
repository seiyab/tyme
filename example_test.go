package tyme_test

import (
	"fmt"
	"time"

	"github.com/seiyab/tyme"
)

func ExampleLocalYear() {
	year := tyme.NewLocalYear(2006)
	fmt.Println(year)

	// Output:
	// 2006
}

func ExampleLocalMonth() {
	month := tyme.NewLocalMonth(2006, time.January)
	fmt.Println(month)

	// Output:
	// 2006-01
}

func ExampleLocalDate() {
	date := tyme.NewLocalDate(2006, time.January, 2)
	fmt.Println(date)

	// Output:
	// 2006-01-02
}