package util

import (
	"strconv"
)

func StringToFloat64(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return float64(0), err
	}
	return f, nil
}
