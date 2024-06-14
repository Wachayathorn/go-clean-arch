package util

import "time"

type TimeUtil interface {
	Now() time.Time
}

type timeUtil struct {
}

func NewTimeUtil() TimeUtil {
	return &timeUtil{}
}

func (t *timeUtil) Now() time.Time {
	return time.Now()
}
