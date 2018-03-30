package prettytable

import (
        "unicode"
	"math"
	"strings"
	"fmt"
)

func countChineseChar(str string) int {
    total := 0
    for _, r := range str {
        if unicode.Is(unicode.Scripts["Han"], r) {
            total++
        }
    }
    return total
}

func stringLength(str string) int {
    chinese := countChineseChar(str)
    return len([]rune(str)) + chinese
}


func calculateSectionLengths(labels []string, table [][]string) []int {
	largest := make([]int, len(labels))
	for i := range largest {
		largest[i] = stringLength(labels[i])
	}

	for i := 0 ; i < len(table) ; i += 1 {
		for j := 0 ; j < len(table[i]) ; j += 1 {
			content := table[i][j]
			if len(content) > largest[j] {
				largest[j] = stringLength(content)
			}
		}
	}

	return largest
}

func getCenteredLabel(sectionLength int, label string) string {
	numWhitespace := sectionLength - stringLength(label)
	numSpace := math.Floor(float64(numWhitespace) / 2)
	str := ""
	str += strings.Repeat(" ", int(numSpace))
	str += label
	str += strings.Repeat(" ", int(numSpace))
	if int(numWhitespace) % 2 != 0 {
		str += " "
	}
	return str
}

func getHeader(labels []string, sectionLengths []int) string {
	accumulator := "\u2503"
	for i, label := range labels {
		accumulator += (" " + getCenteredLabel(sectionLengths[i], label) + " \u2503")
	}
	return accumulator
}

func getTableWidth(sectionLengths []int) int {
	accumulator := 2
	for _, sectionLength := range sectionLengths{
		accumulator += (sectionLength + 2)
	}
	return accumulator
}

func createRegularLine(sectionLengths []int) string {
	accumulator := "\u2502"
	for i, sectionLength := range sectionLengths{
		accumulator += (strings.Repeat("\u2500", sectionLength + 2))
		if i == len(sectionLengths) - 1 {
			accumulator += "\u2502"
		} else {
			accumulator += "\u253C"
		}
	}
	return accumulator
}

func createTopLine(sectionLengths []int) string {
	accumulator := "\u250F"
	for i, sectionLength := range sectionLengths{
		accumulator += (strings.Repeat("\u2501", sectionLength + 2))
		if i == len(sectionLengths) - 1 {
			accumulator += "\u2513"
		} else {
			accumulator += "\u2533"
		}
	}
	return accumulator
}

func createBottomLine(sectionLengths []int) string {
	accumulator := "\u2514"
	for i, sectionLength := range sectionLengths{
		accumulator += (strings.Repeat("\u2500", sectionLength + 2))
		if i == len(sectionLengths) - 1 {
			accumulator += "\u2518"
		} else {
			accumulator += "\u2534"
		}
	}
	return accumulator
}

func createSectionLine(sectionLengths []int) string {
	accumulator := "\u2521"
	for i, sectionLength := range sectionLengths{
		accumulator += (strings.Repeat("\u2501", sectionLength + 2))
		if i == len(sectionLengths) - 1 {
			accumulator += "\u2529"
		} else {
			accumulator += "\u2547"
		}
	}
	return accumulator
}

func getContentPadded(content string, sectionLength int) string {
	if sectionLength > len(content) {
		return content + strings.Repeat(" ", sectionLength - len(content))
	} else {
		return content
	}
}

func PrintTable(labels []string, table [][]string) {
	sectionLengths := calculateSectionLengths(labels, table)
	header := getHeader(labels, sectionLengths)
	topLine := createTopLine(sectionLengths)
	bottomLine := createBottomLine(sectionLengths)
	regularLine := createRegularLine(sectionLengths)
	sectionLine := createSectionLine(sectionLengths)

	fmt.Println(topLine)
	fmt.Println(header)
	fmt.Println(sectionLine)

	for i := range table {
		rowString := "\u2502"
		for j := range table[i] {
			contentPadded := getContentPadded(table[i][j], sectionLengths[j])
			rowString += (" " + contentPadded + " \u2502")
		}
		fmt.Println(rowString)
		if i == len(table) - 1 {
			fmt.Println(bottomLine)
		} else {
			fmt.Println(regularLine)
		}
	}
}
