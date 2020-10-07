package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

const slash = `\`

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) { //nolint:gocognit
	var strBuilder strings.Builder
	var charSequance = []rune(str)
	var lastCharPos = len(charSequance) - 1 // char positions starts from 0
	var needToCheckLastChar bool

	if lastCharPos < 0 { // if empty string
		return "", nil
	}

	for i := 0; i < lastCharPos; i++ {
		if string(charSequance[i]) == slash && i+2 <= lastCharPos { // if current char escaped next char(s) and we've got more than 2 chars before end of string
			if (unicode.IsDigit(charSequance[i+1]) && unicode.IsDigit(charSequance[i+2])) || // next two chars are numeric, so got escaped numeric and repeat count
				(string(charSequance[i+1]) == slash && unicode.IsDigit(charSequance[i+2])) { // next two chars are slash and numeric, so got escaped slash and repeat count
				strBuilder.WriteString(repeatChar(charSequance[i+1], charSequance[i+2]))
				i += 2
				continue
			}
		}

		if string(charSequance[i]) == slash { // if we've just got escaped numeric symbol or slash
			if unicode.IsDigit(charSequance[i+1]) || string(charSequance[i+1]) == slash { // if we've just got escaped numeric symbol or slash
				strBuilder.WriteRune(charSequance[i+1])
				i++
				if i+1 == lastCharPos { // if only last char is unchecked
					needToCheckLastChar = true
				}
				continue
			} else {
				return "", ErrInvalidString
			}
		}

		if unicode.IsDigit(charSequance[i]) { // if current char is separated unescaped numeric
			return "", ErrInvalidString
		}

		if unicode.IsDigit(charSequance[i+1]) { // if next char is repeat count for current char
			strBuilder.WriteString(repeatChar(charSequance[i], charSequance[i+1]))
			i++
		} else {
			strBuilder.WriteRune(charSequance[i])
		}

		if i+1 == lastCharPos { // if only last char is unchecked
			needToCheckLastChar = true
		}
	}

	if needToCheckLastChar {
		if !unicode.IsDigit(charSequance[lastCharPos]) && string(charSequance[lastCharPos]) != slash {
			strBuilder.WriteRune(charSequance[lastCharPos])
		} else {
			return "", ErrInvalidString
		}
	}

	return strBuilder.String(), nil
}

func repeatChar(char, count rune) string {
	number, err := strconv.Atoi(string(count))
	if err != nil {
		log.Fatalf("Smth goes wrong. Can't convert %c, to numeral", count)
	}
	return strings.Repeat(string(char), number)
}
