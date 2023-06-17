package day8

import (
	"sort"
	"strconv"
	"strings"

	"github.com/divxvid/aoc-go-2021/utils"
)

type Query struct {
	jumbledSegments []string
	digits          []string
}

type Day8 struct {
	queries []Query
}

func New(data string) (*Day8, error) {
	queries := make([]Query, 0)

	for _, line := range utils.SplitIntoLines(data) {
		parts := strings.Split(line, "|")
		segments := strings.Split(strings.Trim(parts[0], " "), " ")
		digits := strings.Split(strings.Trim(parts[1], " "), " ")
		for i := range segments {
			segments[i] = sortString(strings.Trim(segments[i], " "))
		}
		for i := range digits {
			digits[i] = sortString(strings.Trim(digits[i], " "))
		}

		queries = append(queries, Query{jumbledSegments: segments, digits: digits})
	}

	return &Day8{
		queries: queries,
	}, nil
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

// Number of Digits:
// 0: 6
// 1: 2
// 2: 5
// 3: 5
// 4: 4
// 5: 5
// 6: 6
// 7: 3
// 8: 7
// 9: 6
func (d *Day8) SolveA() (string, error) {
	result := 0
	for _, query := range d.queries {
		for _, v := range query.digits {
			switch len(v) {
			case 2:
				result += 1
			case 4:
				result += 1
			case 3:
				result += 1
			case 7:
				result += 1
			}
		}
	}
	return strconv.Itoa(result), nil
}

// digits 1, 4, 7, 8 can be found easily
// digit 6 is a string of length 6 where missing character belongs to characters of 1
// digit 0 is a string of length 6 where missing character belongs to characters of 4
// remaining 6 length string is digit 9
// check which character is missing only once. The number of length 5 where it is missing is digit 2
// the one which has all the characters of digit 7 is digit 3, other is digit 5
func (d *Day8) SolveB() (string, error) {
	result := 0
	for _, query := range d.queries {
		ret := solveSingleQuery(&query)
		result += ret
	}
	return strconv.Itoa(result), nil
}

func solveSingleQuery(query *Query) int {
	digitMap := make(map[string]int)
	inverseDigitMap := make(map[int]string)

	//first pass to get the simple ones
	for _, v := range query.jumbledSegments {
		switch len(v) {
		case 2:
			digitMap[v] = 1
			inverseDigitMap[1] = v
		case 4:
			digitMap[v] = 4
			inverseDigitMap[4] = v
		case 3:
			digitMap[v] = 7
			inverseDigitMap[7] = v
		case 7:
			digitMap[v] = 8
			inverseDigitMap[8] = v
		}
	}
	//second pass to get all the 6 digit guys
	sixDigits := make([]string, 0)
	for _, v := range query.jumbledSegments {
		if len(v) != 6 {
			continue
		}
		sixDigits = append(sixDigits, v)
	}

	for len(sixDigits) > 1 {
		toRemove := -1
		for idx, v := range sixDigits {
			missingCharacter := findMissingCharacter(inverseDigitMap[8], v)
			_, ok6 := inverseDigitMap[6]
			if !ok6 && checkForIntersection(missingCharacter, inverseDigitMap[1]) {
				digitMap[v] = 6
				inverseDigitMap[6] = v
				toRemove = idx
				break
			}
			_, ok0 := inverseDigitMap[0]
			if !ok0 && checkForIntersection(missingCharacter, inverseDigitMap[4]) {
				digitMap[v] = 0
				inverseDigitMap[0] = v
				toRemove = idx
				break
			}
		}
		sixDigits = append(sixDigits[:toRemove], sixDigits[toRemove+1:]...) //deletes the element at index toRemove
	}
	digitMap[sixDigits[0]] = 9
	inverseDigitMap[9] = sixDigits[0]

	//last pass to get all the 5 digit bois
	//first lets try to get digit 2

	digit2String := findDigit2String(query.jumbledSegments)
	digitMap[digit2String] = 2
	inverseDigitMap[2] = digit2String

	fiveDigits := make([]string, 0)
	for _, v := range query.jumbledSegments {
		if len(v) != 5 || v == digit2String {
			continue
		}

		fiveDigits = append(fiveDigits, v)
	}

	if checkForIntersection(inverseDigitMap[7], fiveDigits[0]) {
		digitMap[fiveDigits[0]] = 3
		inverseDigitMap[3] = fiveDigits[0]

		digitMap[fiveDigits[1]] = 5
		inverseDigitMap[5] = fiveDigits[1]
	} else {
		digitMap[fiveDigits[0]] = 5
		inverseDigitMap[5] = fiveDigits[0]

		digitMap[fiveDigits[1]] = 3
		inverseDigitMap[3] = fiveDigits[1]
	}

	result := 0
	for _, s := range query.digits {
		result = result*10 + digitMap[s]
	}

	return result
}

func findDigit2String(data []string) string {
	counts := make(map[rune]int)
	for _, s := range data {
		for _, ch := range s {
			counts[ch] += 1
		}
	}

	missingCharacter := ' '
	for k, v := range counts {
		if v == 9 {
			missingCharacter = k
		}
	}

	for _, s := range data {
		found := false
		for _, ch := range s {
			if ch == missingCharacter {
				found = true
				break
			}
		}
		if !found {
			return s
		}
	}
	return " "
}

func findMissingCharacter(v, s string) string {
	for _, chv := range v {
		found := false
		for _, chs := range s {
			if chs == chv {
				found = true
				break
			}
		}
		if !found {
			return string(chv)
		}
	}
	return ""
}

func checkForIntersection(a string, b string) bool {
	for _, cha := range a {
		found := false
		for _, chb := range b {
			if cha == chb {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
