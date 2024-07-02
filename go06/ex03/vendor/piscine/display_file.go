package piscine

import "os"

const BufferSize = 1024

func DisplayFile(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		PrintErrln(err.Error())
		return
	}
	defer file.Close()

	ReadAndPrintFile(file)
}

func ReadAndPrintFile(file *os.File) {
	buf := [BufferSize]byte{}
	for {
		n, err := file.Read(buf[:])
		if n > 0 {
			Print(string(buf[:n]))
		}
		if err != nil {
			break
		}
	}
}
