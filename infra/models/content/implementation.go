package content

import (
	"strings"
)

func (c Model) GetCaracteresFromEachLinesColumns() []string {
	split := splitByNewLine(c)
	array := splitEachCaracteres(split)
	codes := getCodes(array)
	return codes
}

func splitByNewLine(c Model) []string {
	return strings.Split(c.Content, "\n")
}

func splitEachCaracteres(split []string) [][]string {
	array := make([][]string, len(split))
	for i, s := range split {
		currentLine := strings.Split(s, "")
		array[i] = currentLine
	}
	return array
}

func getNumberFromCaractere(array [][]string, i int, j int) string {
	carateres := array[i][j] +
		array[i][j + 1] +
		array[i][j + 2] +
		array[i + 1][j] +
		array[i + 1][j + 1] +
		array[i + 1][j + 2] +
		array[i + 2][j] +
		array[i + 2][j + 1] +
		array[i + 2][j + 2] +
		array[i + 3][j] +
		array[i + 3][j + 1] +
		array[i + 3][j + 2]

	switch carateres {
	case "     |  |   ":
		return "1"
	case " _  _||_    ":
		return "2"
	case " _  _| _|   ":
		return "3"
	case "   |_|  |   ":
		return "4"
	case " _ |_  _|   ":
		return "5"
	case " _ |_ |_|   ":
		return "6"
	case " _   |  |   ":
		return "7"
	case " _ |_||_|   ":
		return "8"
	case " _ |_| _|   ":
		return "9"
	default:
		return "?"
	}
}

func getCodes(array [][]string) []string {
	codes := make([]string, 1)
	for i := 0; i < len(array); i+=4 {
		currentCode := ""
		for j := 0; j < len(array[i]); j+=3 {
			currentCode += getNumberFromCaractere(array, i, j)
		}
		codes = append(codes, currentCode)
	}
	return codes
}
