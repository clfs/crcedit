package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const usage = `Usage:
    crcedit FILENAME -c CHECKSUM -a ALGORITHM

By default, crcedit will create a new file with a '.crcedit' extension - your
files won't be overwritten unless you specify it. Additional documentation can
be found at https://github.com/clfs/crcedit inside the README.

Options:
	-a, --algorithm ALGORITHM     Pick a checksum algorithm.
	-c, --checksum                Specify the desired checksum.
	    --dangerous-no-confirm    Disable all confirmation prompts.
	-e, --export FILENAME         Export to a filename; auto-names otherwise.
	-f, --filetype FILETYPE       Pick a filetype; auto-detects otherwise.
	-h, --help                    Print this help message.
	-l, --list                    List available algorithms and filetypes.
	-m, --magic MAGICFILE         Use a user-created magic file.
	-o, --overwrite               Overwrite the input file (dangerous!).
	-s, --strategy STRATEGY       Either smart (default), prefix, or suffix.
`

func main() {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	var (
		strategyFlag, filetypeFlag, magicFlag, algorithmFlag, exportFlag, checksumFlag string
		overwriteFlag, dangerousNoConfirmFlag, listFlag, helpFlag                      bool
	)

	flag.StringVar(&strategyFlag, "s", "smart", "")
	flag.StringVar(&strategyFlag, "strategy", "smart", "")
	flag.StringVar(&filetypeFlag, "f", "auto", "")
	flag.StringVar(&filetypeFlag, "filetype", "auto", "")
	flag.StringVar(&magicFlag, "m", "", "")
	flag.StringVar(&magicFlag, "magic", "", "")
	flag.StringVar(&algorithmFlag, "a", "", "")
	flag.StringVar(&algorithmFlag, "algorithm", "", "")
	flag.StringVar(&exportFlag, "e", "", "")
	flag.StringVar(&exportFlag, "export", "", "")
	flag.StringVar(&checksumFlag, "c", "", "")
	flag.StringVar(&checksumFlag, "checksum", "", "")
	flag.BoolVar(&overwriteFlag, "o", false, "")
	flag.BoolVar(&overwriteFlag, "overwrite", false, "")
	flag.BoolVar(&dangerousNoConfirmFlag, "dangerous-no-confirm", false, "")
	flag.BoolVar(&listFlag, "l", false, "")
	flag.BoolVar(&listFlag, "list", false, "")
	flag.BoolVar(&helpFlag, "h", false, "")
	flag.BoolVar(&helpFlag, "help", false, "")
	flag.Parse()

	if flag.NArg() == 0 || helpFlag {
		flag.Usage()
		return
	}
	if flag.NArg() > 1 {
		log.Fatalln("error: too many arguments provided")
	}
	if listFlag {
		fmt.Println("listFlag: print out a list of all supported algorithms and filetypes")
		return
	}
	if filetypeFlag != "" && magicFlag != "" {
		log.Fatalln("error: can't set --filetype and --magic at the same time")
	}

	(&crcedit{
		strategy:  strategyFlag,
		filetype:  filetypeFlag,
		algorithm: algorithmFlag,
		checksum:  checksumFlag,

		inputFilename:  flag.Arg(0),
		outputFilename: exportFlag,
		magicFilename:  magicFlag,

		shouldOverwrite: overwriteFlag,
		shouldConfirm:   !dangerousNoConfirmFlag,
	}).Run()
}

type crcedit struct {
	strategy, filetype, algorithm, checksum      string
	inputFilename, outputFilename, magicFilename string
	shouldOverwrite, shouldConfirm               bool
}

func (c crcedit) Run() {
	fmt.Println("crcedit.Run(): ran")
	return
}
