package problemsolving

import (
	utils "problemsolving/utils"
)

func EncodeOccurrences(text string, words []string) {

	positionWordMapping := make([]int, len(text))
	for i, word := range words {
		result := find(text, word)
		if result >= 0 {
			positionWordMapping[result] = i + 1
		}
	}

	offsetMapping := make([]int, len(text))
	count := 0
	for i := range text {
		if positionWordMapping[i] == 0 {
			continue
		}
		count++
		index := positionWordMapping[i] - 1
		word := words[index]

		for j := range word {
			offsetMapping[i+j] = count
		}
	}

	for i, ch := range text {
		utils.PrintRune(applyOffset(ch, offsetMapping[i]))
	}

	utils.PrintRune('\n')
}

func find(text string, word string) int {
	runesText := []rune(text)
	runesWord := []rune(word)

	wordIndex := 0
	startPos := -1

	for i := 0; i < len(text); i++ {
		if !isEqualIgnoreCase(runesText[i], runesWord[wordIndex]) {
			wordIndex = 0
			if startPos != -1 {
				i = startPos
			}
			startPos = -1
			continue
		}
		wordIndex++
		if startPos == -1 {
			startPos = i
		}
		if wordIndex == len(word) {
			return startPos
		}
	}
	return -1
}

func applyOffset(r rune, offsetAmount int) rune {
	offsetAdded := r + rune(offsetAmount)
	if isLower(r) && offsetAdded > 'z' {
		return offsetAdded - 26
	}
	if isUpper(r) && offsetAdded > 'Z' {
		return offsetAdded - 26
	}
	return offsetAdded
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func toLower(r rune) rune {
	if isUpper(r) {
		return r - 'A' + 'a'
	}
	return r
}

func isEqualIgnoreCase(r1 rune, r2 rune) bool {
	return toLower(r1) == toLower(r2)
}
