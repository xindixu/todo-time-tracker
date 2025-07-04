package accessors

import (
	"fmt"
	"time"
)

// MakeTSRange TODO: use this when querying and creating sessions
func MakeTSRange(start time.Time, end *time.Time) string {
	if end == nil {
		return fmt.Sprintf("[%s,infinity)", start.Format(time.RFC3339))
	}
	return fmt.Sprintf("[%s,%s)", start.Format(time.RFC3339), end.Format(time.RFC3339))
}
