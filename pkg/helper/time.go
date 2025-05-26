package helper

import (
	"time"
	"fmt"
)

func ParseOptionalTimeRFC3339(s *string) (*time.Time, error) {
	if s == nil || *s == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, fmt.Errorf("invalid time format: %w", err)
	}
	return &t, nil
}
