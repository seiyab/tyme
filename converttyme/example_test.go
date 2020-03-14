package converttyme_test

import (
	"fmt"

	"github.com/seiyab/tyme"
	"github.com/seiyab/tyme/converttyme"
)

func ExampleFromLocalYear() {
	year := tyme.NewLocalYear(2006)
	fmt.Println(year)

	time := converttyme.FromLocalYear(year).ToTime()
	fmt.Println(time)

	month := converttyme.FromLocalYear(year).ToLocalMonth()
	fmt.Println(month)
}
