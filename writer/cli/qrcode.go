package cli

import (
	"io"

	"github.com/fox-one/hodl/writer"
	"github.com/mdp/qrterminal"
)

func QRCodeWriter(w io.Writer) io.Writer {
	return &qrWriter{w}
}

type qrWriter struct {
	w io.Writer
}

func (w *qrWriter) Write(p []byte) (n int, err error) {
	if raw, _, ok := writer.ExtractLabel(string(p)); ok {
		qrterminal.GenerateHalfBlock(raw, qrterminal.H, w.w)
		p = []byte(raw)
	}

	return w.w.Write(p)
}
