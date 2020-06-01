package main

import "fmt"

type crcedit struct {
	strategy, filetype, algorithm, checksum      string
	inputFilename, outputFilename, magicFilename string
	shouldOverwrite, shouldConfirm               bool
}

func (c crcedit) Run() {
	fmt.Println("crcedit.Run(): ran")
	return
}
