package main

import (
	"errors"
)

type pars map[rune]rune

func newOpenPars() pars {
	return map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
}

func newClosePars() pars {
	return map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
}

type Stack struct {
	size int
	data []rune
}

func NewStack(len int) Stack {
	return Stack{
		size: 0,
		data: make([]rune, len, len),
	}
}

func (s *Stack) pop() (rune, error) {
	if s.size == 0 {
		return 0, errors.New("stack is empty")
	}
	s.size--
	return s.data[s.size], nil
}

func (s *Stack) push(value rune) error {
	if s.size+1 > cap(s.data) {
		return errors.New("stack is full")
	}
	s.data[s.size] = value
	s.size++
	return nil
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	runes := []rune(s)
	l := len(runes)

	stackOpen := NewStack(l)
	openPars := newOpenPars()
	closePars := newClosePars()

	for i := 0; i < l; i++ {
		_, found := openPars[runes[i]]
		if found {
			err := stackOpen.push(runes[i])
			if err != nil {
				return false
			}
			continue
		}

		openPar, _ := closePars[runes[i]]
		v, err := stackOpen.pop()
		if err != nil {
			return false
		}
		if openPar != v {
			return false
		}
	}
	if !stackOpen.isEmpty() {
		return false
	}

	return true
}
