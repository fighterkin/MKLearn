package mqbasic
import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"
	"time"
)

type writer struct {
	w io.Writer
}

func (w *writer) WriteFrame(frame frame) (err error) {
	if err = frame.write(w.w); err != nil {
		return
	}
	if buf, ok := w.w(*bufio.Writer); ok {
		err = buf.Flush()		
	}
	return 
}

func(f *methodFrame) write(w io.Writer) (err error) {
	var payload byte.Buffer
	if f.Method == nil {
		return errors.New("malformed frame: missing method")	
	}
}
