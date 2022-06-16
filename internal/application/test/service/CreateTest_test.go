package service

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCresteTest(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var createTest = &CreateTestService{}
	createTest.CreateTest()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if !strings.Contains(string(out), "Creating tests") {
		t.Errorf("Expected %s, got %s", "Creating tests", out)
	}
}
