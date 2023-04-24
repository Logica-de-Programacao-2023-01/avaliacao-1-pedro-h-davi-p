package q1

import "fmt"

func DivideWatermelon(weight int) (bool, error) {

	if weight <= 2 {
		return false, fmt.Errorf("o peso precisa ser maior que 0")
	}
	if weight%2 == 0 {
		return true, nil
	}

	return false, nil
}
