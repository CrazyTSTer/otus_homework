package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var strBuilder strings.Builder
	var charSequance = []rune(str)
	var lastCharPos = len(charSequance) - 1 //char positions starts from 0
	needToCheckLastChar := false
	if lastCharPos < 0 { //if empty string
		return "", nil
	}

	for i := 0; i < lastCharPos; i++ {
		//needToCheckLastChar = false
		if string(charSequance[i]) == "\\" { //if current char escaped next char(s)
			if i+2 <= lastCharPos { //if we've got more than 2 chars before end of string
				if unicode.IsDigit(charSequance[i+1]) && unicode.IsDigit(charSequance[i+2]) { //next two chars are numeric, so got escaped numeric and repeat count
					number, err := strconv.Atoi(string(charSequance[i+2]))
					if err != nil {
						log.Fatalf("Smth goes wrong. Can't convert %c, to numeral", charSequance[i+2])
					}
					strBuilder.WriteString(strings.Repeat(string(charSequance[i+1]), number))
					i += 2
				} else if string(charSequance[i+1]) == "\\" && unicode.IsDigit(charSequance[i+2]) { //next two chars are slash and numeric, so got escaped slash and repeat count
					number, err := strconv.Atoi(string(charSequance[i+2]))
					if err != nil {
						log.Fatalf("Smth goes wrong. Can't convert %c, to numeral", charSequance[i+2])
					}
					strBuilder.WriteString(strings.Repeat(string(charSequance[i+1]), number))
					i += 2
				} else if unicode.IsDigit(charSequance[i+1]) || string(charSequance[i+1]) == "\\" { //if we've just got escaped numeric symbol or slash
					strBuilder.WriteRune(charSequance[i+1])
					i++
				} else {
					return "", ErrInvalidString
				}
			} else { //if current char is the penultimate in the string
				if unicode.IsDigit(charSequance[i+1]) || string(charSequance[i+1]) == "\\" { //if last char numeric or slash
					strBuilder.WriteRune(charSequance[i+1])
					i++
				} else {
					return "", ErrInvalidString
				}
			}
		} else {
			if !unicode.IsDigit(charSequance[i]) && unicode.IsDigit(charSequance[i+1]) { //if current char is not numeric and got repeat count
				number, err := strconv.Atoi(string(charSequance[i+1]))
				if err != nil {
					log.Fatalf("Smth goes wrong. Can't convert %c, to numeral", charSequance[i+1])
				}
				strBuilder.WriteString(strings.Repeat(string(charSequance[i]), number))
				i++
			} else if unicode.IsDigit(charSequance[i]) { //if got separate unescaped numeric char
				return "", ErrInvalidString
			} else {
				strBuilder.WriteRune(charSequance[i])
			}
		}
		if i+1 == lastCharPos { //if only last char is unchecked
			needToCheckLastChar = true
		}
	}

	if needToCheckLastChar {
		if unicode.IsDigit(charSequance[lastCharPos]) || string(charSequance[lastCharPos]) == "\\" {
			return "", ErrInvalidString
		} else {
			strBuilder.WriteRune(charSequance[lastCharPos])
		}
	}

	return strBuilder.String(), nil
}
