package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `Usage:
    crcedit FILENAME CHECKSUM -a ALGORITHM

By default, crcedit will create a new file with a '.crcedit' extension - your
files won't be overwritten unless you specify it. Additional documentation can
be found at https://github.com/clfs/crcedit inside the README.

Options:
	-a, --algorithm ALGORITHM     Pick a checksum algorithm.
	    --dangerous-no-confirm    Disable all confirmation prompts.
	-e, --export FILENAME         Export to a filename; auto-names otherwise.
	-f, --filetype FILETYPE       Pick a filetype; auto-detects otherwise.
	-l, --list                    List available algorithms and filetypes.
	-m, --magic MAGICFILE         Use a user-created magic file.
	-o, --overwrite               Overwrite the input file (dangerous!).
	-s, --strategy STRATEGY       Either smart (default), prefix, or suffix.
`

func main() {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	var (
		strategyFlag, filetypeFlag, magicFlag, algorithmFlag, exportFlag string
		overwriteFlag, dangerousNoConfirmFlag, listFlag                  bool
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
	flag.BoolVar(&overwriteFlag, "o", false, "")
	flag.BoolVar(&overwriteFlag, "overwrite", false, "")
	flag.BoolVar(&dangerousNoConfirmFlag, "dangerous-no-confirm", false, "")
	flag.BoolVar(&listFlag, "l", false, "")
	flag.BoolVar(&listFlag, "list", false, "")
	flag.Parse()
}
