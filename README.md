# goprobe

### Build
```bash
go build -o ./bin/goprobe
```

### Example
```bash
> goprobe https://github.com/gaitr/goprobe
> cat links.txt | goprobe
> echo "https://github.com/gaitr/goprobe" | goprobe
```

### Command line options
```bash
> goprobe --help
goprobe is a CLI app to check the Status of accessible URLs

Usage:
  goprobe [flags]

Flags:
  -f, --file string   path to the file
  -G, --get           to send get request
  -h, --help          help for goprobe
  -v, --verbose       log verbose output
      --version       version for goprobe
```
