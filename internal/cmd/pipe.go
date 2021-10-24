package cmd

import (
	"bufio"
	"errors"
	"github.com/gaitr/goprobe/internal"
	"github.com/gaitr/goprobe/internal/util"
	"os"
)

func fileRead() error {
	file, e := getFile()
	if e != nil {
		return e
	}
	defer file.Close()

	return scanList(file)
}

func scanList(input *os.File) error {
	scanner := bufio.NewScanner(bufio.NewReader(input))
	for scanner.Scan() {
		e := internal.Router(client, scanner.Text(), &flagPool)
		if e != nil {
			return e
		}
	}
	return nil
}

func getFile() (*os.File, error) {
	if flags.filepath == "" {
		return nil, errors.New("please input a file")
	}
	if !util.FileExist(flags.filepath) {
		return nil, errors.New("the file provided does not exist")
	}
	file, e := os.Open(flags.filepath)
	if e != nil {
		return nil, errors.New("error")
	}
	return file, nil
}

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}
