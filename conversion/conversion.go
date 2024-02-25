package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringValue := range strings {
		floatPrice, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, errors.New("Failed to Conver String to Float!")
		}

		floats = append(floats, floatPrice)
	}

	return floats, nil
}
