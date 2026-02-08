package constants

import (
	"log"
	"path/filepath"
	"runtime"
)

/*
	How many dirs the root dir is
		above the current file
*/

const parentLvls = 4

var LocalDir string
var LogDir string
var LogFilePath string

func init() {
	bd := BaseDir()
	LocalDir = filepath.Join(bd, "_local")
	LogDir = filepath.Join(bd, "logs")
	LogFilePath = filepath.Join(LogDir, "app.log")
}

func BaseDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error getting BaseDir")
	}
	baseDir := filename
	for range parentLvls {
		baseDir = filepath.Dir(baseDir)
	}
	return baseDir
}
