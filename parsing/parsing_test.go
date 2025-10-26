package parsing

import (
	"os"
	"testing"
)

func TestParsingUtils(t *testing.T) {
	testCases := []struct {
		inputFile  string
		outputFile string
	}{
		{"../testfiles/action-composite.yml", "../testfiles/readme-composite.md"},
		{"../testfiles/action-js.yml", "../testfiles/readme-js.md"},
	}
	for _, tc := range testCases {
		dt, err := ParseYml(tc.inputFile)
		if err != nil {
			t.Errorf("Not expecting any error, got %s", err.Error())
		}
		mdStr, err := ParseActionData(dt)
		if err != nil {
			t.Errorf("Not expecting any error, got %s", err.Error())
		}
		cont, _ := os.ReadFile(tc.outputFile)
		if mdStr != string(cont) {
			t.Errorf("Expecting %s when converting %s, got %s", cont, tc.outputFile, mdStr)
		}
	}
}
