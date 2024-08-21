package document

import (
	"bytes"
	"os/exec"
)

func ConvertPDFToText(filepath string) (string, error) {
	cmd := exec.Command("pdftotext", filepath, "-")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
