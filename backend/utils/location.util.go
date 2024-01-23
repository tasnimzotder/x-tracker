package utils

import "strconv"

func ConvStrToFloat(value string) float32 {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	return float32(val)
}
