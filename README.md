# crcedit
`crcedit` is a Go CLI tool for editing file checksums.

## Project Status
Most key functionality is missing. This isn't in a usable state yet.

## Installation
Clone the repository and install with `go install`.
```bash
git clone https://github.com/clfs/crcedit
cd crcedit
go install
```

Afterwards (if you've configured Go correctly), you can call it from
the command line.
```bash
crcedit -help
```

## Usage
Here's the usage message. This might be less up-to-date than the actual
code itself, so check there for the source of truth.
```
Usage:
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
```

## Roadmap
- Stub out all functionality.
- TBD

## Contributing
I'm not open to contributions just yet, but issues are fine.

## License
[MIT](https://choosealicense.com/licenses/mit/)