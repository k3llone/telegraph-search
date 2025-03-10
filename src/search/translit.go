package search

import (
	"strings"
	"unicode"
)

func Translit(text string) string {
	res := ""
	translit_table := map[rune]string{
		'а': "a",
		'б': "b",
		'в': "v",
		'г': "g",
		'д': "d",
		'е': "e",  //
		'ё': "yo", //
		'ж': "zh", //
		'з': "z",
		'и': "i",
		'й': "j",
		'к': "k",
		'л': "l",
		'м': "m",
		'н': "n",
		'о': "o",
		'п': "p",
		'р': "r",
		'с': "s",
		'т': "t",
		'у': "u", //
		'ф': "f",
		'х': "h",
		'ц': "c",    //
		'ч': "ch",   //
		'ш': "sh",   //
		'щ': "shch", //
		'ъ': "",
		'ы': "y", //
		'ь': "",
		'э': "e",  //
		'ю': "yu", //
		'я': "ya", //
		'1': "1",
		'2': "2",
		'3': "3",
		'4': "4",
		'5': "5",
		'6': "6",
		'7': "7",
		'8': "8",
		'9': "9",
		'0': "0",
		'-': "-",
		'a': "a",
		'b': "b",
		'c': "c",
		'd': "d",
		'e': "e",
		'f': "f",
		'g': "g",
		'h': "h",
		'i': "i",
		'j': "j",
		'k': "k",
		'l': "l",
		'm': "m",
		'n': "n",
		'o': "o",
		'p': "p",
		'q': "q",
		'r': "r",
		's': "s",
		't': "t",
		'u': "u",
		'v': "v",
		'w': "w",
		'x': "x",
		'y': "y",
		'z': "z",
	}

	for _, v := range text {
		if !unicode.IsLower(v) {
			res += strings.ToUpper(translit_table[unicode.ToLower(v)])
		} else {
			res += translit_table[v]
		}
	}

	return res
}
