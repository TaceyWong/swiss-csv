package scsv

import (
	"bytes"
	"io"
)

func LineCounterWC(r io.Reader) (num uint64, err error) {
	return
}

func LineCounter(r io.Reader) (count uint64, err error) {
	buf := make([]byte, 32*1024)
	count = 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += uint64(bytes.Count(buf[:c], lineSep))
		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
