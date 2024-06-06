package mog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

/*
fmt.Fprintf(mog.MW, "\n %s something %d, once = %t", me.Id, 2, true)

fmt.Fprintln(mog.MW, "\n hello")
*/
var MW io.Writer = nil

var verboseOn = false

func Init(verbose bool) error {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	mogsDir := fmt.Sprintf("%s%s%s", exPath, string(os.PathSeparator), "mogs")
	errDir := os.MkdirAll(mogsDir, os.ModePerm)
	if errDir != nil {
		return errDir
	}

	dir := fmt.Sprintf("%s%s", mogsDir, string(os.PathSeparator))

	filePath := dir + time.Now().Format("2006.01.02.15.04.05") + ".mog.txt"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	MW = io.MultiWriter(os.Stdout, file)

	verboseOn = verbose

	return nil
}

func Error(fmtStr string, rest ...any) {
	fmt.Fprintf(MW, "\n[Error] "+fmtStr, rest...)
}

func Warn(fmtStr string, rest ...any) {
	fmt.Fprintf(MW, "\n[Warn] "+fmtStr, rest...)
}

func Verbose(fmtStr string, rest ...any) {
	if !verboseOn {
		return
	}

	fmt.Fprintf(MW, "\n"+fmtStr, rest...)
}
