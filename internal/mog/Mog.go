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
var hasBeenCalled = false
var logFile *os.File = nil
var fileCounter = 0
var ts = time.Now().Format("2006.01.02.15.04.05")
var cache = make([]func(), 0)
var hasInit = false

func Init(verbose bool) error {
	if hasInit {
		return nil
	}

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

	filePath := fmt.Sprintf("%s%s--%d.crowform.mog.txt", dir, ts, fileCounter)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	logFile = file

	MW = io.MultiWriter(os.Stdout, file)

	verboseOn = verbose

	hasInit = true

	return nil
}

func Error(fmtStr string, rest ...any) {
	cacheOrLog(func() {
		logError(fmtStr, rest...)
	})
}
func logError(fmtStr string, rest ...any) {
	hasBeenCalled = true
	fmt.Fprintf(MW, "\n[Error] "+fmtStr, rest...)
}

func Warn(fmtStr string, rest ...any) {
	cacheOrLog(func() {
		logWarn(fmtStr, rest...)
	})
}
func logWarn(fmtStr string, rest ...any) {
	hasBeenCalled = true
	fmt.Fprintf(MW, "\n[Warn] "+fmtStr, rest...)
}

func Verbose(fmtStr string, rest ...any) {
	cacheOrLog(func() {
		logVerbose(fmtStr, rest...)
	})
}
func logVerbose(fmtStr string, rest ...any) {
	if !verboseOn {
		return
	}

	hasBeenCalled = true
	fmt.Fprintf(MW, "\n"+fmtStr, rest...)
}

func Debug(fmtStr string, rest ...any) {
	cacheOrLog(func() {
		logDebug(fmtStr, rest...)
	})
}
func logDebug(fmtStr string, rest ...any) {
	var dTs = time.Now().Format("2006.01.02.15.04.05")
	fmt.Fprintf(MW, "\nDebug ["+dTs+"] "+fmtStr, rest...)
}

func CleanUp() {
	if hasBeenCalled || logFile == nil || !hasInit {
		return
	}

	err1 := logFile.Close()
	if err1 != nil {
		// panic(err1)
		return // Already Closed
	}

	err := os.Remove(logFile.Name())
	if err != nil {
		panic(err)
	}

	logFile = nil
	hasInit = false
	hasBeenCalled = false
}

func checkFileSize() {
	if !hasInit || logFile == nil {
		return
	}

	fi, err := logFile.Stat()
	if err != nil {
		// panic(err)
		return //@TODO investigate Panics because of handle sometimes
	}

	// fmt.Printf("The file is %d bytes long", fi.Size())
	if fi.Size() > 2048*1024 {
		err1 := logFile.Close()
		if err1 != nil {
			// panic(err1)
			return // Already Closed
		}

		logFile = nil
		hasInit = false
		hasBeenCalled = false

		fileCounter++
		Init(verboseOn)
	}
}

func cacheOrLog(logFunc func()) {
	if logFile == nil {
		cache = append(cache, logFunc)
		return
	}

	if len(cache) > 0 {
		for i := range cache {
			if i >= len(cache) {
				break
			}
			cache[i]()
		}

		cache = nil
	}

	logFunc()
	go checkFileSize()
	// checkFileSize()
}
