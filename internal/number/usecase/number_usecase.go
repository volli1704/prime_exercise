package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/volli1704/prime_api/internal/utils"
	"github.com/volli1704/prime_api/pkg/primechecker"
)

var ErrWrongElementFormat = errors.New("wrong format of element")

// Usecases for numbers
type NumberUseCase struct {
	checker primechecker.PrimeChecker
}

func NewNumberUseCase(checker primechecker.PrimeChecker) NumberUseCase {
	return NumberUseCase{checker}
}

// FindPrimesForArray receive array of interfaces, ensures that all values in
// the array can be casted to number and check is this number prime or not
func (u *NumberUseCase) FindPrimesForArray(input []interface{}) ([]bool, error) {
	res := make([]bool, len(input))
	intArr, err := u.ensureIntArray(input)
	if err != nil {
		return nil, err
	}

	for i, num := range intArr {
		res[i] = u.checker.Check(num)
	}

	return res, nil
}

func (u *NumberUseCase) ensureIntArray(input []interface{}) ([]int, error) {
	res := make([]int, len(input))

	for i, v := range input {
		intVal, err := u.ensureInt(v)
		if err != nil {
			return nil, fmt.Errorf("the given input is invalid. Element on index %d is not a number", i+1)
		}

		res[i] = intVal
	}

	return res, nil
}

func (u *NumberUseCase) ensureInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case int:
		return v, nil
	case string:
		intVal, err := strconv.Atoi(v)
		if err == nil {
			return intVal, nil
		}

		floatVal, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return utils.RoundedFloatToI(floatVal)
		}
		return strconv.Atoi(v)
	case float64:
		val := v
		return utils.RoundedFloatToI(val)
	default:
		return 0, ErrWrongElementFormat
	}
}
