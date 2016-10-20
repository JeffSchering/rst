package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"path/filepath"
	"os"
	"regexp"
	"unicode/utf8"
)

func  main()  {
		arg := os.Args[1]
		fmt.Printf("Starting at dir: %s\n", arg)
		Walk(arg)
}

func Walk(rootdir string) {
		filepath.Walk(rootdir, WalkFunc)
}

func WalkFunc(rootdir string, info os.FileInfo, err error) error {
	s := info.Name()
	if !info.IsDir() && strings.HasSuffix(s, ".rst") {
		OpenRstFile(rootdir)
	}
	return nil
}

func OpenRstFile(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error opening", filepath)
		return
	}
	fmt.Println(filepath)
	status, fileData := ProcessRstFile(string(data))
	if status != "ok" {
		fmt.Println("Error while processing", filepath, "Reason:", status)
	} else {
		ioutil.WriteFile(filepath, []byte(fileData), 0644)
	}
}

func ProcessRstFile(data string) (string, string) {
	lines := strings.SplitAfter(data, "\n")
	var lineLength, overlineLength, titlelineLength int
	for i, line := range lines {
		matched, err := regexp.MatchString("^[=\\-\\`:\\.'\\~\\^_*#\\\"]{4,}", line)
		if err != nil {
			return "regex fail", ""
		}
		if matched {
			underlineCharacter := string(line[0])
			if i == 0 {
				overlineLength = len(strings.TrimSpace(lines[i]))
				titlelineLength = utf8.RuneCountInString(strings.TrimSpace(lines[i+1]))
				if overlineLength > titlelineLength {
					lineLength = overlineLength
				}
			} else if i == 1 || i == 2 {
				lineLength = utf8.RuneCountInString(strings.TrimSpace(lines[i-1]))
				if len(strings.TrimSpace(lines[i])) > lineLength {
					lineLength = len(strings.TrimSpace(lines[i]))
				}
			} else if len(strings.TrimSpace(lines[i-1])) < 1 {
				lineLength = utf8.RuneCountInString(strings.TrimSpace(lines[i+1]))
				if len(strings.TrimSpace(lines[i+2])) > lineLength {
					lineLength = len(strings.TrimSpace(lines[i+2]))
				}
				if len(strings.TrimSpace(lines[i+1])) < 1 {
					lineLength = len(strings.TrimSpace(lines[i]))
				}
			} else if lines[i-2][0] == lines[i][0] {
				lineLength = len(strings.TrimSpace(lines[i-2]))
			}
			line = strings.Repeat(underlineCharacter, lineLength)
			lines[i] = line + "\r\n"
		}
		if !matched {
			lineLength = utf8.RuneCountInString(strings.TrimSpace(line))
		}
	}
	return "ok", strings.Join(lines,"")
}
