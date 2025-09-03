package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Transform(data string) (string, error) {

	//проверка на наличие символов
	if len(data) == 0 {
		return "", errors.New("error data")
	}

	//проверка на наличие символов, отличных от Морзе
	check := func(chk rune) bool {
		return chk != '-' && chk != '.' && !unicode.IsSpace(chk)
	}
	//опредление входного текста
	isMorse := strings.ContainsFunc(data, check)

	if isMorse {
		return morse.ToMorse(data), nil
	} else {
		return morse.ToText(data), nil
	}

}
