package main

import (
	"bytes"
	"testing"
)

var correctCases = []struct {
	in  string
	out string
}{
	{"1 2 3 4 + * + =", "Result = 15\n"},
	{"1 2 + 3 4 + * =", "Result = 21\n"},
	{"111++=", "Result = 3\n"},
	{"1 1 1 + + =", "Result = 3\n"},
	{"9 7 / =", "Result = 1\n"},
	{"8 2 / =", "Result = 4\n"},
	{"1 6 + =", "Result = 7\n"},
	{"6 1 - =", "Result = 5\n"},
	{"1 6 - =", "Result = -5\n"},
	{"2 8 * =", "Result = 16\n"},
	{"  2      8    *=", "Result = 16\n"},
}

func TestCorrectCases(t *testing.T) {
	for _, tCase := range correctCases {
		t.Run(tCase.in, func(t *testing.T) {
			in := new(bytes.Buffer)
			out := new(bytes.Buffer)

			in.WriteString(tCase.in)
			err := Calc(in, out)
			if err != nil {
				t.Errorf("test for OK Failed - error in Calc:\n%v", err)
			}
			result := out.String()
			if result != tCase.out {
				t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, tCase.out)
			}
		})
	}
}

var inCorrectCases = []string{
	"20 2 / =",
	"skldjagh",
	"-2 6 + =",
	"=",
}

func TestInCorrectCases(t *testing.T) {
	for _, tCase := range inCorrectCases {
		t.Run(tCase, func(t *testing.T) {
			in := new(bytes.Buffer)
			out := new(bytes.Buffer)

			in.WriteString(tCase)
			err := Calc(in, out)
			if err == nil {
				t.Errorf("test for FAIL Failed with input:\n%v", tCase)
			}
		})
	}
}
