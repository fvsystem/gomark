package cli

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Execute()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if !strings.Contains(string(out), "Executing cli") || !strings.Contains(string(out), "Creating tests") {
		t.Errorf("Expected %s, got %s", "Executing cli", out)
	}
	Execute()
}
