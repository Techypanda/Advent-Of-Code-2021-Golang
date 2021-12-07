package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Submarine struct {
	Readings string
	scanner  *bufio.Scanner
}

func (s *Submarine) ReadFromFile(relativeLoc string) error {
	dat, err := os.ReadFile(relativeLoc)
	s.Readings = string(dat)
	return err
}

func (s *Submarine) recurse(previous *float64, count int, lineNumber int) (int, error) {
	if s.scanner.Scan() {
		current, err := strconv.ParseFloat(s.scanner.Text(), 64)
		if err != nil {
			panic(fmt.Errorf("error on line %d in file input - %s", lineNumber, err.Error()))
		}
		if previous != nil && current > *previous {
			count += 1
		}
		return s.recurse(&current, count, lineNumber+1)
	} else {
		return count, s.scanner.Err()
	}
}

func (s *Submarine) recurseSliding(previous *Truple, count int, lineNumber int) (int, error) {
	if s.scanner.Scan() {
		current, err := strconv.ParseFloat(s.scanner.Text(), 64)
		if err != nil {
			panic(fmt.Errorf("error on line %d in file input - %s", lineNumber, err.Error()))
		}
		previousComputation := previous.ComputeSum()
		previous.SetLatestVal(current)
		currentComputation := previous.ComputeSum()
		if previousComputation != nil && currentComputation != nil && *previousComputation < *currentComputation {
			count += 1
		}
		return s.recurseSliding(previous, count, lineNumber+1)
	} else {
		return count, s.scanner.Err()
	}
}

func (s *Submarine) ComputeSlidingWindow(relativeLoc string) (int, error) {
	file, err := os.Open(relativeLoc)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	s.scanner = scanner
	truple := Truple{}
	return s.recurseSliding(&truple, 0, 1)
}

func (s *Submarine) ComputeIncreasing(relativeLoc string) (int, error) {
	file, err := os.Open(relativeLoc)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	s.scanner = scanner
	return s.recurse(nil, 0, 1)
}
