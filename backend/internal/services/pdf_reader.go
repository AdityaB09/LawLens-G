package services

import (
	"bytes"

	pdf "github.com/zacharysyoung/rsc-thuc-pdf"
)

// ReadPDFToText returns plain text from a PDF file path.
func ReadPDFToText(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer func() { _ = f.Close() }()

	reader, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(reader); err != nil {
		return "", err
	}
	return buf.String(), nil
}
