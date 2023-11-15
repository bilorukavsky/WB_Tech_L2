package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", " ", errors.New("Некорректная строка")},
		{"", "", nil},
	}

	for _, tc := range testCases {
		result, err := UnpackString(tc.input)

		if result == tc.expected {
			t.Errorf("Для входной строки %s ожидался результат %s, получено %s", tc.input, tc.expected, result)
		}

		if err == nil && tc.err != nil {
			t.Errorf("Для входной строки %s ожидалась ошибка: %v, получена ошибка nil", tc.input, tc.err)
		}

		if err != nil && tc.err == nil {
			t.Errorf("Для входной строки %s ожидался успешный результат, получена ошибка: %v", tc.input, err)
		}

		if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("Для входной строки %s ожидалась ошибка: %v, получена ошибка: %v", tc.input, tc.err, err)
		}
	}
}
