package main

import "fmt"

type crcedit struct {
	algorithm, checksum, filetype                          string
	shouldOverwrite, shouldSkipConfirm, shouldPreserveSize bool
}

func (c crcedit) run(args []string) {
	fmt.Printf("crcedit.run(): ran with args %v", args)
	return
}
