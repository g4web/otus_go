package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(stringForUnpack string) (string, error) {
	runes := []rune(stringForUnpack)

	lastRuneIndex := len(runes) - 1

	firstRuneIndex := 0

	var resultBuilder strings.Builder

	for runeIndex, thisRune := range stringForUnpack {
		// если первый символ число, то ошибка
		if runeIndex == firstRuneIndex {
			if unicode.IsDigit(thisRune) {
				return "", ErrInvalidString
			}
		}

		// добавляем последний символ в резултат если это буква
		if runeIndex == lastRuneIndex {
			if !unicode.IsDigit(thisRune) {
				resultBuilder.WriteRune(thisRune)
			}
			continue
		}

		nextRune := runes[runeIndex+1]

		// если подряд два числа, то ошибка
		if unicode.IsDigit(thisRune) && unicode.IsDigit(nextRune) {
			return "", ErrInvalidString
		}

		// если текущий символ число, то не добавляем его в результат
		if unicode.IsDigit(thisRune) {
			continue
		}

		// Выводит текущий символ один или несколько раз
		if unicode.IsDigit(nextRune) {
			numberOfRepeats, _ := strconv.Atoi(string(nextRune))
			for i := 0; i < numberOfRepeats; i++ {
				resultBuilder.WriteRune(thisRune)
			}
		} else {
			resultBuilder.WriteRune(thisRune)
		}
	}

	result := resultBuilder.String()

	return result, nil
}
