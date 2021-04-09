package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Alphabet map[rune]int

var alphabets = map[string]Alphabet{
	"en": {
		'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7, 'i': 8, 'j': 9, 'k': 10, 'l': 11, 'm': 12, 'n': 13, 'o': 14, 'p': 15, 'q': 16, 'r': 17, 's': 18, 't': 19, 'u': 20, 'v': 21, 'w': 22, 'x': 23, 'y': 24, 'z': 25,
	},
	"ru": {
		'ё': 6, 'ъ': 27, 'б': 1, 'в': 2, 'д': 4, 'о': 15, 'т': 19, 'у': 20, 'ю': 31, 'я': 32, 'х': 22, 'ц': 23, 'щ': 26, 'й': 10, 'р': 17, 'а': 0, 'г': 3, 'к': 11, 'н': 14, 'п': 16, 'и': 9, 'м': 13, 'с': 18, 'е': 5, 'ж': 7, 'з': 8, 'ф': 21, 'ч': 24, 'ы': 28, 'л': 12, 'ш': 25, 'ь': 29, 'э': 30,
	},
}

func (al *Alphabet) getChar(i int) rune {
	for char, index := range *al {
		if index == i {
			return char
		}
	}
	return 32
}

func Shift(char rune, shift int, lang string) rune {
	if !unicode.IsLetter(char) {
		return char
	}
	upper := unicode.IsUpper(char)
	char = unicode.ToLower(char)
	alphabet := alphabets[lang]
	length := len(alphabet)
	newPlace := (alphabet[char] + shift) % length
	if newPlace < 0 {
		char = alphabet.getChar(newPlace + length)
	} else {
		char = alphabet.getChar(newPlace)
	}
	if upper {
		return unicode.ToUpper(char)
	} else {
		return char
	}
}

func Caesar(str string, shift int, lang string, encode bool) string {
	res := ""
	for _, value := range str {
		if encode {
			res += string(Shift(value, shift, lang))
		} else {
			res += string(Shift(value, -shift, lang))
		}

	}
	return res
}

func Vigenere(str string, key string, lang string, encode bool) string {
	alphabet := alphabets[lang]
	length := len(alphabet)
	// full key
	fullKey := strings.ToLower(key)
	for len(fullKey) < len(str) {
		fullKey += fullKey
	}
	fullKey = fullKey[:len(str)]
	// case map
	caseMap := make([]bool, len(str))
	for index, value := range str {
		caseMap[index] = unicode.IsUpper(value)
	}
	str = strings.ToLower(str)

	res_raw, res, char := "", "", ' '
	for index, value := range str {
		m := alphabet[value]
		k := alphabet[rune(fullKey[index])]
		if encode {
			char = alphabet.getChar((m + k) % length)
		} else {
			char = alphabet.getChar((m + length - k) % length)
		}
		res_raw += string(char)
	}
	for index, isUpper := range caseMap {
		if isUpper {
			res += string(unicode.ToUpper(rune(res_raw[index])))
		} else {
			res += string(res_raw[index])
		}
	}

	return res
}

func encodeInput() bool {
	fmt.Printf("Choose direction:\n1: encode\n2: decode\n")
	scanner := bufio.NewScanner(os.Stdin)
	choice := ""
	if scanner.Scan() {
		choice = scanner.Text()
	}
	encode := false
	switch string(choice) {
	case "1":
		encode = true
	case "2":
		encode = false
	default:
		panic("wrong input")
	}
	return encode
}

func langInput() string {
	fmt.Printf("Choose language:\n1: ru\n2: eng\n")
	scanner := bufio.NewScanner(os.Stdin)
	choice := ""
	if scanner.Scan() {
		choice = scanner.Text()
	}
	lang := ""
	switch string(choice) {
	case "1":
		lang = "ru"
	case "2":
		lang = "en"
	default:
		panic("not yet implemetned")
	}
	return lang
}

func main() {
	result, inputReader, scanner := "", bufio.NewReader(os.Stdin), bufio.NewScanner(os.Stdin)
	input, output := "input.txt", "output.txt"
	fmt.Printf("Input file path: ")
	if scanner.Scan() {
		input = scanner.Text()
	}
	fmt.Printf("Output file path: ")
	if scanner.Scan() {
		output = scanner.Text()
	}
	str, err := ioutil.ReadFile(string(input))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Choose method:\n1: Caesar\n2: Vigenere\n")
	choice, _, _ := inputReader.ReadLine()
	switch string(choice) {
	case "1":
		fmt.Printf("Shift length (can be less than 0): ")
		shiftRaw, _, _ := inputReader.ReadLine()
		shift, _ := strconv.Atoi(string(shiftRaw))
		lang := langInput()
		encode := encodeInput()
		result = Caesar(string(str), shift, lang, encode)
	case "2":
		fmt.Printf("Key: ")
		key, _, _ := inputReader.ReadLine()
		lang := langInput()
		encode := encodeInput()
		result = Vigenere(string(str), string(key), lang, encode)
	default:
		panic("wrong input")
	}
	fo, _ := os.Create(string(output))
	fo.Write([]byte(result))
	fmt.Printf("\n---done---\n")
}
