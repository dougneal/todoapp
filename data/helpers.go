package data

import (
	"os"
	"path/filepath"
)

func dataDirPath() string {
	var path string
	var err error

	base := os.Getenv("DATA_DIR")
	if base == "" {
		path, err = filepath.Abs("db")
	} else {
		path, err = filepath.Abs(filepath.Join(base, "db"))
	}

	if err != nil {
		panic(err)
	}

	return path
}
