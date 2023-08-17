package helpers

import (
	"time"
)

func TimeToString(dateStr *string, date *time.Time) {
	*dateStr = date.Format(time.RFC3339Nano)
}
