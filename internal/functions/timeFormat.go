package functions

import (
	"fmt"
	"time"
)

func TimeFormat(date time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", date.Hour(), date.Minute(), date.Second())
}
