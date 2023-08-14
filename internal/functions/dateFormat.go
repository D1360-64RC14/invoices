package functions

import (
	"fmt"
	"time"
)

func DateFormat(date time.Time) string {
	return fmt.Sprintf("%02d/%02d/%d", date.Day(), date.Month(), date.Year())
}
