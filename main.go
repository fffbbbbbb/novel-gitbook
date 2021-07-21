package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	expression = "第.*章 "
	filename   = "1.txt"
)

const (
	chapTemp     = `* [%s](%s.md)`
	summaryStart = `# Summary

* [简介](README.md)
`
	output     = "./output/"
	InputSpace = "/gitbook-build/"
)

func init() {
	if os.Getenv("filename") != "" {
		filename = os.Getenv("filename")
	}
	if os.Getenv("expression") != "" {
		expression = os.Getenv("expression")
	}
}

func main() {
	_, err := os.Stat(output)
	if err != nil {
		os.Mkdir(output, os.ModePerm)
	}

	summary := summaryStart
	f, err := os.Open(InputSpace + filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	re, err := regexp.Compile(expression)
	if err != nil {
		panic(err)
	}
	topic := "README.md"
	now := ""
	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = strings.TrimSpace(line)
		if re.MatchString(line) {
			makeFile(output+topic, now)
			now = ""
			topic = line + ".md"
			summary = summary + fmt.Sprintf(chapTemp, line, line) + "\n"
		}
		now += line + "\n"

	}
	makeFile(output+topic, now)
	makeFile(output+"SUMMARY.md", summary)

}

func makeFile(topic, content string) {
	f, err := os.Create(topic)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
}
