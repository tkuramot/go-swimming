package piscine

import (
	"os"
)

func Ztail(args []string, numBytes int64) bool {
	ok := true

	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			PrintErrln(err.Error())
			ok = false
			continue
		}

		s, ok := tail(file, numBytes)
		if ok {
			if Len(args) > 1 {
				Println("==> " + arg + " <==")
			}
			Print(*s)
		} else {
			ok = false
		}
	}

	return ok
}

func tail(file *os.File, numBytes int64) (*string, bool) {
	stat, err := file.Stat()
	if err != nil {
		PrintErrln(err.Error())
		return nil, false
	}
	fileSize := stat.Size()

	var startPos int64
	if fileSize > numBytes {
		startPos = fileSize - numBytes
	}

	_, err = file.Seek(startPos, 0)
	if err != nil {
		PrintErrln(err.Error())
		return nil, false
	}

	buf := make([]byte, numBytes)
	n, err := file.Read(buf)
	if err != nil {
		PrintErrln(err.Error())
		return nil, false
	}

	str := string(buf[:n])
	return &str, true
}
