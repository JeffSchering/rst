package main

import (
	"io/ioutil"
	"testing"
)

func TestProcessRstFile(t *testing.T) {
	fileIn := "tests/test-underline-01-in.rst"
	fileOut := "tests/test-underline-01-out.txt"
	RunTest(t, fileIn, fileOut)
	fileIn = "tests/test-underline-02-in.rst"
	fileOut = "tests/test-underline-02-out.txt"
	RunTest(t, fileIn, fileOut)
	fileIn = "tests/test-underline-03-in.rst"
	fileOut = "tests/test-underline-03-out.txt"
	RunTest(t, fileIn, fileOut)
}

func RunTest(t *testing.T, fileIn string, fileOut string) {
	dataIn, err := ioutil.ReadFile(fileIn)
	if err != nil {
		t.Errorf("Error opening %s\n", fileIn)
	}
	dataOut, err := ioutil.ReadFile(fileOut)
	if err != nil {
		t.Errorf("Error opening %s\n", fileOut)
	}
	status, processedFile := ProcessRstFile(string(dataIn))
	if status == "ok" && processedFile != string(dataOut) {
		t.Errorf("%s: Output does not match expected output", fileIn)
		t.Errorf("%s\n\n%s", processedFile, string(dataOut))
	}

}
