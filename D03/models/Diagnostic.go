package models

import (
	"os"
	"strings"
)

type Diagnostic struct {
	rawBits [][]string
}

func NewDiagnostic(relLocation string) *Diagnostic {
	diag := Diagnostic{}
	dat, err := os.ReadFile(relLocation)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")
	horizontalLength := len(strings.ReplaceAll(lines[0], " ", ""))
	verticalLength := len(lines)
	rawBits := make([][]string, verticalLength)
	for i := range rawBits {
		rawBits[i] = make([]string, horizontalLength)
	}
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			rawBits[y][x] = char
		}
	}
	diag.rawBits = rawBits
	return &diag
}

// gT then return largest else smallest
func (d *Diagnostic) ComputeCommonBit(xPos int, gT bool) string {
	key := ""
	var count int
	counter := d.getCountMap(xPos)
	for k := range counter {
		if key == "" || (gT && counter[k] > count) {
			key = k
			count = counter[k]
		} else if !gT && counter[k] < count {
			key = k
			count = counter[k]
		}
	}
	return key
}

func computeCommonBit(gT bool, countMap map[string]int) string {
	key := ""
	var count int
	counter := countMap
	for k := range counter {
		if key == "" || (gT && counter[k] > count) {
			key = k
			count = counter[k]
		} else if !gT && counter[k] < count {
			key = k
			count = counter[k]
		}
	}
	return key
}

func oxygenRecurse(remainingRows [][]string, pos int) []string {
	filteredRows := [][]string{}
	if len(remainingRows) > 1 {
		cMap := getCountMap(remainingRows, pos)
		if cMap["0"] == cMap["1"] {
			cMap["1"] += 1
		}
		commonBit := computeCommonBit(true, cMap)
		for _, row := range remainingRows {
			if row[pos] == commonBit {
				filteredRows = append(filteredRows, row)
			}
		}
		return oxygenRecurse(filteredRows, pos+1)
	} else {
		return remainingRows[0]
	}
}

func co2Recurse(remainingRows [][]string, pos int) []string {
	filteredRows := [][]string{}
	if len(remainingRows) > 1 {
		cMap := getCountMap(remainingRows, pos)
		if cMap["0"] == cMap["1"] {
			cMap["0"] -= 1
		}
		commonBit := computeCommonBit(false, cMap)
		for _, row := range remainingRows {
			if row[pos] == commonBit {
				filteredRows = append(filteredRows, row)
			}
		}
		return co2Recurse(filteredRows, pos+1)
	} else {
		return remainingRows[0]
	}
}

func (d *Diagnostic) CO2() string {
	allRows := d.rawBits
	outStr := ""
	for _, char := range co2Recurse(allRows, 0) {
		outStr += char
	}
	return outStr
}

func (d *Diagnostic) Oxygen() string {
	allRows := d.rawBits
	outStr := ""
	for _, char := range oxygenRecurse(allRows, 0) {
		outStr += char
	}
	return outStr
}

func (d *Diagnostic) ComputeGammaRate() string {
	gamma := ""
	for i := 0; i < len(d.rawBits[0]); i++ {
		gamma += d.ComputeCommonBit(i, true)
	}
	return gamma
}

func (d *Diagnostic) ComputeEpislonRate() string {
	epsilon := ""
	for i := 0; i < len(d.rawBits[0]); i++ {
		epsilon += d.ComputeCommonBit(i, false)
	}
	return epsilon
}

func getCountMap(rows [][]string, xPos int) map[string]int {
	counter := map[string]int{}
	for _, row := range rows {
		counter[row[xPos]] += 1
	}
	return counter
}

func (d *Diagnostic) getCountMap(xPos int) map[string]int {
	counter := map[string]int{}
	for _, row := range d.rawBits {
		counter[row[xPos]] += 1
	}
	return counter
}
