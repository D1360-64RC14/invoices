package functions

import (
	"fmt"
	"strings"
)

func PriceFormat(price float32) string {
	integerPortion := Reverse(fmt.Sprintf("%.0f", price))
	decimalPortion := strings.Split(fmt.Sprintf("%.2f", price), ".")[1]

	integerPortionDotted := ""

	for i := len(integerPortion); i > 0; i-- {
		if i%3 == 0 && i < len(integerPortion) {
			integerPortionDotted += "."
		}

		integerPortionDotted += string(integerPortion[i-1])
	}

	return fmt.Sprintf(
		"R$ %s,%s",
		integerPortionDotted,
		decimalPortion,
	)
}
