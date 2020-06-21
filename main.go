package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `Usage:
    crcedit -c CHECKSUM FILENAME

By default, crcedit creates a new file with a '.crcedit' extension - the
original file won't be overwritten unless you want it to be.

Options:
    -a, -algorithm ALGORITHM     Select a checksum algorithm.
    -c, -checksum CHECKSUM       Specify the desired checksum.
        -dangerous-no-confirm    Disable all confirmation prompts.
        -dangerous-overwrite     Overwrite the input file.
    -h, -help                    Print this help message.
    -l, -list                    List available algorithms and filetypes.
        -preserve-file-size      Preserve the original file size.
    -t, -type FILETYPE           State the file type, ignoring auto-detection.`

func main() {
	flag.Usage = func() { fmt.Fprintln(os.Stderr, usage) }

	var (
		algorithmFlag, checksumFlag, typeFlag          string
		dangerousNoConfirmFlag, dangerousOverwriteFlag bool
		helpFlag, listFlag, preserveFileSizeFlag       bool
	)

	flag.StringVar(&algorithmFlag, "a", "", "")
	flag.StringVar(&algorithmFlag, "algorithm", "", "")
	flag.StringVar(&checksumFlag, "c", "", "")
	flag.StringVar(&checksumFlag, "checksum", "", "")
	flag.StringVar(&typeFlag, "t", "", "")
	flag.StringVar(&typeFlag, "type", "", "")

	flag.BoolVar(&dangerousOverwriteFlag, "dangerous-overwrite", false, "")
	flag.BoolVar(&dangerousNoConfirmFlag, "dangerous-no-confirm", false, "")

	flag.BoolVar(&helpFlag, "h", false, "")
	flag.BoolVar(&helpFlag, "help", false, "")
	flag.BoolVar(&listFlag, "l", false, "")
	flag.BoolVar(&listFlag, "list", false, "")
	flag.BoolVar(&preserveFileSizeFlag, "preserve-file-size", false, "")

	flag.Parse()

	if helpFlag {
		flag.Usage()
		return
	} else if listFlag {
		fmt.Println("listFlag: print out a list of all supported algorithms and filetypes")
		return
	}

	(&crcedit{
		algorithm: algorithmFlag,
		checksum:  checksumFlag,
		filetype:  typeFlag,

		shouldOverwrite:    dangerousOverwriteFlag,
		shouldSkipConfirm:  dangerousNoConfirmFlag,
		shouldPreserveSize: preserveFileSizeFlag,
	}).run(flag.Args())
}
