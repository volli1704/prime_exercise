package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
)

// JsonReaderToArray converts reader with JSON array data in it to array
func JsonReaderToArray(input io.Reader) ([]interface{}, error) {
	var res []interface{}

	bb := new(bytes.Buffer)
	_, err := bb.ReadFrom(input)
	if err != nil {
		return nil, err
	}
	raw := bb.Bytes()

	err = json.Unmarshal(raw, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// RoundedFloatToI check if float number can be converted to integer without precision loss
func RoundedFloatToI(num float64) (int, error) {
	if math.Ceil(num) == num {
		return int(num), nil
	}

	return 0, fmt.Errorf("float %f can't be rounded", num)
}
