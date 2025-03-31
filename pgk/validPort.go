package pgk

import (
	"fmt"
	"strconv"
)

func ValidPort(input string) (port int, err error) {
	if input == "" {
		err = fmt.Errorf("ошибка: Вы ничего не ввели")
		return 0, err
	}

	port, err = strconv.Atoi(input)
	if err != nil {
		err = fmt.Errorf("ошибка: Введено не целое число")
		return 0, err
	}

	if port < 1 || port > 65535 {
		err = fmt.Errorf("ошибка: Номер порта должен быть в диапазоне от 1 до 65535")
		return 0, err
	}

	return port, nil
}
